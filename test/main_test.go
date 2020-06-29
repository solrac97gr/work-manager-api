package test

import (
	"github.com/solrac97gr/yendoapi/database/user"
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
	token := "BearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI1ZWFkYTRiN2VlMzI3ODU1NWI5NTgzYzIiLCJhdmF0YXIiOiJodHRwczovL3Njb250ZW50LmZsaW0xOC0zLmZuYS5mYmNkbi5uZXQvdi90MS4wLTkvODkwMzc2NzNfMTEwODU2OTI3NjE2Mjc2Nl84ODk0NzY3NjUzODYwNjcxNDg4X24uanBnP19uY19jYXQ9MTA4XHUwMDI2X25jX3NpZD04NWE1NzdcdTAwMjZfbmNfZXVpMj1BZUd5NFNCSi1wMnRFWE8xT25nUVEwY3FZY29QaFFLVEVSUmh5Zy1GQXBNUkZBZ2dLMUd1VHlJMGpmT2wtX2xBaUcyd0lyMmp4TDFzakhXdmlnMjdOeUYyXHUwMDI2X25jX29oYz1SaWRwUnI5RFIwNEFYOTNyY0lEXHUwMDI2X25jX2h0PXNjb250ZW50LmZsaW0xOC0zLmZuYVx1MDAyNm9oPTYwMTA3NGQzZjQyMTBjOGY3YmMxZWIzZDdlYzhmY2FlXHUwMDI2b2U9NUVEM0FBRkMiLCJiaXJ0aGRhdGUiOiIxOTk3LTA0LTEzVDAwOjAwOjAwWiIsImVtYWlsIjoicHJ1ZWJhQGdtYWlsLmNvbSIsImV4cCI6MTU5MDU1MjQ0MCwibGFzdG5hbWUiOiJHYXJjaWEgUm9zYWxlcyIsImxvY2F0aW9uIjoiTGltYSwgUGVyw7oiLCJuYW1lIjoiQ2FybG9zIEFsYmVydG8ifQ.5vP0JrcTsTZTqZKkSOuFH2JIpfBGhkjartHwWuGWbqw"

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

	_, exist, userID := user.UserExist("prueba@gmail.com")

	if expectedID != userID {
		t.Error("ID's not equals")
	}
	if exist != expectedExistence {
		t.Error("Not found the user correctly")
	}
}
