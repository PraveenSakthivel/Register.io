package controller

import (
	"crypto/md5"
	"encoding/hex"
	"main/models"
	Tokens "main/protobuf"
	"main/service"

	"github.com/gin-gonic/gin"
)

var gradeMap = map[string]int{
	"A":  400,
	"B+": 350,
	"B":  300,
	"C+": 250,
	"C":  200,
	"D":  100,
	"F":  0,
}

// LoginCredentials ...
type LoginCredentials struct {
	NetID    string `form:"NetID"`
	Password string `form:"Password"`
	Classes  map[string]int
}

// LoginController ...
type LoginController interface {
	Login(ctx *gin.Context) string
	LoginEndpoint(creds *Tokens.Credentials) string
}

type loginController struct {
	jWtService service.JWTService
}

// LoginHandler ...
func LoginHandler(jWtService service.JWTService) LoginController {
	return &loginController{
		jWtService: jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	// Gets login form information and binds to struct
	var credential LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	// Hash and update password
	data := []byte(credential.Password)
	hash := md5.Sum(data)
	credential.Password = hex.EncodeToString(hash[:])

	// if user is valid, generate token
	isUserAuthenticated := service.LoginUser(credential.NetID, credential.Password)
	if isUserAuthenticated {
		classesList := []models.CourseHistory{}
		credential.Classes = make(map[string]int)
		models.DB.Where("netid = ?", credential.NetID).Find(&classesList)
		for _, item := range classesList {
			credential.Classes[item.CourseNumber] = gradeMap[item.Grade]
		}
		students := models.Students{}
		models.DB.Where("netid = ?", credential.NetID).First(&students)
		cases := students.SpecialCases
		casesArray := []int{}
		cases.Scan(casesArray)
		casesMap := make(map[int]bool)
		for _, value := range casesArray {
			casesMap[value] = true
		}
		user := models.User{}
		models.DB.Where("netid = ?", credential.NetID).Find(&user)
		return controller.jWtService.GenerateToken(credential.NetID, true, user.Type, credential.Classes, casesMap)
	}
	return ""
}

func (controller *loginController) LoginEndpoint(creds *Tokens.Credentials) string {
	// Gets login form information and binds to struct
	var credential LoginCredentials
	credential.NetID = creds.NetID
	credential.Password = creds.Password
	// Hash and update password
	data := []byte(credential.Password)
	hash := md5.Sum(data)
	credential.Password = hex.EncodeToString(hash[:])

	// if user is valid, generate token
	isUserAuthenticated := service.LoginUser(credential.NetID, credential.Password)
	if isUserAuthenticated {
		classesList := []models.CourseHistory{}
		credential.Classes = make(map[string]int)
		models.DB.Where("netid = ?", credential.NetID).Find(&classesList)
		for _, item := range classesList {
			credential.Classes[item.CourseNumber] = gradeMap[item.Grade]
		}
		students := models.Students{}
		models.DB.Where("netid = ?", credential.NetID).First(&students)
		cases := students.SpecialCases
		casesArray := []int{}
		cases.Scan(casesArray)
		casesMap := make(map[int]bool)
		for _, value := range casesArray {
			casesMap[value] = true
		}
		user := models.User{}
		models.DB.Where("netid = ?", credential.NetID).Find(&user)
		return controller.jWtService.GenerateToken(credential.NetID, true, user.Type, credential.Classes, casesMap)
	}
	return ""
}
