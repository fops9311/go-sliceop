package main

import (
	"fmt"
	"fops9311/go-sliceop"
)

func main() {
	type User struct {
		Name  string
		Score int
	}
	users := []User{
		{
			Name:  "Alex",
			Score: 1,
		},
		{
			Name:  "Vlad",
			Score: 241,
		},
		{
			Name:  "Alex",
			Score: 22,
		},
		{
			Name:  "Vlad",
			Score: 2,
		},
		{
			Name:  "Alex",
			Score: 3,
		},
	}
	users = sliceop.Group(
		users,
		func(u User) string {
			return u.Name
		},
		func(u1 User, u2 User) User {
			u1.Score += u2.Score
			return u1
		},
	)
	fmt.Println(users) //[{Alex 26} {Vlad 243}]
}
