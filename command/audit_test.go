package command

import (
	"testing"

	"os"
)


// Checking that audit runs without erroring
func TestAudit(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fail()
	}
	Audit(pwd)
	t.Fail()
}