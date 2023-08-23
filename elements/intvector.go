package elements

import "fmt"

// optArg type is used to describe optional arguments
type optArg interface{}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// IntVector is the container for vector containing
// INTEGER values.
type IntVector struct {
	Values []int // use architecture-independent 'int' datatype
}

// CopyFrom (*IntVector method) copies the elements from one slice into the vector.
//
// If the slice has less number of elements than the vector, the slice containing
// the vector elements (= v.Values) is reallocated to a shorter underlying array
// that contains the elements of slice (param c).
func (v *IntVector) CopyFrom(c []int) bool {
	if len(c) == 0 {
		return false
	}
	if v.Values == nil {
		v.Values = make([]int, len(c))
	} else if len(c) > len(v.Values) {
		return false
	} else if len(c) < len(v.Values) {
		v.Values = make([]int, len(c))
		copy(v.Values, c)
	}
	for i := range v.Values {
		v.Values[i] = c[i]
	}
	return true
}

func (v *IntVector) Zeros() {
	for i := range v.Values {
		v.Values[i] = 0
	}
}

// Padding (*IntVector method) pads the vector either before its
// ``start`` or after its ``end``, depending upon the <param optArg>
// value.
func (v *IntVector) Padding(val, len int, prefix optArg) error {
	if len <= 0 || prefix == nil {
		return nil
	}
	d := make([]int, len)
	for i := range d {
		d[i] = val
	}
	if val, ok := prefix.(bool); ok {
		if val {
			v.Values = append(d, v.Values...)
		} else {
			v.Values = append(v.Values, d...)
		}
		return nil
	} else {
		return fmt.Errorf("[PADDING-ERROR]: provide boolean argument for optional parameter 'prefix'. Provided %T", prefix)
	}
}

// Remove (*IntVector method) removes element at index idx and returns a new vector.
func (v *IntVector) Remove(idx int) *IntVector {
	if idx < 0 || idx >= len(v.Values) {
		return nil
	}

	values := v.Values
	newValues := make([]int, len(values)-1)
	newValues = append(newValues, values[:idx]...)
	newValues = append(newValues, values[idx+1:]...)

	return &IntVector{Values: newValues}
}

// Add (*IntVector method) adds the corresponding elements of the vector
// supplied in argument with the ``v.Values`` elements.
func (v *IntVector) Add(a *IntVector) {
	// (a) v.Values=[0 0 0 0] is differet from (b) v.Values=[]
	// (a) returns a.Values; (b) returns nil
	if len(v.Values) == 0 || len(a.Values) == 0 || len(v.Values) != len(a.Values) {
		return
	}
	for i := range v.Values {
		v.Values[i] = v.Values[i] + a.Values[i]
	}
}

// multHelper uses generics to generate switch-case output
// for similar Number datatypes like int | int8... | float32 | float64
func multHelper[T Number](v *IntVector, a T) {
	for i := range v.Values {
		v.Values[i] = int(a) * v.Values[i]
	}
}

// Mult (*IntVector method) returns as per the following cases :
//
// - scalar quantity is supplied -> []int (scalar multiplication)
// - vector of type (*IntVector) is supplied -> int (dot product)
//
// ... rest all cases result in errors
func (v *IntVector) Mult(a interface{}) (interface{}, error) {
	switch x := a.(type) {
	case nil:
		return nil, fmt.Errorf("[MULT-ERROR]: no argument provided")
	case int:
		multHelper(v, x)
		return v.Values, nil
	case float32:
		multHelper(v, x)
		return v.Values, nil
	case float64:
		multHelper(v, x)
		return v.Values, nil
	case *IntVector:
		if len(x.Values) != len(v.Values) {
			return nil, fmt.Errorf("[MULT-ERROR]: vectors are of different lengths")
		}
		multValue := 1
		for i := range v.Values {
			multValue = v.Values[i] * x.Values[i]
		}
		return multValue, nil
	default:
		return nil, fmt.Errorf("[MULT-ERROR]: type implements interface, does not support multiplication")
	}
}
