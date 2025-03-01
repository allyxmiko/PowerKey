package utils

import "testing"

func TestCreateJwt(t *testing.T) {
	subject := "powerkey"
	token := CreateJwt(subject)
	if len(token) == 0 {
		t.Errorf("CreateJwt() failed")
	}
}

func TestParseJwt(t *testing.T) {
	jwtString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwb3dlcmtleSIsInN1YiI6InBvd2Vya2V5IiwiZXhwIjoxNzQwNzg2NjM5LCJuYmYiOjE3NDA3ODY2MjksImlhdCI6MTc0MDc4NjYyOSwianRpIjoiMTgzYTZmMzgtMDRmMC00OGZiLWExZTYtMmZjNzQ1ZTRiY2VkIn0.f8uZ0WdQtK95_Hyz5boRKgWdphiOBi5A6HUg3OVGn7M"
	token, err := ParseJwt(jwtString)
	if err != nil {
		t.Errorf("ParseJwt() failed: %s", err)
	}
	if token != "powerkey" {
		t.Errorf("ParseJwt() failed: %s", err)
	}
}
