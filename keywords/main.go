package main

import (
	"fmt"
	format "fmt"
	"log"

	. "github.com/dgkg/gomasterclasses/keywords/service"
)

func init() {
	// TODO: add config file.
}

type AccessLevel uint8

func (a AccessLevel) String() string {
	return tblAccessLevel[a]
}

const (
	SuperAdmin AccessLevel = iota + 1
	Admin
)

var tblAccessLevel = [...]string{
	SuperAdmin: "super-admin",
	Admin:      "admin",
}

const (
	PI        = 3.141595
	NbMaxCall = 100
)

func main() {
	// TODO exec program.
	format.Println("Hello")
	res, err := Run("Server")
	if err != nil {
		log.Println(err)
	}
	format.Println(res)
	format.Printf("%v - %T\n", SuperAdmin, SuperAdmin)
	format.Printf("%v - %T\n", Admin, Admin)

	format.Printf("%v - %T\n", NbMaxCall, NbMaxCall)
	var i int16 = 101
	if i > NbMaxCall {
		log.Println("max call !!!")
	}
	var j int64 = 1000
	if j > NbMaxCall {
		log.Println("max call !!!")
	}

	listUser := map[int]*User{
		0: &User{
			"509b6424-567e-11ec-8dbf-27dfb4b91470",
			"Bob",
			"L'Eponge",
		},
		3: &User{
			"71b6ccf2-567e-11ec-b0b4-078b2459bca4",
			"Robert",
			"Miles",
		},
	}

	delete(listUser, 3)

	for i := 0; i < len(listUser); i++ {
		u := listUser[i]
		if u == nil {
			continue
		}
		log.Printf("%#v", u)
	}

	for k := range listUser {
		log.Printf("%#v", listUser[k])
	}

	tblUser := []User{
		User{
			"509b6424-567e-11ec-8dbf-27dfb4b91470",
			"Bob",
			"L'Eponge",
		},
		User{
			"71b6ccf2-567e-11ec-b0b4-078b2459bca4",
			"Robert",
			"Miles",
		},
	}
	fmt.Println("tblUser", tblUser[len(tblUser)-1:])
	tblUser = append(tblUser, User{
		UUID:      "924c7680-5681-11ec-b0be-7b88d45fe63d",
		FirstName: "Rene",
		LastName:  "French",
	})
	fmt.Println("tblUser new", tblUser)
}

type User struct {
	UUID      string
	FirstName string
	LastName  string
}
