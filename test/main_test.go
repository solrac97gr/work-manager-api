package test

import (
	"testing"

	"github.com/solrac97gr/yendo/database"
)

func testDbconection(t *testing.T) {
	actual := true
	isConected := database.ConnectionOK()
	if actual != isConected {
		t.Logf("The database conection fail")
		t.Fail()
	}
}
