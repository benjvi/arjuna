package filter

import (
	"testing"
	"github.com/benjvi/arjuna/assertion"
	"github.com/benjvi/arjuna/provider"
	"fmt"
	"encoding/json"
)

type testState struct {
	Key 	*string
}

type testResource struct {
	state 	*testState
}

func (this testResource) State() interface{} {
	return this.state
}

func (this testResource) Summary() string {
	return fmt.Sprintf("%v", this.state)
}

func TestFilter(t *testing.T) {
	filterConf := FilterConfig{
		Type: 	"JMESPath",
		Assert: assertion.AssertionConfig{
			Value: "abc",
		},
		Expression: "Key",
	}
	f, err := New(filterConf)
	if err != nil {
		t.Fatal(err)
	}


	resources := getTestResources()
	result, err := f.Run(resources)
	if err != nil {
		t.Fatal(err)
	}
	for _, res := range result {
		fmt.Printf("Resource matched filter: %v\n", res)
	}
	if len(result) != 1 {
		t.Fatal(fmt.Sprintf("Filtered result set was wrong size: %d, expected: 1", len(result)))
	}

}

func getTestResources() []provider.Resource {
	jsonStr1 := "{\"key\":\"abc\"}"
	state1 := testState{}
	json.Unmarshal([]byte(jsonStr1), &state1)
	res1 := testResource{
		state:	&state1,
	}
	jsonStr2 := "{\"key\":\"def\"}"
	state2 := testState{}
	json.Unmarshal([]byte(jsonStr2), &state2)
	res2 := testResource{
		state:	&state2,
	}
	resources := []provider.Resource{res1, res2}
	return resources
}
