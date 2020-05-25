package test

import (
	"github.com/solrac97gr/yendoapi/utilities"
	"testing"

	"github.com/solrac97gr/yendoapi/database"
)

func TestDbConnection(t *testing.T) {
	expected := true
	isConnected := database.ConnectionOK()
	if expected != isConnected {
		t.Error("The database isn't conected")
	}
}

func TestProcessToken(t *testing.T) {
	token := "BearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI1ZWFkYTRiN2VlMzI3ODU1NWI5NTgzYzIiLCJhdmF0YXIiOiIiLCJiaXJ0aGRhdGUiOiIyMDA2LTAxLTAyVDE1OjA0OjA1WiIsImVtYWlsIjoicHJ1ZWJhQGdtYWlsLmNvbSIsImV4cCI6MTU4ODUyNDgxOSwibGFzdG5hbWUiOiJHYXJjaWEiLCJsb2NhdGlvbiI6IiIsIm5hbWUiOiJDYXJsb3MifQ.nDRH76k2zChrd4V-FCQZzm5HocutXzdn2OWTjVnoxnY"

	expectedID := "5eada4b7ee3278555b9583c2"
	expectedEmail := "prueba@gmail.com"

	claims, _, userID, _ := utilities.ProcessToken(token)
	if expectedID != userID {
		t.Error("User ID not save properly")
		t.Fail()
	}
	if claims.Email != expectedEmail {
		t.Error("Email not save properly")
	}
}

func TestUserExist(t *testing.T) {
	expectedID := "5eada4b7ee3278555b9583c2"
	expectedExistence := true

	_, exist, userID := database.UserExist("prueba@gmail.com")

	if expectedID != userID {
		t.Error("ID's not equals")
	}
	if exist != expectedExistence {
		t.Error("Not found the user correctly")
	}
}
