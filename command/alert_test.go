package command

import (
	"testing"

	"os"
)

func TestAlert(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fail()
	}
	Alert(pwd)
}