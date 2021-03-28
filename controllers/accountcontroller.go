package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/xceejay/boilerplate/services"
)

type AccountController struct {
}

type User struct {
	username string
	password string
	name     string
	age      string
}

const (
	userkey = "user"
)

// AuthRequired is a simple middleware to check the session
func (accountcontroller AccountController) HandleLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// login is a handler that parses a form and checks for specific data
func (accountcontroller AccountController) PerformLogin(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if username != "xceejay" || password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set(userkey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusPermanentRedirect, "/account/"+username)

}

func (accountcontroller AccountController) HandleLogout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (accountController AccountController) HandleAccountPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get(userkey)
	usernameSessionstring := fmt.Sprintf("%s", username)

	myUrl, err := url.Parse(c.Request.RequestURI)

	if err != nil {
		log.Errorf("ERROR: %s\nHandleAccountPage: Could Not Parse Request URI\nRedirecting to 404\n", err)
		c.Redirect(http.StatusPermanentRedirect, "/404")
	}
	urlUsername := path.Base(myUrl.Path)
	fmt.Printf("Request URI: %s\nBase: %s\n", c.Request.RequestURI, urlUsername)

	if usernameSessionstring != urlUsername {

		c.Redirect(http.StatusPermanentRedirect, "/login")
		return
	}

	if urlUsername != "xceejay" {
		c.Redirect(http.StatusNotFound, "/404")
		return
	}

	if urlUsername == "xceejay" {

		accountController.ServeAccountPage(c)
		// location := fmt.Sprintf("/%s/%s", "account", username)

		// fmt.Println(location)
		// c.Redirect(http.StatusPermanentRedirect, location)
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

func (accountController AccountController) GetAllUserData(username string) *User {
	user := &User{
		username: "xceejay",
		name:     "Joel Kofi Amoako",
		age:      "21",
	}

	return user

}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
