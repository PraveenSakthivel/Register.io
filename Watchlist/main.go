package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"main/models"
	"net/http"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func amain() {
	// Router & template Setup
	router := gin.Default()

	// Intiialize SQLite DB
	models.ConnectDB()

	router.POST("/AddToList", func(c *gin.Context) {
		netid := c.PostForm("NetID")
		course := c.PostForm("Course")
		entry := models.Watchlist{Netid: netid, Course: course}
		result := models.DB.Create(&entry)
		if result.Error == nil {
			c.String(http.StatusOK, "Created")
		} else {
			c.String(http.StatusInternalServerError, "Error Creating Entry")
		}
	})

	router.POST("/DropFromList", func(c *gin.Context) {
		netid := c.PostForm("NetID")
		course := c.PostForm("Course")
		result := models.DB.Where("netid = ? AND course = ?", netid, course).Delete(&models.Watchlist{})
		if result.Error == nil {
			c.String(http.StatusOK, "Deleted")
		} else {
			c.String(http.StatusInternalServerError, "Error Deleting Entries")
		}
	})

	router.POST("/PingList", func(c *gin.Context) {
		course := c.PostForm("Course")
		smtpServer := "smtp.gmail.com"
		auth := smtp.PlainAuth(
			"",
			"registeriowatchlist",
			os.Getenv("EMAILPASS"),
			"smtp.gmail.com",
		)
		entries := []models.Watchlist{}
		from := mail.Address{Name: "Register.io", Address: "registeriowatchlist@gmail.com"}
		models.DB.Where("course = ?", course).Find(&entries)

		for _, entry := range entries {
			recip := entry.Netid + "@scarletmail.rutgers.edu"
			to := mail.Address{Name: "", Address: recip}
			title := "Register.io=?utf-8?Q?=F0=9F=93=9A?= | Opening for " + entry.Course
			body := "There is currently an open space for the course: " + entry.Course + "\r\nThanks for using Register.io📚"
			header := make(map[string]string)
			header["From"] = from.String()
			header["To"] = to.String()
			header["Subject"] = title
			header["MIME-Version"] = "1.0"
			header["Content-Type"] = "text/plain; charset=\"utf-8\""
			header["Content-Transfer-Encoding"] = "base64"
			message := ""
			for k, v := range header {
				message += fmt.Sprintf("%s: %s\r\n", k, v)
			}
			message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

			err := smtp.SendMail(
				smtpServer+":587",
				auth,
				from.Address,
				[]string{to.Address},
				[]byte(message),
			)
			if err != nil {
				log.Fatal(err)
				c.String(http.StatusInternalServerError, "Error sending emails")
				return
			}
		}

		c.String(http.StatusOK, "Emails sent")
	})

	router.Run()
}
