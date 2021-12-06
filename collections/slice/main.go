package main

import (
	"errors"
	"fmt"
	"log"
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

	var tbl []uint8
	var capTbl int = cap(tbl)
	for i := 0; i < 200; i++ {
		tbl = append(tbl, uint8(i))
		if cap(tbl) != capTbl {
			log.Println("new capacity", capTbl)
			capTbl = cap(tbl)
		}
	}
	for i := 0; i < len(tbl); i++ {
		log.Println(&tbl[i])
	}
	fmt.Printf("tbl pointer value:%p\n", tbl)
	DisplayPointer(tbl)
}

func DisplayPointer(tbl []uint8) {
	fmt.Printf("tbl pointer value:%p\n", tbl)
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
	if err != ErrNotFound {
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
