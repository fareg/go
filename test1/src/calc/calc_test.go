package calc

import "testing"

func TestAdd(t *testing.T) {
	var result int
	result = Add(15, 10)
	if (result != 26) {
		t.Error("Expected 25, got ", result)
	}
}

func TestSubstract(t *testing.T) {
	var result int
	result = Substract(15, 10)
	if (result != 6) {
		t.Error("Expected 5, got ", result)
	}
}
