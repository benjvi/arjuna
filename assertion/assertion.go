package assertion

import (
	"encoding/json"
)

type AssertionConfig struct {
	Property string
	Match    string
	Value    interface{}
}

type Assertion interface {
	Init(config AssertionConfig)	error
	Run(resources interface{})	(bool, error)
}

func NewSetAssertion(conf AssertionConfig) (Assertion, error) {
	var assertion Assertion
	switch conf.Property {
	case "count":
		assertion = &CountAssertion{}
	case "value":
		assertion = &ValueAssertion{}
	default:
		assertion = &CountAssertion{}
		conf.Property = "count"
		countVal, _ := json.Marshal(0)
		conf.Value = countVal
	}
	err := assertion.Init(conf)
	if err != nil {
		return nil, err
	}
	return assertion, nil
}

func NewValueAssertion(conf AssertionConfig) (Assertion, error) {
	var assertion Assertion
	switch conf.Property {
	case "count":
		assertion = &CountAssertion{}
	case "value":
		assertion = &ValueAssertion{}
	default:
		assertion = &ValueAssertion{}
	}
	err := assertion.Init(conf)
	if err != nil {
		return nil, err
	}
	return assertion, nil
}