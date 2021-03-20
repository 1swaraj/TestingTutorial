package database

import (
	"fmt"
)

// db act as a dummy database.
var db map[string]int

// init initialize a the db with some data
func init() {
	db = make(map[string]int)
	// we assign a userid corresponding to the email
	db["1swarajphadtare@gmail.com"] = len(db)+1
	db["xyz@gmail.com"] = len(db)+1
}

// AddUser adds the user (email) to the database
func AddUser(email string) bool {
	fmt.Println("AddUser called in database package")
	if UserExists(email) {
		return false
	}
	db[email]=len(db)+1
	return true
}

// UserExists check if the User is registered with the provided email.
func UserExists(email string) bool {
	fmt.Println("UserExists called in database package")
	if _, ok := db[email]; !ok {
		return false
	}
	return true
}

// FetchUserID returns the userid corresponding to the email of the user
func FetchUserID(email string) int {
	fmt.Println("FetchUserID called in database package")
	return db[email]
}