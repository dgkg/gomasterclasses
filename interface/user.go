package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	var u User = User{
		UUID:      "alsdjflkjsd",
		FirstName: "Bob",
		LastName:  "L'Eponge",
		Password:  "password123",
		CreateAt:  time.Now(),
	}
	data, _ := json.Marshal(u)
	fmt.Println(string(data))
}

type User struct {
	UUID      string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
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

func (u User) MarshalJSON() ([]byte, error) {
	aux := map[string]string{
		"id":         u.UUID,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
	}
	return json.Marshal(aux)
}
