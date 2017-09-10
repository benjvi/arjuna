package assertion

import (
	"errors"
	"reflect"
)

// Comparer returns an error only if type of expected value is not supported by the comparer implementation
// e.g. for equals operator, is not a simple json type or list thereof
// or for gt operator, is not an orderable type
// if actual value (from json doc being filtered) is unsupported/unexpected by comparer just return false
type Comparer func(expected, actual interface{})(bool, error)

func NewComparer(conf string) (Comparer, error) {
	switch conf {
	case "":
		return Equals, nil
	case "equals", "=", "==":
		return Equals, nil
	case "notEquals", "!=":
		return NotEquals, nil
	case "greaterThan", ">":
		return nil, errors.New("Not implemented")
	}

	return nil, errors.New("No matching operator found")
}

func Equals(expected, actual interface{}) (bool, error) {
	// should support simple json types
	// ie number, string, null, boolean and arrays of these types
	// note not supporting arrays, just slices
	// TODO preconditions on expected values to ensure no pointer values

	if expected != nil && actual != nil {
		switch reflect.TypeOf(actual).Kind() {
		// this is possibly the most complicated switch statement in the world
		// if we didn't need to compare pointer/non-pointer values this would be unnecessary
		case reflect.Ptr:
			println("Actual is pointer")
			value := reflect.Indirect(reflect.ValueOf(actual))
			return isPointedToValueEqual(value.Interface(), expected)
		case reflect.Bool, reflect.Float64, reflect.String, reflect.Int:
			println("Actual is regular value")
			return (actual == expected), nil
		case reflect.Slice:
			println("Actual is slice")
			if reflect.TypeOf(actual).Elem().Kind() == reflect.Ptr {
				println("Actual slice contains pointers")
				return isPointerSliceEqual(actual, expected)
			} else if reflect.TypeOf(actual).Elem().Kind() == reflect.Interface {
				println("Actual slice contains interfaces")
				// could have anything in here so go one by one comparing
				result, err := isInterfaceSliceEqual(actual, expected)
				println("final result")
				println(result)
				return result, err
			} else {
				println("Actual slice contains regular values")
				// list of regular value types
				return reflect.DeepEqual(actual, expected), nil
			}
		default:
			return false, errors.New("Unsupported value type in expected value: "+reflect.TypeOf(actual).Kind().String())
		}
	} else {
		return (actual == expected), nil
	}
}

func isInterfaceSliceEqual(iSlice interface{}, expected interface{}) (bool, error) {
	a := reflect.ValueOf(iSlice)
	e := reflect.ValueOf(expected)

	if reflect.TypeOf(expected).Kind() != reflect.Slice {
		println("expected is not a slice as expected")
		return false, nil
	}
	if a.Len() != e.Len() {
		println("actual slice length different than expected")
		return false, nil
	}

	println("cty736rcwycddc3r4c93c")
	for i := 0; i < a.Len(); i++ {
		aIndexVal := a.Index(i).Interface()
		eIndexVal := e.Index(i).Interface()
		println(reflect.TypeOf(aIndexVal).String())

		var isEqual bool
		var err error
		switch reflect.TypeOf(aIndexVal).Kind() {
		case reflect.Ptr:
			value := reflect.Indirect(reflect.ValueOf(aIndexVal))
			println(value.Interface().(string))
			println(eIndexVal.(string))
			isEqual, err = isPointedToValueEqual(value.Interface(), eIndexVal)
			println(isEqual)
		case reflect.Bool, reflect.Float64, reflect.String, reflect.Int:
			isEqual, err = (aIndexVal == eIndexVal), nil
		default:
			isEqual, err = false, errors.New("Unsupported value type in expected value: "+reflect.TypeOf(iSlice).Kind().String())
		}
		if !isEqual || err != nil {
			println("something went wrong")
			return isEqual, err
		}
	}
	println("interface slice fonund equal")
	return true, nil
}

func isPointedToValueEqual(value interface{}, expected interface{}) (bool, error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Bool, reflect.Float64, reflect.String, reflect.Int:
		println("reached the end")
		return (value == expected), nil
	case reflect.Slice:
		if reflect.TypeOf(value).Elem().Kind() == reflect.Ptr {
			// elements in list are pointers
			return isPointerSliceEqual(value, expected)
		} else if reflect.TypeOf(value).Elem().Kind() == reflect.Interface {
			// could have anything in here so go one by one comparing
			return isInterfaceSliceEqual(value, expected)
		} else {
			// list of regular value types
			return reflect.DeepEqual(value, expected), nil
		}
	default:
		return false, errors.New("Unsupported value type in actual value: "+reflect.TypeOf(value).String())
	}
}

func isPointerSliceEqual(pointerSlice interface{}, expected interface{}) (bool, error) {
	// elements in slice are pointers
	// TODO make sure expected is a slice!
	switch reflect.TypeOf(pointerSlice).Elem().Elem().Kind() {
	case reflect.Bool, reflect.Float64, reflect.String, reflect.Int:
		// types in slice need to be identical for comparison to work
		a := reflect.ValueOf(pointerSlice)
		e := reflect.ValueOf(expected)

		expValueList := make([]interface{}, a.Len())
		actualValueList := make([]interface{}, a.Len())
		for i := 0; i < a.Len(); i++ {
			actualValueList[i] = reflect.Indirect(a.Index(i)).Interface()
			expValueList[i] = e.Index(i).Interface()
		}
		return reflect.DeepEqual(actualValueList, expValueList), nil
	default:
		return false, errors.New("Unsupported value type for list elements in expected value: " + reflect.TypeOf(expected).Elem().Elem().Kind().String())
	}
}

func NotEquals(expected, actual interface{}) (bool, error) {
	result, err := Equals(expected, actual)
	return !result, err
}