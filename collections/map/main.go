package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
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

var list map[string]*User

func main() {
	u, err := Add("Bob", "Pike")
	if err != nil {
		log.Println(err)
	}
	Update(u.UUID, "Rene", "French")
	u2, _ := Get(u.UUID)
	fmt.Printf("user 2 is: %#v", u2)
	fmt.Printf("%#v", list)
	fmt.Println("before delete")
	Delete(u.UUID)
	fmt.Println("after delete")
	fmt.Printf("%#v", list)
}

func Add(fn, ln string) (*User, error) {
	u := NewUser(fn, ln)
	_, ok := list[u.UUID]
	if ok {
		return nil, errors.New("map: user not found")
	}
	list[u.UUID] = u
	return u, nil
}

func Delete(uuid string) error {
	_, ok := list[uuid]
	if !ok {
		return errors.New("map: user not found")
	}
	delete(list, uuid)
	return nil
}

func Update(uuid, fn, ln string) (*User, error) {
	u, ok := list[uuid]
	if !ok {
		return nil, errors.New("map: user not found")
	}
	u.FirstName = fn
	u.LastName = ln
	return u, nil
}

func Get(uuid string) (*User, error) {
	u, ok := list[uuid]
	if !ok {
		return nil, errors.New("map: user not found")
	}
	return u, nil
}
