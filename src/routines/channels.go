package routines

import "math/rand"

type User struct {
	name string
	age  int
}

func getUser(name string) User {
	return User{
		name: name,
		age:  rand.Intn(100),
	}
}

