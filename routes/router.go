package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xceejay/boilerplate/controllers"
)

func InitRouter(engine *gin.Engine) {

	accountController := new(controllers.AccountController)
	homeController := new(controllers.HomeController)
	engine.NoRoute(homeController.Handle404)
	engine.GET("/", homeController.ServeHomePage)
	engine.GET("/about", homeController.ServeAboutPage)
	engine.GET("/login", accountController.HandleLogin)
	engine.POST("/login", accountController.PerformLogin)
	// engine.POST("/login", accountController.PerformLogin)
	// engine.POST("/account/:username", accountController.ServeAccountPage)
	// engine.GET("/account/:username", accountController.HandleAccountPage)
	engine.POST("/account/admin", accountController.ServeAdminAccountPage)
	engine.GET("/account/admin", accountController.HandleAdminAccountPage)
	engine.GET("/account/admin/:dashboard_content", accountController.HandleAdminDashboardContent)

	engine.GET("/logout", accountController.HandleLogout)
	engine.GET("/payment", accountController.HandlePayment)

	// engine.GET("/template", accountController.ServeTemplate)

}
