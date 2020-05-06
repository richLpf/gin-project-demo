package common

import (
	"testing"
)

//TestArrayHasContainValue test
func TestArrayHasContainValue(t *testing.T) {
	a := 1
	b := []int{1, 2}
	result := ArrayHasContainValue(a, b)
	if result != true {
		t.Errorf("result shold be %t", result)
	}
}
