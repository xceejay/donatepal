package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/xceejay/boilerplate/cache"
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

func (transaction Transaction) NewTransaction() *Transaction {

	return &Transaction{}
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

	start := time.Now()

	redisTransactions, err := transaction.GetAllTransactionsByFundRaiserWithCache(fundraiser)

	if err != nil {
		fmt.Println("GETTING WITH CACHE ERROR:", err)
	}
	// fmt.Println("long innit", len(redisTransactions))
	if len(redisTransactions) > 0 {
		// fmt.Println("WHY THE FUCK IS THERE NO DATA", redisTransactions)
		elapsed := time.Since(start)
		log.Printf("Search took %s using redis\n", elapsed)
		return redisTransactions, nil
	}

	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM transactions").Scan(&count)
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

	transactionsCache := new(cache.TransactionsCache)

	go func(fundraiser string, transactions []Transaction, transactionsCache *cache.TransactionsCache) {
		transactionsCache.FundRaiser = fundraiser

		// 	Email         string    `json:"Email"`
		// DonationType  string    `json:"Donation_Type"`
		// PaymentMethod string    `json:"PaymentMethod"`
		// DateCreated   time.Time `json:"DateCreated"`
		// Transactionid int       `json:"Transactionid"`
		// Firstname     string    `json:"Firstname"`
		// Lastname      string    `json:"Lastname"`
		// Amount        float64   `json:"Amount"`
		// Address       string    `json:"Address"`
		// Phone         string    `json:"Phone"`
		// DateDonated   time.Time `json:"DateDonated"`

		var index int = 0
		transactionsCache.Transactions = make([]cache.TransactionCache, len(transactions))
		for _, transaction := range transactions {
			transactionsCache.Transactions[index].Transactionid = transaction.Transactionid
			transactionsCache.Transactions[index].Email = transaction.Email.String
			transactionsCache.Transactions[index].Firstname = transaction.Firstname.String
			transactionsCache.Transactions[index].Lastname = transaction.Lastname.String
			transactionsCache.Transactions[index].Amount = transaction.Amount
			transactionsCache.Transactions[index].PaymentMethod = transaction.PaymentMethod
			transactionsCache.Transactions[index].Phone = transaction.Phone.String
			transactionsCache.Transactions[index].Address = transaction.Address.String
			transactionsCache.Transactions[index].DateDonated = transaction.DateDonated

			index++

		}

		transactionsCache.SetTransactionsByFundRaiser()
	}(fundraiser, transactions, transactionsCache)
	// transactionCache.SetTransactionsByFundRaiser(fundraiser, transactions)

	fmt.Println("Sucessfully Got Transactions By FundRaiser")
	elapsed := time.Since(start)
	fmt.Printf("Search took %s using sql \n", elapsed)
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
		results, err := db.Query("select sum(amount) from transactions where monthname(date_donated)=? and fundraiser=?", month, fundraiser)
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

func (transaction Transaction) GetAllTransactionsByFundRaiserWithCache(fundraiser string) ([]Transaction, error) {

	transactionsCache := new(cache.TransactionsCache)

	// Email         string    `json:"Email"`
	// DonationType  string    `json:"Donation_Type"`
	// PaymentMethod string    `json:"PaymentMethod"`
	// DateCreated   time.Time `json:"DateCreated"`
	// Transactionid int       `json:"Transactionid"`
	// Firstname     string    `json:"Firstname"`
	// Lastname      string    `json:"Lastname"`
	// Amount        float64   `json:"Amount"`
	// Address       string    `json:"Address"`
	// Phone         string    `json:"Phone"`
	// DateDonated   time.Time `json:"DateDonated"`

	transactionsCache.FundRaiser = fundraiser

	transactionscache, err := transactionsCache.GetAllTransactionsByFundRaiser(fundraiser)

	if err != nil {
		return nil, err
	}
	transactions := make([]Transaction, len(transactionscache.Transactions))

	var index int = 0

	for range transactionscache.Transactions {
		transactions[index].Email.String = transactionscache.Transactions[index].Email
		// fmt.Println("EVERYTIME:", transactions[index].Email.String)
		transactions[index].DonationType = transactionscache.Transactions[index].DonationType
		transactions[index].PaymentMethod = transactionscache.Transactions[index].PaymentMethod
		transactions[index].DateCreated = transactionscache.Transactions[index].DateCreated
		transactions[index].Transactionid = transactionscache.Transactions[index].Transactionid
		transactions[index].Firstname.String = transactionscache.Transactions[index].Firstname
		transactions[index].Lastname.String = transactionscache.Transactions[index].Lastname
		transactions[index].Address.String = transactionscache.Transactions[index].Address
		transactions[index].Amount = transactionscache.Transactions[index].Amount
		transactions[index].Phone.String = transactionscache.Transactions[index].Phone
		transactions[index].DateDonated = transactionscache.Transactions[index].DateDonated

		index++

	}

	// transactions := make([]Transaction, count)
	fmt.Println("-----------------------------used cache ------------------------")
	// fmt.Println("YES I GOT IT :", transactionscache.Transactions[0].Firstname)
	return transactions, nil
}
