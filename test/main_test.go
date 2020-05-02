package test

import (
	"github.com/solrac97gr/yendoapi/database"
	"testing"
)

func TestDbConnection(t *testing.T) {
	actual := true
	isConnected := database.ConnectionOK()
	if actual != isConnected {
		t.Fail()
	}
}
