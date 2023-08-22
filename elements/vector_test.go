package elements

import (
	"math"
	"testing"
)

func TestVector_MeanWithEmpty(t *testing.T) {
	vec := &Vector{}
	actual := vec.Mean()
	expected := 0.0
	if expected != actual {
		t.Errorf("Expected %f do not match actual %f", expected, actual)
	}
}

func TestVector_Mean(t *testing.T) {
	vec := &Vector{[]float64{10, 20, 30}}
	expected := 20.0
	actual := vec.Mean()
	if expected != actual {
		t.Errorf("Expected %f do not match actual %f", expected, actual)
	}
}

func TestVector_Exp(t *testing.T) {
	vec := &Vector{[]float64{10, 10, 10}}
	vec.Exp()
	if vec.Values[0] != math.Exp(10) {
		t.Errorf("Expected %f do not match actual %f", math.Exp(10), vec.Values[0])
	}
}
