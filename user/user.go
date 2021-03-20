package user

import (
	"fmt"
	
	"github.com/swaraj1802/TestingTutorial/database"
)

type Impl struct {

}

func NewImpl() *Impl {
	return &Impl{}
}

// User encapsulate a user in the system.
type User struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID int    `json:"userid"`
}

// RegisterUser will register a User if only User.Email doesn't exist in our database
// If the user is successfully registered, it returns the userid else it returns -1 and the error
func RegisterUser(user User,db database.Database_Interface) (int, error) {
	adduser := db.AddUser(user.Email)
	if adduser {
		userid := db.FetchUserID(user.Email)
		return userid,nil
	} else {
		return -1,fmt.Errorf("user already exists")
	}
}