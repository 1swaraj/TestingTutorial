package user

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

)

type User_Interface_Mock struct {
	User_Interface
	mock.Mock
}

func (a *User_Interface_Mock) UserExists(email string) bool {
	args := a.Called(email)
	return args.Get(0).(bool)
}

func TestRegisterUser(t *testing.T) {
	user := User{
		Name:  "demouser",
		Email: "demo@gmail.com",
	}
	usermock := new(User_Interface_Mock)
	usermock.On("UserExists", mock.Anything).Return(false)
	val, err := RegisterUser(user)
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
	
	_, err := RegisterUser(user)
	// Since this is a new user, we should not get any error
	assert.Nil(t,err)
	
	// Now we will try to register the another user
	// with email id johndoe@gmail.com
	user2 := User{
		Name:  "not john doe",
		Email: "johndoe@gmail.com",
	}
	_, err = RegisterUser(user2)
	
	// We should get an error since the user exists
	assert.NotNil(t, err)
	assert.Equal(t, "user already exists",err.Error())
	
}
