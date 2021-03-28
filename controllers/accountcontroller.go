package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/xceejay/boilerplate/services"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

type AccountController struct {
}

type User struct {
	username string
	password string
	name     string
	age      string
}

func (accountController AccountController) HandleLogin(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {

		c.HTML(http.StatusOK, "login.html", nil)
		return
	}

	if cookie == "123456" {

		accountController.PerformLogin(c)
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func (accountController AccountController) PerformLogin(c *gin.Context) {
	username := c.PostForm("username")

	if authenticate(c) {
		if username != "" {
			c.Redirect(http.StatusTemporaryRedirect, "/account/"+username)
			return
		}

		cookieusername, err := c.Cookie("username")
		if err != nil {
			c.Redirect(http.StatusUnauthorized, "/")
		}
		if cookieusername != "" && cookieusername == "xceejay" {
			c.Redirect(http.StatusTemporaryRedirect, "/account/"+cookieusername)
		}
	}

}

func (accountController AccountController) ServeAccountPage(c *gin.Context) {

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

func (accountController AccountController) HandleAccountPage(c *gin.Context) {
	tokencookie, err := c.Cookie("token")
	usernamecookie, err := c.Cookie("username")

	myUrl, err := url.Parse(c.Request.RequestURI)
	if err != nil {
		log.Fatal(err)
	}
	urlUsername := path.Base(myUrl.Path)
	fmt.Printf("Request URI: %s\nBase: %s\n", c.Request.RequestURI, urlUsername)
	if err != nil {
		c.Redirect(http.StatusUnauthorized, "/")
	}
	if tokencookie == "123456" && usernamecookie == "xceejay" && usernamecookie == urlUsername {

		accountController.ServeAccountPage(c)
	} else {

		c.Redirect(http.StatusUnauthorized, "/")

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
		c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		c.SetCookie("username", user.username, 3600, "/", "localhost", false, true)

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
