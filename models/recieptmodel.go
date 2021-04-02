package models

import (
	"database/sql"
	"fmt"
	"log"
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

func (receipt Receipt) GetAllReceipts() ([]Receipt, error) {
	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM receipts").Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Number of rows are %d\n", count)
	}

	receipts := make([]Receipt, count)
	results, err := db.Query("select reciept_id,email,firstname,lastname,amount,payment_method,phone,address,date_donated from receipts")

	if err != nil {
		return receipts, err
	}

	var i int = 0
	for results.Next() {

		err = results.Scan(&receipts[i].Receiptid, &receipts[i].Email, &receipts[i].Firstname.String, &receipts[i].Lastname.String, &receipts[i].Amount, &receipts[i].PaymentMethod, &receipts[i].Phone.String, &receipts[i].Address.String, &receipts[i].DateDonated)

		if err != nil {
			return receipts, err
		}
		i++

	}

	// fmt.Printf("%v", receipts)

	fmt.Println("Sucessfully Got Receipts")
	return receipts, nil
}
