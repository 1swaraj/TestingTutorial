package user

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/swaraj1802/TestingTutorial/database"
)

type Database_Interface_Mock struct {
	database.Database_Interface
	mock.Mock
}

func (a *Database_Interface_Mock) AddUser(email string) bool {
	args := a.Called(email)
	return args.Get(0).(bool)
}

func (a *Database_Interface_Mock) FetchUserID(email string) int {
	args := a.Called(email)
	return args.Get(0).(int)
}

func TestRegisterUser(t *testing.T) {
	user := User{
		Name:  "demouser",
		Email: "demo@gmail.com",
	}
	dbmock := new(Database_Interface_Mock)
	dbmock.On("AddUser", mock.Anything).Return(true)
	dbmock.On("FetchUserID", mock.Anything).Return(3)
	val, err := RegisterUser(user,dbmock)
	// Since this is a new user we expect that no error should occur
	assert.Nil(t,err)
	// We assert that the user id allocated to the user is 3
	// (because we had initialized two users in the init method)
	assert.Equal(t, 3,val)
}

func TestRegisterUser_UserExists(t *testing.T) {
	user := User{
		Name:  "john doe",
		Email: "johndoe@gmail.com",
	}
	dbmock := new(Database_Interface_Mock)
	// We consider that this user is already created so we return false
	// from the mock implementation of user
	dbmock.On("AddUser", mock.Anything).Return(false)
	_, err := RegisterUser(user,dbmock)
	// Because the user exists, we must get an error
	assert.NotNil(t, err)
	assert.Equal(t, "user already exists",err.Error())
}
