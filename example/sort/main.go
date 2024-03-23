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
	sliceop.SortAsc(
		users,
		func(u User) string {
			return u.Name
		},
	)
	fmt.Println(users) // [{Alex 1} {Alex 22} {Alex 3} {Vlad 241} {Vlad 2}]

	sliceop.SortAsc(
		users,
		func(u User) int {
			return u.Score
		},
	)
	fmt.Println(users) // [{Alex 1} {Vlad 2} {Alex 3} {Alex 22} {Vlad 241}]
}
