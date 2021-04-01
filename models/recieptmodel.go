package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Receipt struct {
	Email         sql.NullString
	DonationType  string
	PaymentMethod string
	DateCreated   time.Time
	Receiptid     int
	Firstname     sql.NullString
	Lastname      sql.NullString
	Amount        float64
	Address       sql.NullString
	Phone         sql.NullString
	DateDonated   time.Time
}

func (receipt Receipt) InsertReceipt() error {
	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	statement, err := db.Prepare("insert into receipts (email,payment_method,date_created,firstname,lastname,amount,address,phone,date_donated) values(?,?,curdate(),?,?,?,?,?,?)")

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(receipt.Email.String, receipt.PaymentMethod, receipt.Firstname.String, receipt.Lastname.String, receipt.Amount, receipt.Address.String, receipt.Phone.String, receipt.DateDonated)
	if err != nil {
		return err
	}
	fmt.Println("Sucessfully Inserted Receipt")
	return nil
}
