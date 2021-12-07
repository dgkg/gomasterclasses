package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("slice: user not found")
)

type User struct {
	UUID      string
	FirstName string
	LastName  string
	CreateAt  time.Time
}

func NewUser(fn, ln string) *User {
	return &User{
		UUID:      uuid.NewString(),
		FirstName: fn,
		LastName:  ln,
		CreateAt:  time.Now(),
	}
}

var list []*User

func main() {
	// ADD
	u, err := Add("Bob", "Pike")
	if err != nil {
		panic(err)
	}

	// Update
	_, err = Update(u.UUID, "Rene", "French")
	if err != nil {
		panic(err)
	}

	// Get
	u2, err := Get(u.UUID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user 2 is: %#v\n", u2)

	// Delete
	fmt.Printf("list before delete: %#v\n", list[0])
	err = Delete(u.UUID)
	if err != nil {
		panic(err)
	}

	fmt.Println(findIndex(u.UUID))
	fmt.Printf("list after delete: %#v\n", list)

}

func findIndex(uuid string) (int, error) {
	for k := range list {
		if list[k] != nil && list[k].UUID == uuid {
			return k, nil
		}
	}
	return -1, ErrNotFound
}

func Add(fn, ln string) (*User, error) {
	u := NewUser(fn, ln)
	_, err := findIndex(u.UUID)
	if err != ErrNotFound {
		return nil, err
	}
	list = append(list, u)
	return u, nil
}

func Delete(uuid string) error {
	i, err := findIndex(uuid)
	if err != nil {
		return err
	}
	list = append(list[:i], list[i+1:]...)
	return nil
}

func Update(uuid, fn, ln string) (*User, error) {
	i, err := findIndex(uuid)
	if err != ErrNotFound {
		return nil, err
	}
	list[i].FirstName = fn
	list[i].LastName = ln
	return list[i], nil
}

func Get(uuid string) (*User, error) {
	i, err := findIndex(uuid)
	if err != ErrNotFound {
		return nil, err
	}
	return list[i], nil
}
