package policy

import (
	"testing"

	"os"
	"fmt"
)

func TestLoadPolicies(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fail()
	}
	fmt.Printf("Test pwd: %s\n", pwd)
	policies := FromDir(pwd)

	for _,p := range policies {
		fmt.Printf("policy: %+v\n", p)
	}
	if len(policies) == 0 {
		t.Errorf("No policy files!")
	}
	if len(policies) > 1 {
		t.Errorf("Incorrect number of policy files: %v\n", policies)
	}
	for _,p := range policies {
		if p.Provider != "aws" {
			t.Errorf("unxpected provider value: %s, wanted: aws\n", p.Provider)
		}
		if p.ResourceType != "ec2" {
			t.Errorf("unxpected resource value: %s, wanted: ec2\n", p.ResourceType)
		}
	}
}