package database

type Database_Interface interface {
	AddUser(email string) bool
	FetchUserID(email string) int
	UserExists(email string) bool
}
