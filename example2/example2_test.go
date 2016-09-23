package example2

import "testing"

func TestEmptySample(t *testing.T) {
	expected := 0
	input := []Tuple{}
	obtained := Sample(input)
	if obtained != expected {
		t.Fatalf("TestEmptySample failed, exp: %d, obt: %d\n", expected, obtained)
	}
}

func TestMultiSample(t *testing.T) {
	expected := 17
	input := []Tuple{{10, -5}, {10, 2}}
	obtained := Sample(input)
	if obtained != expected {
		t.Fatalf("TestMultiSample failed, exp: %d, obt: %d\n", expected, obtained)
	}
}
