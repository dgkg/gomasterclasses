package service

import (
	"errors"
	"fmt"
)

func Run(str string) (string, error) {
	fmt.Println(str)
	return str, errors.New("port allready used")
}
