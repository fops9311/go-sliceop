# go-sliceop
Обобщенные операции над слайсами, без лишней аллокации

## Filter

```go
package main

import (
	"fmt"
	sliceop "fops9311/go-sliceop"
)

func main() {
	s := []string{"123", "123", "321", "1515", "123"}

	filter := func(s string) bool {
		return s == ""
	}
	sliceop.Filter(s, filter)

	fmt.Println(s) // [123 123 123]
}
```

## Group
```go
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
```

## Sort

Обобщенная сортировка.

```go
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
```
