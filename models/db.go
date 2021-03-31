package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "1234"
	hostname = "localhost:3306"
	dbname   = "donatepal"
)

type Database struct {
}

///EXPORTS

//INTERNAL
func (database Database) InitDatabase() *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	return db

}

// func (database Database) DBTest() {

// 	var user User
// 	db := database.InitDatabase()
// 	defer db.Close()
// 	results, err := db.Query("select * from users")
// 	if err != nil {
// 		panic(err)
// 	}

// 	for results.Next() {

// 		err = results.Scan(&user.username, &user.age)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("username:%s\nage: %v\n\n", user.username, user.age)
// 	}
// }
