package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xceejay/boilerplate/controllers"
)

func InitRouter(engine *gin.Engine) {

	accountController := new(controllers.AccountController)
	homeController := new(controllers.HomeController)
	engine.GET("/", homeController.ServeHomePage)
	engine.GET("/about", homeController.ServeAboutPage)
	engine.GET("/login", accountController.HandleLogin)
	engine.POST("/login", accountController.PerformLogin)
	engine.POST("/account/:username", accountController.ServeAccountPage)
	engine.GET("/account/:username", accountController.HandleAccountPage)
	engine.GET("/logout", accountController.PerformLogout)

}
