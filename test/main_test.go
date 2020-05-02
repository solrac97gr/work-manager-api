package test

import (
	"testing"

	"github.com/solrac97gr/yendoapi/database"
)

func TestDbConnection(t *testing.T) {
	expected := true
	isConnected := database.ConnectionOK()
	if expected != isConnected {
		t.Fail()
	}
}
