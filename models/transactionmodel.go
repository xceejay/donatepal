package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Transaction struct {
	Email         sql.NullString
	DonationType  string
	PaymentMethod string
	DateCreated   time.Time
	Transactionid int
	Firstname     sql.NullString
	Lastname      sql.NullString
	Amount        float64
	Address       sql.NullString
	Phone         sql.NullString
	DateDonated   time.Time
}

func (transaction Transaction) InsertTransaction() error {
	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	statement, err := db.Prepare("insert into transactions (email,payment_method,firstname,lastname,amount,address,phone,date_donated,donation_type) values(?,?,?,?,?,?,?,curdate(),?)")

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(transaction.Email.String, transaction.PaymentMethod, transaction.Firstname.String, transaction.Lastname.String, transaction.Amount, transaction.Address.String, transaction.Phone.String, transaction.DonationType)
	if err != nil {
		return err
	}
	fmt.Println("Sucessfully Inserted Transaction")
	return nil
}

func (transaction Transaction) GetAllTransactions() ([]Transaction, error) {
	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM transactions").Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Number of rows are %d\n", count)
	}

	transactions := make([]Transaction, count)
	results, err := db.Query("select reciept_id,email,firstname,lastname,amount,payment_method,phone,address,date_donated,donation_type from transactions")

	if err != nil {
		return transactions, err
	}

	var i int = 0
	for results.Next() {

		err = results.Scan(&transactions[i].Transactionid, &transactions[i].Email, &transactions[i].Firstname.String, &transactions[i].Lastname.String, &transactions[i].Amount, &transactions[i].PaymentMethod, &transactions[i].Phone.String, &transactions[i].Address.String, &transactions[i].DateDonated, transaction.DonationType)

		if err != nil {
			return transactions, err
		}
		i++

	}

	// fmt.Printf("%v", transactions)

	fmt.Println("Sucessfully Got Transactions")
	return transactions, nil
}
