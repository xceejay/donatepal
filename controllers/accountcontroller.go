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
	username  string
	firstname string
	lastname  string
	email     string
	address   string
	country   string
	city      string
	age       uint
}

const (
	userkey = "user"
)

/////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////HANDLERS////////////////////////////////////////////////

// HandleLogin is a simple middleware to login
func (accountcontroller AccountController) HandleLogin(c *gin.Context) {
	// session := sessions.Default(c)
	// user := session.Get(userkey)
	// usernameSessionstring := fmt.Sprintf("%s", user)
	// if user == nil {
	// 	c.HTML(http.StatusOK, "login.html", nil)
	// } else {
	// 	c.Redirect(http.StatusPermanentRedirect, "/account/"+usernameSessionstring)
	// }

	c.HTML(http.StatusOK, "login.html", nil)

}

func (accountcontroller AccountController) HandleLogout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	// session.Delete(userkey)

	session.Set("user", "") // this will mark the session as "written" and hopefully remove the username
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1}) // this sets the cookie with a MaxAge of 0

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

// func (accountController AccountController) HandleAccountPage(c *gin.Context) {
// 	session := sessions.Default(c)
// 	username := session.Get(userkey)

// 	usernameSessionstring := fmt.Sprintf("%s", username)

// 	myUrl, err := url.Parse(c.Request.RequestURI)

// 	if err != nil {
// 		log.Errorf("ERROR: %s\nHandleAccountPage: Could Not Parse Request URI\nRedirecting to 404\n", err)
// 		c.Redirect(http.StatusNotFound, "/404")
// 	}
// 	urlUsername := path.Base(myUrl.Path)
// 	fmt.Printf("Request URI: %s\nBase: %s\n", c.Request.RequestURI, urlUsername)

// 	if usernameSessionstring != urlUsername {

// 		c.Redirect(http.StatusPermanentRedirect, "/login")
// 		return
// 	}

// 	if urlUsername == usernameSessionstring {

// 		location := fmt.Sprintf("/%s/%s", "account", username)

// 		fmt.Println("LOCATION:", location)
// 		accountController.ServeAccountPage(c)
// 	}
// }

func (accountController AccountController) HandleAdminAccountPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get(userkey)

	usernameSessionstring := fmt.Sprintf("%s", username)

	myUrl, err := url.Parse(c.Request.RequestURI)

	if err != nil {
		log.Errorf("ERROR: %s\nHandleAccountPage: Could Not Parse Request URI\nRedirecting to 404\n", err)
		c.Redirect(http.StatusNotFound, "/404")
	}
	urlUsername := path.Base(myUrl.Path)
	fmt.Printf("Request URI: %s\nBase: %s\n", c.Request.RequestURI, urlUsername)

	if usernameSessionstring != urlUsername {

		c.Redirect(http.StatusPermanentRedirect, "/login")
		return
	}

	if urlUsername == usernameSessionstring {

		location := fmt.Sprintf("/%s/%s", "account", username)

		fmt.Println("LOCATION:", location)
		accountController.ServeAdminAccountPage(c)
	}
}

func (accountController AccountController) HandlePayment(c *gin.Context) {
	c.HTML(http.StatusOK, "payment.html", nil)

}

/////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////MISC////////////////////////////////////////////////

// // login is a handler that parses a form and checks for specific data
func (accountcontroller AccountController) PerformLogin(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match from a database
	if username != "admin" || password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// session.Options(sessions.Options{MaxAge:	})
	// Save the username in the session
	session.Set(userkey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusPermanentRedirect, "/account/"+username)

}

func (accountController AccountController) ServeAdminAccountPage(c *gin.Context) {
	accountController.ServeAdminAccountOverviewPage(c)
}

