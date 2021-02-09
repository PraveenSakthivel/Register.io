package main

import (
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

	// Present login form
	router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login", gin.H{})
	})

	// Present not found page
	router.GET("/notfound/", func(c *gin.Context) {
		// Get type url parameter
		// If param = "login" -> present invalid credentials, else present username already exists
		c.HTML(200, "notfound", gin.H{"text": "Invalid credentials"})
	})

	// Process user login
	router.POST("/login_user", func(c *gin.Context) {
		// Generate token
		token := loginController.Login(c)
		if token != "" {
			encToken := middleware.Encrypt(token)
			// Set token to cookie & send back home
			c.SetCookie("token", encToken, 48*60, "/", "", false, false)
			// c.Redirect(http.StatusFound, "/")
			c.JSON(200, gin.H{"token": encToken})
			return
		}
		// redirect to error page
		c.Redirect(http.StatusFound, "/notfound")
	})

	// Logout user
	router.GET("/logout", func(c *gin.Context) {
		// delete token cookie and send home
		c.SetCookie("token", "", -1, "/", "", false, false)
		c.Redirect(http.StatusFound, "/")
	})

	router.Run()
}
