package main

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"

	"main/controller"
	"main/middleware"
	"main/models"
	"main/service"
)

// Merge base html page with additional page and set name for each view
func tempRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/welcome.html")
	r.AddFromFiles("signup", "templates/base.html", "templates/signup.html")
	r.AddFromFiles("login", "templates/base.html", "templates/login.html")
	r.AddFromFiles("notfound", "templates/base.html", "templates/notfound.html")
	return r
}

func main() {
	// JWT login setup
	jwtService := service.JWTAuthService()
	loginController := controller.LoginHandler(jwtService)

	// Router & template Setup
	router := gin.Default()
	router.HTMLRender = tempRender()

	// Intiialize SQLite DB
	models.ConnectDB()

	// Get index page
	router.GET("/", func(c *gin.Context) {
		// Check if user has valid token
		token, valid := middleware.ValidToken(c)
		// If valid present login page
		if valid {
			c.HTML(200, "index", gin.H{"userobj": token.Claims.(jwt.MapClaims)["name"]})
			return
		}
		// Present standard welcome page
		c.HTML(200, "index", gin.H{})
	})

	// Present signup form
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup", gin.H{})
	})

	// Present login form
	router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login", gin.H{})
	})

	// Present not found page
	router.GET("/notfound/", func(c *gin.Context) {
		// Get type url parameter
		// If param = "login" -> present invalid credentials, else present username already exists
		login := c.Query("type")
		if login == "login" {
			c.HTML(200, "notfound", gin.H{"text": "Invalid credentials"})
			return
		}
		c.HTML(200, "notfound", gin.H{"text": "Username already exists"})
	})

	// Process user signup
	router.POST("/signup_user", func(c *gin.Context) {
		// get form data
		email := c.PostForm("email")
		password := c.PostForm("password")
		// hash password
		data := []byte(password)
		hash := md5.Sum(data)
		newpass := hex.EncodeToString(hash[:])
		// check if user already exists
		users := []models.User{}
		models.DB.Where("email = ?", email).Find(&users)
		// if not, add to DB
		if len(users) == 0 {
			models.DB.Create(&models.User{Email: email, Password: newpass})
			c.Redirect(http.StatusFound, "/")
			return
		}
		// if exists, route to error page
		c.Redirect(http.StatusFound, "/notfound")
	})

	// Process user login
	router.POST("/login_user", func(c *gin.Context) {
		// Generate token
		token := loginController.Login(c)
		if token != "" {
			// Set token to cookie & send back home
			c.SetCookie("token", token, 48*60, "/", "", false, false)
			c.Redirect(http.StatusFound, "/")
			return
		}
		// redirect to error page
		c.Redirect(http.StatusFound, "/notfound?type=login")
	})

	// Logout user
	router.GET("/logout", func(c *gin.Context) {
		// delete token cookie and send home
		c.SetCookie("token", "", -1, "/", "", false, false)
		c.Redirect(http.StatusFound, "/")
	})

	router.Run()
}
