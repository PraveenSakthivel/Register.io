package middleware

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	Tokens "main/protobuf"
	secret "main/secrets"
	"main/service"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

// TokenSecret ...
type TokenSecret struct {
	Token string `json:"TokenSecret"`
}

// Decrypt ...
func Decrypt(encryptedString string) (decryptedString string) {
	tokenSecret, _ := secret.GetTokenSecret("user/JWTEncryption")
	tokenObj := TokenSecret{}
	json.Unmarshal([]byte(tokenSecret), &tokenObj)
	// key, _ := hex.DecodeString(tokenObj.Token[0:32])
	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher([]byte(tokenObj.Token[0:32]))
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	nonceSize := aesGCM.NonceSize()

	if len(enc) < nonceSize {
		return ""
	}

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	return fmt.Sprintf("%s", plaintext)
}

// Encrypt ...
func Encrypt(stringToEncrypt string) (encryptedString string) {
	tokenSecret, _ := secret.GetTokenSecret("user/JWTEncryption")
	tokenObj := TokenSecret{}
	json.Unmarshal([]byte(tokenSecret), &tokenObj)
	// key, _ := hex.DecodeString(tokenObj.Token[0:32])
	plaintext := []byte(stringToEncrypt)

	block, err := aes.NewCipher([]byte(tokenObj.Token[0:32]))
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

// GenBytesGRPC ...
func GenBytesGRPC(data string) []byte {
	bytes := []byte{}
	stringvals := strings.Split(data, "+")
	for index := 0; index < len(stringvals); index++ {
		value, err := strconv.Atoi(stringvals[index])
		if err != nil {
			return nil
		}
		bytes = append(bytes, byte(value))
	}
	return bytes
}

// GenBytes ...
func GenBytes(data string) []byte {
	bytes := []byte{}
	stringvals := strings.Split(data, " ")
	for index := 0; index < len(stringvals); index++ {
		value, err := strconv.Atoi(stringvals[index])
		if err != nil {
			return nil
		}
		bytes = append(bytes, byte(value))
	}
	return bytes
}

// ValidToken ...
func ValidToken(c *gin.Context) (*jwt.Token, bool) {
	// Get token from cookie and check if valid
	prototoken := &Tokens.Token{}
	cookiedata, err := c.Cookie("token")
	if err != nil {
		return nil, false
	}
	databytes := GenBytes(cookiedata)
	if databytes == nil || len(databytes) == 0 {
		return nil, false
	}
	err = proto.Unmarshal(databytes, prototoken)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	encTokenString := prototoken.Token
	fmt.Println(encTokenString)
	tokenString := Decrypt(encTokenString)
	fmt.Println(tokenString)
	token, err := service.JWTAuthService().ValidateToken(tokenString)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		return token, true
	}
	fmt.Println(err)
	return nil, false

}

// ValidTokenGRPC ...
func ValidTokenGRPC(tokenInput *Tokens.Token) (*jwt.Token, bool) {
	encTokenString := tokenInput.Token
	fmt.Println(encTokenString)
	databytes := GenBytes(encTokenString)
	if databytes == nil || len(databytes) == 0 {
		return nil, false
	}
	newToken := &Tokens.Token{}
	proto.Unmarshal(databytes, newToken)
	tokenString := Decrypt(newToken.Token)
	if tokenString == "" {
		return nil, false
	}
	fmt.Println(tokenString)
	token, err := service.JWTAuthService().ValidateToken(tokenString)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		return token, true
	}
	fmt.Println(err)
	return nil, false
}
