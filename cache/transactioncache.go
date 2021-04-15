package cache

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type TransactionCache struct {
	Email         string    `json:"Email"`
	DonationType  string    `json:"DonationType"`
	PaymentMethod string    `json:"PaymentMethod"`
	DateCreated   time.Time `json:"DateCreated"`
	Transactionid int       `json:"Transactionid"`
	Firstname     string    `json:"Firstname"`
	Lastname      string    `json:"Lastname"`
	Amount        float64   `json:"Amount"`
	Address       string    `json:"Address"`
	Phone         string    `json:"Phone"`
	DateDonated   time.Time `json:"DateDonated"`
}

type TransactionsCache struct {
	Transactions []TransactionCache
	FundRaiser   string
}

// {
// 	"transaction_id": { "Email": "jack", "Donation_Type": "anoymous" }
// }

func (transactions TransactionsCache) SetTransactionsByFundRaiser() error {
	b, err := json.Marshal(transactions)
	if err != nil {
		fmt.Println("ERROR MARSHARLLING DATA:", err)

	}

	var redisCache RedisCache

	redisClient, err := redisCache.InitCache()
	if err != nil {
		return err
	}

	defer redisClient.Close()

	//the 60 is seconds

	_, err = redisClient.Do("SETEX", transactions.FundRaiser+"_transactions", 60, string(b))

	if err != nil {
		return err
	}
	// transactions.GetAllTransactionsByFundRaiser(transactions.FundRaiser)
	return nil
}

func (transactionscache TransactionsCache) GetAllTransactionsByFundRaiser(fundraiser string) (TransactionsCache, error) {

	var redisCache RedisCache

	redisClient, err := redisCache.InitCache()

	if err != nil {
		return TransactionsCache{}, err
	}
	defer redisClient.Close()
	result, err := redis.String(redisClient.Do("GET", fundraiser+"_transactions"))

	if err != nil {
		return TransactionsCache{}, err
	}

	err = json.Unmarshal([]byte(result), &transactionscache)
	if err != nil {
		return TransactionsCache{}, err

	}
	// fmt.Printf("Length of array: %v\n", len(transactionscache.Transactions))

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("I LOVE YOU SO MUCH", transactionscache.Transactions)
	// }
	return transactionscache, nil
}
