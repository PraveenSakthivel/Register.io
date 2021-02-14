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
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func tempRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/welcome.html")
	r.AddFromFiles("signup", "templates/base.html", "templates/signup.html")
	r.AddFromFiles("login", "templates/base.html", "templates/login.html")
	r.AddFromFiles("notfound", "templates/base.html", "templates/notfound.html")
	r.AddFromFiles("viewreg", "templates/base.html", "templates/viewreg.html")
	// r.AddFromFiles("about", "templates/base.html", "templates/about.html")
	// r.AddFromFilesFuncs("about", template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }}, "templates/base.html", "templates/about.html")
	return r
}
func main() {
	// JWT login setup
	jwtService := service.JWTAuthService()
	loginController := controller.LoginHandler(jwtService)

	// Router & template Setup
	router := gin.Default()
	router.HTMLRender = tempRender()
	// router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Intiialize SQLite DB
	models.ConnectDB()

	// Get index page
	router.GET("/", func(c *gin.Context) {
		// token, valid := middleware.ValidToken(c)
		message := authuser(c)
		if message.Token != "" {
			// If valid present login page
			// if valid {
			c.HTML(200, "index", gin.H{"userobj": message.Token})
			return
		}
		// Present standard welcome page
		c.HTML(200, "index", gin.H{})
	})

	router.GET("/viewreg", func(c *gin.Context) {
		message := authuser(c)
		if message.Token != "" {

			regs := []models.CourseRegistration{}
			models.DB.Where("netid = ?", message.Token).Find(&regs)
			// fmt.Println(regs)
			classes := []models.Soc{}
			for _, reg := range regs {
				current := []models.Soc{}
				models.DB.Where("index = ?", reg.ClassIndex).First(&current)
				classes = append(classes, current...)
			}
			// fmt.Println(classes)

			c.HTML(200, "viewreg", gin.H{"classes": classes})
			return
		}
		c.Redirect(http.StatusFound, "login")

	})

	router.GET("/authuser", func(c *gin.Context) {
		message := authuser(c)
		c.ProtoBuf(200, message)
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup", gin.H{})
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login", gin.H{})
	})

	router.POST("/signup_user", func(c *gin.Context) {
		netid := c.PostForm("netid")
		password := c.PostForm("password")
		// fmt.Println(email)
		data := []byte(password)
		hash := md5.Sum(data)
		newpass := hex.EncodeToString(hash[:])
		users := []models.User{}
		models.DB.Where("netid = ?", netid).Find(&users)
		if len(users) == 0 {
			models.DB.Create(&models.User{Netid: netid, Password: newpass})
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.Redirect(http.StatusFound, "/notfound")
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
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.Redirect(http.StatusFound, "/notfound?type=login")
	})

	router.GET("/notfound/", func(c *gin.Context) {
		// Get type url parameter
		// If param = "login" -> present invalid credentials, else present username already exists
		if c.Query("type") == "login" {
			c.HTML(200, "notfound", gin.H{"text": "Invalid credentials"})
		} else {
			c.HTML(200, "notfound", gin.H{"text": "User already exists"})
		}
	})

	// Logout user
	router.GET("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, false)
		c.Redirect(http.StatusFound, "/")
	})

	router.Run()
}

func authuser(c *gin.Context) *protobuf.Token {
	// Check cookie value is set and if cookie corresponds to valid JWT
	token, valid := middleware.ValidToken(c)
	// If valid send username from JWT
	if valid {
		message := &protobuf.Token{Token: token.Claims.(jwt.MapClaims)["name"].(string)}
		data, _ := proto.Marshal(message)
		stringarray := fmt.Sprint(data)
		stringarray = stringarray[1 : len(stringarray)-1]
		// c.ProtoBuf(200, message)
		return message
	}
	// If not, send empty string
	message := &protobuf.Token{Token: ""}
	return message
	// c.ProtoBuf(200, message)
}