func (accountController AccountController) ServeAdminAccountOverviewPage(c *gin.Context) {
	content := accountController.getDashboardContent("overview", c)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (accountController AccountController) ServeAdminAccountTransactionPage(c *gin.Context) {
	content := accountController.getDashboardContent("transactions", c)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (accountController AccountController) ServeAdminAccountBalancePage(c *gin.Context) {
	content := accountController.getDashboardContent("balance", c)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (accountController AccountController) ServeAdminAccountRecieptPage(c *gin.Context) {
	content := accountController.getDashboardContent("reciept", c)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (accountController AccountController) ServeAdminAccountSettingsPage(c *gin.Context) {
	content := accountController.getDashboardContent("settings", c)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (accountController AccountController) getDashboardContent(page string, c *gin.Context) string {
	paths := []string{
		"views/html/account/header.html",
		"views/html/account/footer.html",
		"views/html/account/account_top.html",
		"views/html/account/account_bottom.html",
		"views/html/account/overview.html",
		"views/html/account/transactions.html",
		"views/html/account/balance.html",
		"views/html/account/reciept.html",
		"views/html/account/settings.html",
	}
	vars := make(map[string]interface{})

	user := accountController.GetAllUserData(c.Query("username"))
	vars["username"] = user.username
	vars["email"] = user.email
	vars["firstname"] = user.firstname
	vars["lastname"] = user.lastname
	vars["address"] = user.address
	vars["city"] = user.city
	vars["country"] = user.country

	templateEngine := new(services.TemplateEngine)

	header := templateEngine.ProcessFile(paths[0], vars)
	footer := templateEngine.ProcessFile(paths[1], vars)
	accounthtmlPageBottom := templateEngine.ProcessFile(paths[3], vars)
	overviewContent := templateEngine.ProcessFile(paths[4], vars)

	switch page {
	case "transactions":
		vars["transactions_active"] = "active"
		accounthtmlPageTop := templateEngine.ProcessFile(paths[2], vars)
		transactionContent := templateEngine.ProcessFile(paths[5], vars)

		return header + accounthtmlPageTop + transactionContent + accounthtmlPageBottom + footer

	case "balance":

		vars["balance_active"] = "active"
		accounthtmlPageTop := templateEngine.ProcessFile(paths[2], vars)
		balanceContent := templateEngine.ProcessFile(paths[6], vars)

		return header + accounthtmlPageTop + balanceContent + accounthtmlPageBottom + footer

	case "reciept":
		vars["reciept_active"] = "active"
		accounthtmlPageTop := templateEngine.ProcessFile(paths[2], vars)
		recieptContent := templateEngine.ProcessFile(paths[7], vars)

		return header + accounthtmlPageTop + recieptContent + accounthtmlPageBottom + footer

	case "settings":
		vars["settings_active"] = "active"
		accounthtmlPageTop := templateEngine.ProcessFile(paths[2], vars)
		settingsContent := templateEngine.ProcessFile(paths[8], vars)

		return header + accounthtmlPageTop + settingsContent + accounthtmlPageBottom + footer

	default:
		vars["overview_active"] = "active"
		accounthtmlPageTop := templateEngine.ProcessFile(paths[2], vars)

		return header + accounthtmlPageTop + overviewContent + accounthtmlPageBottom + footer

	}

}

// func (accountController AccountController) ServeAccountPage(c *gin.Context) {

// 	paths := []string{
// 		"views/html/account/header.html",
// 		"views/html/account/footer.html",
// 		"views/html/account/account_top.html",
// 		"views/html/account/account_bottom.html",
// 		"views/html/account/overview.html",
// 		"views/htmlsettings/account/transactions.html",
// 		"views/html/account/balance.html",
// 		"views/html/account/reciept.html",
// 		"views/html/account/settings.html",
// 	}
// 	vars := make(map[string]interface{})

// 	user := accountController.GetAllUserData(c.Query("username"))
// 	vars["name"] = user.name

// 	templateEngine := new(services.TemplateEngine)

// 	header := templateEngine.ProcessFile(paths[0], vars)
// 	footer := templateEngine.ProcessFile(paths[1], vars)
// 	accounthtmlPageTop := templateEngine.ProcessFile(paths[2], vars)
// 	accounthtmlPageBottom := templateEngine.ProcessFile(paths[3], vars)
// 	overviewContent := templateEngine.ProcessFile(paths[4], vars)
// 	transactionContent := templateEngine.ProcessFile(paths[5], vars)
// 	balanceContent := templateEngine.ProcessFile(paths[6], vars)
// 	recieptContent := templateEngine.ProcessFile(paths[7], vars)
// 	settingsContent := templateEngine.ProcessFile(paths[8], vars)

// 	page := "home"
// 	switch page {
// 	case "transactions":
// 		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(header+accounthtmlPageTop+transactionContent+accounthtmlPageBottom+footer))
// 		return
// 	case "balance":
// 		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(header+accounthtmlPageTop+balanceContent+accounthtmlPageBottom+footer))

// 		return
// 	case "reciept":
// 		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(header+accounthtmlPageTop+recieptContent+accounthtmlPageBottom+footer))

// 		return
// 	case "settings":
// 		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(header+accounthtmlPageTop+settingsContent+accounthtmlPageBottom+footer))

// 		return
// 	default:
// 		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(header+accounthtmlPageTop+overviewContent+accounthtmlPageBottom+footer))
// 		return

// 	}

// }

// func status(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
// }

func isLoggedIn(c *gin.Context) bool {

	session := sessions.Default(c)
	user := session.Get(userkey)

	if user == nil {

		return false
	}
	usernameSessionstring := fmt.Sprintf("%s", user)

	return usernameSessionstring == "admin"

}

// func me(c *gin.Context) {
// 	session := sessions.Default(c)
// 	user := session.Get(userkey)
// 	c.JSON(http.StatusOK, gin.H{"user": user})
// }

// func (accountController AccountController) ServeTemplate(c *gin.Context) {

// 	c.HTML(http.StatusOK, "template.html", nil)

// }

func (accountController AccountController) GetAllUserData(username string) *User {
	user := &User{
		username:  "admin",
		firstname: "Joel",
		lastname:  "Amoako",
		email:     "joelkofiamoako@gmail.com",
		address:   "N/A",
		country:   "N/A",
		city:      "N/A",
		age:       21,
	}

	return user

}
