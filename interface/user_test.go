package main

import (
	"encoding/json"
	"testing"
)

func TestPasswordUnmarshalJSON(t *testing.T) {
	payload := []byte(`{
		"first_name":"Jane",
		"last_name":"Doe",
		"pass":"123password"}`)
	var u User
	err := json.Unmarshal(payload, &u)
	if err != nil {
		t.Errorf("got error %v", err)
		return
	}
	t.Log("user:", u)
	if u.Password != "47625ed74cab8fbc0a8348f3df1feb07f87601e34d62bd12eb0d51616566fab5" {
		t.Errorf("password not valid waiting for 123password got %v", u.Password)
	}
}
