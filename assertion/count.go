package assertion

import (
	"errors"
	"reflect"
	"strconv"
	"github.com/kubernetes/kubernetes/pkg/util/integer"
)

type CountAssertion struct {
	Count	int
	Compare Comparer
	config	AssertionConfig
}

func (this *CountAssertion) Init(config AssertionConfig) (err error) {
	var intValue int
	switch config.Value.(type) {
	case float64:
		intValue = int(integer.RoundToInt32(config.Value.(float64)))
	case string:
		intValue, err = strconv.Atoi(config.Value.(string))
		return err
	}
	this.Count = intValue
	compare, err := NewComparer(config.Match)
	if err != nil {
		return err
	}
	this.Compare = compare
	this.config = config
	return nil
}

func (this *CountAssertion) Run(value interface{}) (bool, error) {
	// checking value is slice type to avoid panic
	// http://stackoverflow.com/questions/14025833/range-over-interface-which-stores-a-slice
	// could also accept maps - but in general we don't support assertions on complex types
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(value)

		success, err := this.Compare(this.Count, s.Len())
		if err != nil {
			return success, errors.New("Error occurred comparing length of value lists: "+err.Error())
		} else {
			return success, nil
		}
	default:
		return false, errors.New("Non-countable value found")
	}
}

func (this *CountAssertion) Info() AssertionConfig {
	return this.config
}