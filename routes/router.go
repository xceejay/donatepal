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
	engine.GET("/account/admin/transactions", accountController.ServeAdminAccountTransactionPage)
	engine.GET("/account/admin/balance", accountController.ServeAdminAccountBalancePage)
	engine.GET("/account/admin/reciept", accountController.ServeAdminAccountRecieptPage)
	engine.GET("/account/admin/settings", accountController.ServeAdminAccountSettingsPage)
	engine.GET("/account/admin/", accountController.ServeAdminAccountOverviewPage)
	engine.GET("/account/admin/overview", accountController.ServeAdminAccountOverviewPage)
	engine.GET("/logout", accountController.HandleLogout)
	engine.GET("/payment", accountController.HandlePayment)

	// engine.GET("/template", accountController.ServeTemplate)

}
