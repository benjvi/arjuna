package assertion

import (
	"errors"
)

type ValueAssertion struct {
	Expected 	interface{}
	Compare 	func(expected, actual interface{})(bool, error)
}

func (this *ValueAssertion) Init(config AssertionConfig) error {
	this.Expected = config.Value
	if config.Match == "" {
		config.Match = "equals"
	}
	op, err := NewComparer(config.Match)
	if err != nil {
		return err
	}
	this.Compare = op
	return nil
}

func (this *ValueAssertion) Run(value interface{}) (bool, error) {
	success, err := this.Compare(this.Expected, value)
	if err != nil {
		return false, errors.New("Error occurred comparing values: "+err.Error())
	} else {
		return success, nil
	}
}