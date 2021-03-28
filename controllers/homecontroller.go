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
