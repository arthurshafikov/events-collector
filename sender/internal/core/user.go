package core

type User struct {
	Name     string
	Username string
	Password string
}

type AuthorizeInput struct {
	Username string
	Password string
}
