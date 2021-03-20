package user

type User_Interface interface {
	
	RegisterUser(user User) (int, error)
}
