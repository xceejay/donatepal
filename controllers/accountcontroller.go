package controllers

import (
	"fmt"
	"net/http"

	"github.com/xceejay/boilerplate/services"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
}

type User struct {
	username string
	password string
	name     string
	age      string
}

func (accountController AccountController) ShowLoginPage(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)

}

func (accountController AccountController) PerformLogin(c *gin.Context) {

	if authenticate(c) {
		paths := []string{
			"views/html/account/account.html",
		}
		vars := make(map[string]interface{})

		user := accountController.GetAllUserData(c.Query("username"))
		vars["name"] = user.name

		templateEngine := new(services.TemplateEngine)

		accounthtmlPage := templateEngine.ProcessFile(paths[0], vars)

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(accounthtmlPage))

	}

}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

func (accountController AccountController) PerformLogout(c *gin.Context) {

	c.SetCookie("token", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")

}

func authenticate(c *gin.Context) bool {

	var user User
	user.username = c.PostForm("username")
	user.password = c.PostForm("password")
	fmt.Printf("username: %s\n", user.username)
	fmt.Printf("password: %s\n", user.password)

	if user.username == "xceejay" && user.password == "1234" {
		token := "123456"
		c.SetCookie("token", token, 3600, "", "", false, true)

		return true
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	return false

}

func (accountController AccountController) GetAllUserData(username string) *User {
	user := &User{
		username: "xceejay",
		name:     "Joel Kofi Amoako",
		age:      "21",
	}

	return user

}
