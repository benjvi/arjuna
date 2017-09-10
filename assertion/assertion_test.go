package assertion

import (
	"testing"
	"encoding/json"
	"fmt"
)

func newAssertion(conf []byte) (Assertion,error) {
	fmt.Printf("loaded json: %s \n", string(conf))
	assertionConfig := AssertionConfig{}
	err := json.Unmarshal(conf, &assertionConfig)
	fmt.Printf("Unmarshalled assertion config : %+v \n", assertionConfig)
	if err != nil {
		return nil, err
	}
	assertion, _ := NewValueAssertion(assertionConfig)
	fmt.Printf("Constructed assertion: %+v \n", assertion)
	return assertion, nil
}

// value is required, but by default we assume equals operator
// matching value must pass assertion
func TestStringValueDefaultEqualSuccess(t *testing.T) {
	assertionJson := []byte(`{"value": "derek"}`)
	assertion, err := newAssertion(assertionJson)
	if err != nil {
		t.Fatalf("assertion unmarshalling failed with err: %+v \n", err)
	}
	matchingValue := "derek"
	_, err = assertion.Run(matchingValue)
	if err != nil {
		t.Fatalf("Equal value should have passed assertion, but got error: %+v \n", err)
	}
}

// value is required, but by default we assume equals operator
// matching value must pass assertion
func TestNumberValueDefaultEqualSuccess(t *testing.T) {
	assertionJson := []byte(`{"value": 3500}`)
	assertion, err := newAssertion(assertionJson)
	if err != nil {
		t.Fatalf("assertion unmarshalling failed with err: %+v \n", err)
	}
	matchingValue := float64(3500.00)
	_, err = assertion.Run(matchingValue)
	if err != nil {
		t.Fatalf("Equal value should have passed assertion, but got error: %+v \n", err)
	}

}

func TestValueDefaultEqualFailure(t *testing.T) {
	assertionJson := []byte(`{"value": "derek"}`)
	assertion, err := newAssertion(assertionJson)
	if err != nil {
		t.Fatalf("assertion unmarshalling failed with err: %+v \n" , err)
	}
	nonMatchingValue := "der"
	_, err = assertion.Run(nonMatchingValue)
	if err == nil {
		t.Fatalf("Expected assertion to fail with nonequal value but error was nil\n")
	}
}
