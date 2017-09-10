package assertion

import "testing"

func TestEqualsMatcherSimpleJSONValueTypes(t *testing.T) {
	op, err := NewComparer("equals")
	if err != nil {
		t.Fatalf("Didnt construct operator, got error: %+v", err)
	}
	var exp, actual interface{}
	exp = "test"
	actual = "test"
	result, err := op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b match, but returned false")
	}
	exp = "test"
	actual = "diff"
	result, err = op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if result {
		t.Fatal("a and b match, but returned false")
	}

	exp = 5.0
	actual = 5.0
	result, err = op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b match, but returned false")
	}

	exp = false
	actual = false
	result, err = op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b match, but returned false")
	}


}

func TestEqualsMatcherHandlesInt(t *testing.T) {
	// e.g. equals matcher may be called in count operator
	op, err := NewComparer("equals")
	if err != nil {
		t.Fatalf("Didnt construct operator, got error: %+v", err)
	}
	var exp, actual interface{}
	exp = 1
	actual = 1
	result, err := op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b match, but returned false")
	}

	exp = 1
	actual = 2
	result, err = op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if result {
		t.Fatal("a and b don't match, but returned false")
	}
}

func TestEqualsMatcherHandlesListTypes(t *testing.T) {
	op, err := NewComparer("equals")
	if err != nil {
		t.Fatalf("Didnt construct operator, got error: %+v", err)
	}
	var exp, actual interface{}
	exp = []float64{1.0, 2.0}
	actual = []float64{1.0, 2.0}
	result, err := op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b match, but returned false")
	}

	exp = []string{"a", "b"}
	actual = []string{"a", "b"}
	result, err = op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b match, but returned false")
	}

	exp = []string{"a", "c"}
	actual = []string{"a", "b"}
	result, err = op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if result {
		t.Fatal("a and b match, but returned false")
	}

	exp = []float64{1.0, 2.0}
	actual = []string{"a", "b"}
	result, err = op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if result {
		t.Fatal("a and b match, but returned false")
	}

	exp = []float64{1.0, 2.0}
	actual = []string{"a", "b"}
	result, err = op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if result {
		t.Fatal("a and b match, but returned false")
	}
}

func TestEqualsMatcherHandlesNilValues(t *testing.T) {
	op, err := NewComparer("equals")
	if err != nil {
		t.Fatalf("Didnt construct operator, got error: %+v", err)
	}
	var exp, actual interface{}
	exp = nil
	actual = nil
	result, err := op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b match, but returned false")
	}
}

func TestEqualsMatcherHandlesSimplePointers(t *testing.T) {
	// e.g. 'actual' value may be passed as a pointer
	op, err := NewComparer("equals")
	if err != nil {
		t.Fatalf("Didnt construct operator, got error: %+v", err)
	}
	var exp interface{}
	var actual int
	exp = 1
	actual = 1
	result, err := op(exp, &actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b match, but returned false")
	}

	exp = 1
	actual = 2
	result, err = op(exp, &actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if result {
		t.Fatal("a and b don't match, but returned false")
	}
}

func TestEqualsMatcherHandlesPointersOfSlices(t *testing.T) {
	// i.e. 'actual' value may be passed as a pointer
	op, err := NewComparer("equals")
	if err != nil {
		t.Fatalf("Didnt construct operator, got error: %+v", err)
	}

	actual1 := 1
	exp := []int{1}
	actual := []*int{&actual1}
	result, err := op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b match, but returned false")
	}
}

func TestEqualsMatcherHandlesPointersOfSlicesOfPointers(t *testing.T) {
	// e.g. 'actual' value may be passed as a pointer to slice containing pointer values!
	op, err := NewComparer("equals")
	if err != nil {
		t.Fatalf("Didnt construct operator, got error: %+v", err)
	}

	actual1 := 1
	exp := []int{1}
	actual := []*int{&actual1}
	result, err := op(exp, &actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b match, but returned false")
	}
}

func TestNotEqualsMatcher(t *testing.T) {
	op, err := NewComparer("notEquals")
	if err != nil {
		t.Fatalf("Didnt construct operator, got error: %+v", err)
	}
	var exp, actual interface{}
	exp = "test"
	actual = "test"
	result, err := op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if result {
		t.Fatal("a and b are equal (NOT notEqual) but returned true")
	}

	exp = "test"
	actual = "diff"
	result, err = op(exp, actual)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !result {
		t.Fatal("a and b are notEqual but returned false")
	}
}

func TestListNotEqualToValue(t *testing.T) {
	t.Skip("not implemented")
}
