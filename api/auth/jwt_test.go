package auth

import "testing"

func TestJWTSign(t *testing.T) {
	uuid, name := "001d0684-58da-11ec-b7b2-1f88738963de", "Bob"
	jwtValue, err := JWTSign(uuid, name)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(jwtValue)
	data, err := JWTParse(jwtValue)
	if err != nil {
		t.Error(err)
		return
	}
	if data["uuid"] != uuid || data["full_name"] != name {
		t.Errorf("exected uuid %v got uuid %v and\nexpected %v and got %v", uuid, data["uuid"], name, data["full_name"])
		return
	}
}
