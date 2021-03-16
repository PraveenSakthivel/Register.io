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

	// NOTE View Registration
	router.GET("/viewreg", func(c *gin.Context) {
		message := authuser(c)
		if message != nil {

			regs := []models.CourseRegistration{}
			models.DB.Where("netid = ?", message.Claims.(jwt.MapClaims)["name"]).Find(&regs)
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

	router.GET("/viewreg_endpoint", func(c *gin.Context) {
		message := authuser(c)
		if message != nil {

			regs := []models.CourseRegistration{}
			models.DB.Where("netid = ?", message.Claims.(jwt.MapClaims)["name"]).Find(&regs)
			// fmt.Println(regs)
			classes := []models.Soc{}
			for _, reg := range regs {
				current := []models.Soc{}
				models.DB.Where("index = ?", reg.ClassIndex).First(&current)
				classes = append(classes, current...)
			}
			// fmt.Println(classes)
			protomessage := protobuf.TokenList{}
			for _, class := range classes {
				protomessage.Token = append(protomessage.Token, class.Name)
			}
			c.ProtoBuf(http.StatusFound, &protomessage)
			return
		}
		protomessage := protobuf.TokenList{}
		c.ProtoBuf(http.StatusFound, &protomessage)

	})

	// Present login form
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", gin.H{})
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup", gin.H{})
	})

	// NOTE Signup User Logic
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
		} else {
			c.Redirect(http.StatusOK, "notfound")
			return
		}
		c.Redirect(http.StatusFound, "/")
		return
	})

	// NOTE Login User Logic
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

	// NOTE : Login User Endpoint
	// * Pass in Username & password
	// * Returns protobuf w/ encrypted JWT token
	// * JWT token contains netID of user
	router.POST("/login_user_endpoint", func(c *gin.Context) {
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
			c.ProtoBuf(http.StatusFound, message)
			return
		}
		message := &protobuf.Token{Token: ""}
		c.ProtoBuf(http.StatusFound, message)
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

	router.GET("/prevCourses", func(c *gin.Context) {
		message := authuser(c)
		if message != nil {
			fmt.Println("MESSAGE MAP CLAIMS")
			fmt.Println(message.Claims.(jwt.MapClaims))
			c.JSON(200, gin.H{"prev": message.Claims.(jwt.MapClaims)["classes"]})
			return
		}
		c.JSON(200, gin.H{})
	})

	router.Run()
}

// NOTE : authuser endpoint
// * Pass in context w/ cookie saved
// * Returns protobuf w/ netID
func authuser(c *gin.Context) *jwt.Token {
	// Check cookie value is set and if cookie corresponds to valid JWT
	token, valid := middleware.ValidToken(c)
	// If valid send username from JWT
	if valid {
		// message := &protobuf.Token{Token: token.Claims.(jwt.MapClaims)["name"].(string)}
		// data, _ := proto.Marshal(message)
		// stringarray := fmt.Sprint(data)
		// stringarray = stringarray[1 : len(stringarray)-1]
		// // c.ProtoBuf(200, message)
		// return message
		return token
	}
	// If not, send empty string
	// message := &protobuf.Token{Token: ""}
	// return message
	return nil
	// c.ProtoBuf(200, message)
}
