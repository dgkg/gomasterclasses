package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	UUID      string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  Password  `json:"pass"`
	CreateAt  time.Time `json:"create_at"`
}

type Password string

func (Password) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

func (p *Password) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	h := sha256.New()
	h.Write([]byte(s))
	*p = Password(fmt.Sprintf("%x", h.Sum(nil)))
	return nil
}

type LoginUser struct {
	Email    string   `json:"email"`
	Password Password `json:"pass"`
}
