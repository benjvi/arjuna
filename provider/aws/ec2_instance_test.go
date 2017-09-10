package aws

import (
	"testing"
	"fmt"
)

func TestDescribeInstances(t *testing.T) {
	collection := EC2InstanceCollection{}
	resources := collection.List()
	fmt.Printf("EC2 resources: %+v", resources)
	//TODO: mock out aws lib
	if len(resources) == 0 {
		t.Error("Describe instances didn't work - no instances!")
	}
}