package configs

import (
	"testing"
)

func Test_DatabaseDB(t *testing.T) {
	DatabaseDB()
	if Database == nil {
		t.Error("NOT CONNECTED")
	}
}
