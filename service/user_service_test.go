package service

import (
	"PowerKey/db"
	"testing"
)

var userService IUserService

func init() {
	err := db.Init()
	if err != nil {
		return
	}
	userService = NewUserService()
}

func TestFindUserByUsername(t *testing.T) {
	user, err := userService.FindByUsername("admin")
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
		return
	}

	if user.Username != "admin" {
		t.Errorf("expected username '%s', but got '%s'", "admin", user.Username)
	}
}

func TestUpdatePassword(t *testing.T) {
	err := userService.UpdatePassword("admin", "newpassword")
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
		return
	}
}
