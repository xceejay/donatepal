package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (homeController HomeController) ServeHomePage(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
}

func (homeController HomeController) ServeAboutPage(c *gin.Context) {

	c.HTML(http.StatusOK, "about.html", nil)
}

func (homeController HomeController) ServeDonationPage(c *gin.Context) {
	// get all the fundraiser but for now its hard coded
	c.HTML(http.StatusOK, "donate.html", nil)
}

func (homeController HomeController) ServeRegistrationPage(c *gin.Context) {
	// get all the fundraiser but for now its hard coded
	c.HTML(http.StatusOK, "register.html", nil)
}

func (homeController HomeController) Handle404(c *gin.Context) {

	c.HTML(http.StatusNotFound, "404.html", nil)

}
