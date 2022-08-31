package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xceejay/donate-pal/controllers"
)

func InitRouter(engine *gin.Engine) {

	accountController := new(controllers.AccountController)
	homeController := new(controllers.HomeController)
	receiptController := new(controllers.ReceiptController)
	transactionController := new(controllers.TransactionController)

	engine.NoRoute(homeController.Handle404)
	engine.GET("/", homeController.ServeHomePage)
	engine.POST("/", homeController.ServeHomePage)

	engine.GET("/about", homeController.ServeAboutPage)
	engine.GET("/donate", homeController.ServeDonationPage)
	engine.GET("/register", homeController.ServeRegistrationPage)
	engine.GET("/login-page", accountController.ServeLoginPage)
	engine.POST("/login-page", accountController.ServeLoginPage)
	engine.GET("/login", accountController.HandleLogin)
	engine.POST("/login", accountController.PerformLogin)
	// engine.POST("/login", accountController.PerformLogin)
	// engine.POST("/account/:username", accountController.ServeAccountPage)
	// engine.GET("/account/:username", accountController.HandleAccountPage)
	engine.POST("/logout", accountController.HandleLogout)

	engine.POST("/register", accountController.HandleRegistration)
	engine.POST("/account/admin", accountController.ServeAdminAccountPage)
	engine.POST("/successful-registration", accountController.HandleSuccessfulRegistration)
	engine.GET("/account/admin", accountController.HandleAdminAccountPage)
	engine.GET("/account/admin/:dashboard_content", accountController.HandleAdminDashboardContent)
	engine.POST("/account/admin/:dashboard_content/*savereceipt", receiptController.HandleSaveReceipt)
	engine.POST("/payment", transactionController.HandlePayment)
	engine.POST("/card-payment", transactionController.HandleCardPayment)
	engine.POST("/donate", transactionController.HandleDonation)
	engine.POST("/successful-payment", transactionController.HandleSuccessfulPayment)

	// engine.GET("/template", accountController.ServeTemplate)

}
