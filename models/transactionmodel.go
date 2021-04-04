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
	FundRaiser    string
	DateDonated   time.Time
}

func (transaction Transaction) InsertTransaction() error {
	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	statement, err := db.Prepare("insert into transactions (email,payment_method,firstname,lastname,amount,address,phone,date_donated,donation_type,fundraiser) values(?,?,?,?,?,?,?,curdate(),?,?)")

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(transaction.Email.String, transaction.PaymentMethod, transaction.Firstname.String, transaction.Lastname.String, transaction.Amount, transaction.Address.String, transaction.Phone.String, transaction.DonationType, transaction.FundRaiser)
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
		// fmt.Printf("Number of rows are %d\n", count)
	}

	transactions := make([]Transaction, count)
	results, err := db.Query("select transaction_id,email,firstname,lastname,amount,payment_method,phone,address,date_donated,donation_type,fundraiser from transactions")

	if err != nil {
		return transactions, err
	}

	var i int = 0
	for results.Next() {

		err = results.Scan(&transactions[i].Transactionid, &transactions[i].Email, &transactions[i].Firstname.String, &transactions[i].Lastname.String, &transactions[i].Amount, &transactions[i].PaymentMethod, &transactions[i].Phone.String, &transactions[i].Address.String, &transactions[i].DateDonated, &transactions[i].DonationType, &transactions[i].FundRaiser)

		if err != nil {
			return transactions, err
		}
		i++

	}

	// fmt.Printf("%v", transactions)

	fmt.Println("Sucessfully Got Transactions")
	return transactions, nil
}

func (transaction Transaction) GetAllTransactionsByFundRaiser(fundraiser string) ([]Transaction, error) {
	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM transactions").Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		// fmt.Printf("Number of rows are %d\n", count)
	}

	transactions := make([]Transaction, count)
	results, err := db.Query("select transaction_id,email,firstname,lastname,amount,payment_method,phone,address,date_donated,donation_type,fundraiser from transactions where fundraiser = ?", fundraiser)

	if err != nil {
		return transactions, err
	}

	var i int = 0
	for results.Next() {

		err = results.Scan(&transactions[i].Transactionid, &transactions[i].Email, &transactions[i].Firstname.String, &transactions[i].Lastname.String, &transactions[i].Amount, &transactions[i].PaymentMethod, &transactions[i].Phone.String, &transactions[i].Address.String, &transactions[i].DateDonated, &transactions[i].DonationType, &transactions[i].FundRaiser)

		if err != nil {
			return transactions, err
		}
		i++

	}

	// fmt.Printf("%v", transactions)

	fmt.Println("Sucessfully Got Transactions By FundRaiser")
	return transactions, nil
}

func (transaction Transaction) GetTotalAmountOfTransactionsByFundraiser(fundraiser string) int {
	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	var count int = 0
	err := db.QueryRow("SELECT COUNT(*) FROM transactions where fundraiser = ?", fundraiser).Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		// fmt.Printf("Number of rows are %d\n", count)
	}

	return count

}

// fmt.Printf("%v", transactions)

func (transaction Transaction) GetTotalAmountRaisedByFundaiser(fundraiser string) float64 {
	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	var Amount sql.NullFloat64
	err := db.QueryRow("SELECT SUM(amount) FROM transactions where fundraiser=?", fundraiser).Scan(&Amount)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		// fmt.Printf("Number of rows are %d\n", Amount)
	}
	return Amount.Float64
}

// get transaction by month and transacion amounts in array
// select * from transactions where monthname(date_donated)="April"

func (transaction Transaction) GetMonthlyTransactionAmountsByFundRaiser(fundraiser string) ([12]float64, error) {
	var transactionAmountsSql [12]sql.NullFloat64
	var transactionAmounts [12]float64

	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	months := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM transactions where fundraiser=?", fundraiser).Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		// fmt.Printf("Number of rows are %d\n", count)
	}

	// var results []*sql.Rows
	var i int = 0
	for _, month := range months {
		results, err := db.Query("select sum(amount) from transactions where monthname(date_donated)=?", month)
		if err != nil {
			return transactionAmounts, err
		}

		for results.Next() {
			err = results.Scan(&transactionAmountsSql[i])
			if err != nil {
				return transactionAmounts, err
			}

		}
		// fmt.Printf("AMOUNT FOR %s : %.2f\n", month, transactionAmountsSql[i].Float64)
		i++
	}

	for index, transactionAmountSql := range transactionAmountsSql {
		transactionAmounts[index] = transactionAmountSql.Float64

	}

	return transactionAmounts, nil
}
