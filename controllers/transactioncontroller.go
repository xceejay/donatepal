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
	session.Set("", c.PostForm(""))
	session.Set("", c.PostForm(""))
	session.Set("", c.PostForm(""))
	session.Set("", c.PostForm(""))
	session.Set("", c.PostForm(""))
	session.Set("", c.PostForm(""))

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, "/payment")
}

func (transactionController TransactionController) HandlePayment(c *gin.Context) {
	// session := sessions.Default(c)
	// var tempvar1 string
	// tempvar := session.Get(tempvar1)
	// tempvar = session.Get(tempvar1)
	// tempvar = session.Get(tempvar1)
	// tempvar = session.Get(tempvar1)
	// tempvar = session.Get(tempvar1)
	// tempvar = session.Get(tempvar1)
	// tempvar = session.Get(tempvar1)
	// tempvar = session.Get(tempvar1)

	c.HTML(http.StatusOK, "success.html", nil)

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
