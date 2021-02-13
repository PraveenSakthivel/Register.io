package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"main/controller"
	"main/middleware"
	"main/models"
	"main/protobuf"
	"main/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func main() {
	// JWT login setup
	jwtService := service.JWTAuthService()
	loginController := controller.LoginHandler(jwtService)

	// Router & template Setup
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Intiialize SQLite DB
	models.ConnectDB()

	// Get index page
	router.GET("/", func(c *gin.Context) {
		token, valid := middleware.ValidToken(c)
		// If valid present login page
		if valid {
			c.HTML(200, "index", gin.H{"userobj": token.Claims.(jwt.MapClaims)["name"]})
			return
		}
		// Present standard welcome page
		c.HTML(200, "index", gin.H{})
	})

	router.GET("/authuser", func(c *gin.Context) {
		// Check if user has valid token
		token, valid := middleware.ValidToken(c)
		// If valid present login page
		if valid {
			message := &protobuf.Token{Token: token.Claims.(jwt.MapClaims)["name"].(string)}
			data, _ := proto.Marshal(message)
			stringarray := fmt.Sprint(data)
			stringarray = stringarray[1 : len(stringarray)-1]
			c.ProtoBuf(200, message)

			return
		}
		message := &protobuf.Token{Token: ""}
		c.ProtoBuf(200, message)
	})

	// Present login form
	router.GET("/login", func(c *gin.Context) {
		c.SetCookie("state", "login", 10*365*24*60*60, "/", "", false, false)
		c.Redirect(http.StatusFound, "/")
	})

	router.GET("/signup", func(c *gin.Context) {
		c.SetCookie("state", "signup", 10*365*24*60*60, "/", "", false, false)
		c.Redirect(http.StatusFound, "/")
	})

	router.POST("/signup_user", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		data := []byte(password)
		hash := md5.Sum(data)
		newpass := hex.EncodeToString(hash[:])
		users := []models.User{}
		models.DB.Where("email = ?", email).Find(&users)
		if len(users) == 0 {
			models.DB.Create(&models.User{Email: email, Password: newpass})
			c.SetCookie("state", "", 10*365*24*60*60, "/", "", false, false)
		} else {
			c.SetCookie("state", "notfound", 10*365*24*60*60, "/", "", false, false)
		}
		c.Redirect(http.StatusFound, "/")
	})

	// Process user login
	router.POST("/login_user", func(c *gin.Context) {
		// Generate token
		token := loginController.Login(c)
		if token != "" {
			encToken := middleware.Encrypt(token)
			// Set token to cookie & send back home
			message := &protobuf.Token{Token: encToken}
			data, err := proto.Marshal(message)
			stringarray := fmt.Sprint(data)
			stringarray = stringarray[1 : len(stringarray)-1]
			fmt.Println(stringarray)
			if err != nil {
				log.Fatal("marshaling error: ", err)
			}

			c.SetCookie("token", stringarray, 48*60, "/", "", false, false)

			c.SetCookie("state", "", 10*365*24*60*60, "/", "", false, false)
		} else {
			c.SetCookie("state", "invalid", 10*365*24*60*60, "/", "", false, false)
		}
		c.Redirect(http.StatusFound, "/")
	})

	// Logout user
	router.GET("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, false)
		c.Redirect(http.StatusFound, "/")
	})

	router.Run()
}
