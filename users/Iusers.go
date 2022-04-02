package users

type IUser interface{
	GetUser(username string) string
	GetUsers() [] string
	CreateUser(user User)
}