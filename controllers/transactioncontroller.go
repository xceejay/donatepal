package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xceejay/boilerplate/models"
)

type TransactionController struct{}

func (transactionController TransactionController) HandleDonation(c *gin.Context) {
	session := sessions.Default(c)

	session.Set("payment_method", c.PostForm("payment_method"))

	session.Set("firstname", c.PostForm("firstname"))

	session.Set("lastname", c.PostForm("lastname"))
	session.Set("email", c.PostForm("email"))
	session.Set("address", c.PostForm("address"))
	session.Set("amount", c.PostForm("amount"))
	session.Set("date_donated", c.PostForm("date_donated"))
	session.Set("donation_type", c.PostForm("donation_type"))
	session.Set("fundraiser", c.PostForm("fundraiser"))

	session.Set("phone", c.PostForm("phone"))

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/payment")

}

func (transactionController TransactionController) HandleCardPayment(c *gin.Context) {
	session := sessions.Default(c)
	if c.PostForm("number") != "" {

		// this will mark the session as "written" and hopefully remove the username

		session.Set("payment_method", "")
		session.Set("firstname", "")
		session.Set("lastname", "")
		session.Set("email", "")
		session.Set("address", "")
		session.Set("amount", "")
		session.Set("date_donated", "")
		session.Set("donation_type", "")
		session.Set("fundraiser", "")

		session.Set("phone", "")
		session.Clear()
		session.Options(sessions.Options{Path: "/", MaxAge: -1}) // this sets the cookie with a MaxAge of 0

		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}
		c.Redirect(http.StatusPermanentRedirect, "/successful-payment")
		return
	}
	transactionController.ServePaymentPage(c)
}

func (transactionController TransactionController) HandleSuccessfulPayment(c *gin.Context) {
	c.HTML(http.StatusOK, "success.html", nil)
}

func (transactionController TransactionController) HandlePayment(c *gin.Context) {

	// fmt.Println("payment handled")

	transactionModel := new(models.Transaction)

	homeController := new(HomeController)

	session := sessions.Default(c)

	transactionModel.Firstname.String = fmt.Sprintf("%s", session.Get("firstname"))
	transactionModel.Lastname.String = fmt.Sprintf("%s", session.Get("lastname"))
	transactionModel.Email.String = fmt.Sprintf("%s", session.Get("email"))
	transactionModel.Address.String = fmt.Sprintf("%s", session.Get("address"))
	transactionModel.Phone.String = fmt.Sprintf("%s", session.Get("phone"))
	transactionModel.DonationType = fmt.Sprintf("%s", session.Get("donation_type"))
	transactionModel.FundRaiser = fmt.Sprintf("%s", session.Get("fundraiser"))

	amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err != nil {
		fmt.Printf("ERROR CONV: %v", err)
	}
	transactionModel.Amount = amount
	// donated_time, err := time.Now().UTC()

	// if err != nil {

	// 	fmt.Printf("Error converting time: %v", err)
	// }
	// transactionModel.Amount = amount
	transactionModel.DateDonated = time.Now()
	transactionModel.PaymentMethod = fmt.Sprintf("%s", session.Get("payment_method"))
	// if tempvar empty

	if session.Get("amount") == "" {
		homeController.ServeDonationPage(c)
		return
	}

	err = transactionModel.InsertTransaction()
	if err != nil {
		fmt.Printf("ERROR Inserting transactions: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction"})

	}
	fmt.Println(transactionModel.PaymentMethod)
	if transactionModel.PaymentMethod == "card" {
		// fmt.Println("its card")
		transactionController.HandleCardPayment(c)
	}

}

func (transactionController TransactionController) HandleSaveTransaction(c *gin.Context) {
	// fmt.Printf("transaction handled")
	if c.Param("dashboard_content") == "transaction" && isLoggedIn(c) {
		transactionModel := new(models.Transaction)

		transactionModel.Email.String = c.PostForm("email")
		transactionModel.Firstname.String = c.PostForm("firstname")
		transactionModel.Lastname.String = c.PostForm("lastname")

		amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
		if err != nil {
			fmt.Printf("ERROR CONV: %v", err)
		}
		transactionModel.Amount = amount
		donated_time, err := time.Parse("2006-01-02", c.PostForm("date_donated"))

		if err != nil {
			fmt.Printf("Error converting time: %v", err)
		}

		transactionModel.DateDonated = donated_time
		transactionModel.PaymentMethod = c.PostForm("payment_method")
		transactionModel.FundRaiser = c.PostForm("fundraiser")

		transactionModel.Address.String = c.PostForm("address")
		transactionModel.Phone.String = c.PostForm("phone")

		err = transactionModel.InsertTransaction()
		if err != nil {

			//change to better error page
			c.JSON(http.StatusExpectationFailed, gin.H{"message": "failed", "error": fmt.Sprintf("\n%v", err)})
		}

		// c.JSON(http.StatusOK, gin.H{"message": "sucess", "error": fmt.Sprintf("\n%v", err)})

		c.HTML(http.StatusOK, "transaction-pdf.html", nil)

	}

}

func (transactionController TransactionController) ServePaymentPage(c *gin.Context) {

	c.HTML(http.StatusOK, "payment.html", nil)
}
