package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	Username  string
	Password  string
	Firstname sql.NullString
	Lastname  sql.NullString
	Email     sql.NullString
	Address   sql.NullString
	Country   sql.NullString
	City      sql.NullString
	Age       sql.NullString
	Phone     sql.NullString
}

func (usr User) AuthencateUser(user *User) bool {

	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()
	results, err := db.Query("select username,firstname from users where password=? and username=?", user.Password, user.Username)

	if err != nil {
		fmt.Printf("ERORR QUERYING: %v", err)
	}

	defer results.Close()

	user.Username = ""

	for results.Next() {

		err = results.Scan(&user.Username, &user.Firstname)
		if err != nil {
			fmt.Printf("Database Scan ERROR:%v", err)
		}

		if len(user.Username) >= 0 {
			fmt.Printf("\nAuthenticated User\nUsername:%s\nFirstName:%s\n\n", user.Username, user.Firstname.String)
			return true
		} else {
			return false
		}

	}
	return false
}

func (user User) GetAllUserData(username string) (User, error) {

	// Username  string
	// Password  string
	// Firstname string
	// Lastname  string
	// Email     string
	// Address   string
	// Country   string
	// City      string
	// Age       uint
	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()
	results, err := db.Query("select username,age,firstname,lastname,email,address,country,city from users where username=?", "admin")

	if err != nil {
		return user, err
	}

	defer results.Close()

	user.Username = ""

	for results.Next() {

		err = results.Scan(&user.Username, &user.Age, &user.Firstname, &user.Lastname, &user.Email, &user.Address, &user.Country, &user.City)
		// fmt.Printf("ALL DATA:%v", user.Username)

		if err != nil {

			fmt.Printf("scan error: %v", err)
			return user, err
		}
	}
	return user, nil
}
