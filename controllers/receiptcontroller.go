package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xceejay/donatepal/models"
	"github.com/xceejay/donatepal/services"
)

type ReceiptController struct{}

func (receiptController ReceiptController) HandleSaveReceipt(c *gin.Context) {
	// fmt.Printf("receipt handled")
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
		receiptModel.Fundraiser = c.PostForm("fundraiser")

		err = receiptModel.InsertReceipt()
		if err != nil {

			//change to better error page
			c.JSON(http.StatusExpectationFailed, gin.H{"message": "failed", "error": fmt.Sprintf("\n%v", err)})
		}

		// c.JSON(http.StatusOK, gin.H{"message": "sucess", "error": fmt.Sprintf("\n%v", err)})
		path := "views/html/account/receipt-pdf.html"

		vars := make(map[string]interface{})

		vars["receipt_id"] = "2021DR"
		vars["date_donated"] = receiptModel.DateDonated.Format("2020/12/01")
		vars["fundraiser_name"] = receiptModel.GetFundraisername()
		vars["amount_donated"] = receiptModel.Amount
		templateEngine := new(services.TemplateEngine)

		htmlContent := templateEngine.ProcessFile(path, vars)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.HTML(htmlContent)))

	}

}

func GetAllReceipts() {

}
func GetReceiptsByEmail() {

}

func GetReceiptsByDateCreated() {

}
