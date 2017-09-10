package assertion

import (
	"testing"
)

func TestCountRunsFloatSlice(t *testing.T) {
	countAssert := CountAssertion{}
	countAssert.Init(AssertionConfig{
		Value: 2,
		Match: "equals",
	})
	_, err := countAssert.Run([]float64{1,2})
	if err !=nil {
		t.Fatal(err)
	}
	_, err = countAssert.Run([]float64{2})
	if err ==nil {
		t.Fatal(err)
	}
}

func TestCountRunsInterfaceSlice(t *testing.T) {
	countAssert := CountAssertion{}
	countAssert.Init(AssertionConfig{
		Value: 2,
		Match: "equals",
	})
	var s []interface{}
	_, err := countAssert.Run(s)
	if err ==nil {
		t.Fatal(err)
	}

	s = append(s, 1)
	s = append(s, 2)
	_, err = countAssert.Run(s)
	if err !=nil {
		t.Fatal(err)
	}
}

func TestCountFailsMap(t *testing.T) {
	countAssert := CountAssertion{}
	countAssert.Init(AssertionConfig{
		Value: 2,
		Match: "equals",
	})
	s := make(map[int]int)
	_, err := countAssert.Run(s)
	if err == nil {
		t.Fatal(err)
	}

	s[1]=1
	s[2]=2
	_, err = countAssert.Run(s)
	if err == nil {
		t.Fatal(err)
	}
}