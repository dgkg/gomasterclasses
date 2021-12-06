/*
break - ok
default - ok
func - ok
interface - ok
select - ok
case -ok
defer - ok
go - ok
map - ok
struct - ok
chan - ok
else - ok
goto -
package - ok
switch -
const - ok
fallthrough -
if - ok
range - ok
type - ok
continue - ok
for - ok
import - ok
return - ok
var - ok
*/
package main

import (
	format "fmt"
	"log"
	"time"

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
			SuperAdmin,
		},
		3: &User{
			"71b6ccf2-567e-11ec-b0b4-078b2459bca4",
			"Robert",
			"Miles",
			Admin,
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
		if listUser[k].AccessLevel == Admin {
			format.Println(listUser[k].FirstName, " is ", Admin)
		} else if listUser[k].UUID == "71b6ccf2-567e-11ec-b0b4-078b2459bca4" {
			format.Println(listUser[k].FirstName, " is ", listUser[k].AccessLevel)
		}
	}

	tblUser := []User{
		User{
			"509b6424-567e-11ec-8dbf-27dfb4b91470",
			"Bob",
			"L'Eponge",
			Admin,
		},
		User{
			"71b6ccf2-567e-11ec-b0b4-078b2459bca4",
			"Robert",
			"Miles",
			SuperAdmin,
		},
	}
	format.Println("tblUser", tblUser[len(tblUser)-1:])
	tblUser = append(tblUser, User{
		UUID:        "924c7680-5681-11ec-b0be-7b88d45fe63d",
		FirstName:   "Rene",
		LastName:    "French",
		AccessLevel: SuperAdmin,
	})
	format.Println("tblUser new", tblUser)
	execSomethingHeavy()

	tblAll := []interface{}{
		tblUser[:],
		Animal{
			Name: "Mistigris",
		},
	}
	_ = tblAll

	for _, v := range tblAll {
		switch v.(type) {
		case User, *User:
			format.Println("this is a Human")
		case []User:
			format.Println("this is a group of Humans")
		case Animal, *Animal:
			format.Println("this is an Animal")
		default:
			format.Println("we don't know what it is")
		}
	}
}

type User struct {
	UUID        string
	FirstName   string
	LastName    string
	AccessLevel AccessLevel
}

type Animal struct {
	UUID string
	Name string
}

func newTicker(sec uint) *time.Ticker {
	return time.NewTicker(time.Duration(sec) * time.Second)
}

func doSomethingHeavy(todo chan string) {
	defer format.Println("finished")
	time.Sleep(2 * time.Second)
	todo <- "done"
}

func execSomethingHeavy() {
	ticker := newTicker(10)
	todo := make(chan string)
	go doSomethingHeavy(todo)
loop:
	for {
		select {
		case <-ticker.C:
			format.Println("ticker has tick")
			break loop
		case res := <-todo:
			format.Println("res is:", res)
			return
		default:
			//format.Print("wait for result...")
		}
	}
}
