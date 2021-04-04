package models

import (
	"database/sql"
	"fmt"
	"log"
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

func (user User) GetAllUserData() ([]User, error) {

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

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Number of rows are %d\n", count)
	}

	users := make([]User, count)

	results, err := db.Query("select username,age,firstname,lastname,email,address,country,city from users")

	if err != nil {
		return users, err
	}

	defer results.Close()

	user.Username = ""
	var i int = 0

	for results.Next() {

		err = results.Scan(&users[i].Username, &users[i].Age, &users[i].Firstname, &users[i].Lastname, &users[i].Email, &users[i].Address, &users[i].Country, &users[i].City)
		// fmt.Printf("ALL DATA:%v", users[i].users[i]name)

		if err != nil {

			fmt.Printf("scan error: %v", err)

			return users, err

		}
		i++
	}
	return users, nil
}

func (user User) GetAllUserDataByUsername(username string) (User, error) {

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
	results, err := db.Query("select username,age,firstname,lastname,email,address,country,city from users where username=?", username)

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

func (user User) InsertUser() error {
	database := new(Database)
	db := database.InitDatabase()
	defer db.Close()

	statement, err := db.Prepare("insert into users (username,password,email,firstname,lastname,address,country,city) values(?,?,?,?,?,?,?,?)")

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Username, user.Password, user.Email.String, user.Firstname.String, user.Lastname.String, user.Address.String, user.Country.String, user.City.String)
	if err != nil {
		return err
	}
	fmt.Println("Sucessfully Inserted User")
	return nil
}
