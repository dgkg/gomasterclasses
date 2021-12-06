package main

import (
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
	format.Println(SuperAdmin)
	format.Println(Admin)

	format.Printf("%v - %T", NbMaxCall, NbMaxCall)
	var i int16 = 101
	if i > NbMaxCall {
		format.Println("max call !!!")
	}
	var j int64 = 1000
	if j > NbMaxCall {
		format.Println("max call !!!")
	}
}
