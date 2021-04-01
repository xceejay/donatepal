package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xceejay/boilerplate/models"
)

type ReceiptController struct{}

func (receiptController ReceiptController) HandleSaveReceipt(c *gin.Context) {
	fmt.Printf("receipt handled")
	if c.Param("dashboard_content") == "receipt" && isLoggedIn(c) {
		receiptModel := new(models.Receipt)

		receiptModel.Email.String = c.PostForm("email")
		receiptModel.Firstname.String = c.PostForm("firstname")
		receiptModel.Lastname.String = c.PostForm("lastname")

		amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
		if err != nil {
			fmt.Printf("ERROR CONV: %v", err)
		}
		receiptModel.Amount = amount
		donated_time, err := time.Parse("2006-01-02", c.PostForm("date_donated"))

		if err != nil {
			fmt.Printf("Error converting time: %v", err)
		}

		receiptModel.DateDonated = donated_time
		receiptModel.PaymentMethod = c.PostForm("payment_method")
		receiptModel.Address.String = c.PostForm("address")
		receiptModel.Phone.String = c.PostForm("phone")

		err = receiptModel.InsertReceipt()
		if err != nil {

			//change to better error page
			c.JSON(http.StatusExpectationFailed, gin.H{"message": "failed", "error": fmt.Sprintf("\n%v", err)})
		}

		// c.JSON(http.StatusOK, gin.H{"message": "sucess", "error": fmt.Sprintf("\n%v", err)})

		c.HTML(http.StatusOK, "receipt-pdf.html", nil)

	}

}

func GetAllReceipts() {

}
func GetReceiptsByEmail() {

}

func GetReceiptsByDateCreated() {

}
