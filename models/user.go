package models

type User struct {
	id       string
	username string
	email    string
}

func  createUser(id, username, email string) (error, User) {
	newUser := &User{}
	return nil, *newUser
}

