package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xceejay/boilerplate/models"
	"github.com/xceejay/boilerplate/services"
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
	paths := []string{
		"views/html/home/donate.html",
	}
	vars := make(map[string]interface{})

	user := new(models.User)
	fundraisers, err := user.GetAllUserData()
	if err != nil {
		panic(err)
	}

	var fundraiserOptions string

	for _, fundraiser := range fundraisers {

		if len(fundraiser.Firstname.String) < 1 {
			continue
		}

		fundraiserOptions += "<option " + "value= '" + fundraiser.Username + "'>" + fundraiser.Firstname.String + " " + fundraiser.Lastname.String + "</option>\n"

	}

	vars["fundraiser_options"] = template.HTML(fundraiserOptions)
	templateEngine := new(services.TemplateEngine)

	donateHtmlFile := templateEngine.ProcessFile(paths[0], vars)

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(donateHtmlFile))
}

func (homeController HomeController) ServeRegistrationPage(c *gin.Context) {
	// get all the fundraiser but for now its hard coded
	c.HTML(http.StatusOK, "register.html", nil)
}

func (homeController HomeController) Handle404(c *gin.Context) {

	c.HTML(http.StatusNotFound, "404.html", nil)

}
