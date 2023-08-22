package elements

import (
	"math"

	"github.com/ssd0247/mini-gonum/elements/utils"
)

type Vector struct {
	Values []float64
}

func (v *Vector) Bind(vfunc func(val interface{}, args ...interface{}) interface{}, args ...interface{}) {
	for i, val := range v.Values {
		v.Values[i] = vfunc(val, args...).(float64)
	}
}

func (v *Vector) Reduce(rfunc func(red, val interface{}, args ...interface{}) interface{}, args ...interface{}) float64 {
	reduced := args[0].(float64)
	for _, val := range v.Values {
		reduced = rfunc(reduced, val, args...).(float64)
	}
	return reduced
}

func (v *Vector) Clip(min, max float64) {
	v.Bind(utils.Clip, min, max)
}

func (v *Vector) Fill(value float64) {
	v.Bind(utils.Fill, value)
}

func (v *Vector) FillIndex(values []float64, index int) {
	vLen := len(v.Values)
	for i, val := range values {
		rIndex := index + i
		if rIndex < vLen {
			v.Values[index+i] = val
		}
	}
}

func (v *Vector) Zeros() {
	v.Bind(utils.Fill, 0.0)
}

func (v *Vector) Maximum() float64 {
	return v.Reduce(utils.Max, math.Inf(-1))
}

func (v *Vector) Minimum() float64 {
	return v.Reduce(utils.Min, math.Inf(1))
}

func (v *Vector) Exp() {
	v.Bind(utils.Exp)
}

func (v *Vector) Exp2() {
	v.Bind(utils.Exp2)
}

func (v *Vector) Log() {
	v.Bind(utils.Log)
}

func (v *Vector) Log2() {
	v.Bind(utils.Log2)
}

func (v *Vector) Pow(p float64) {
	v.Bind(utils.Pow, p)
}

func (v *Vector) Sum() float64 {
	return v.Reduce(utils.Sum, float64(0))
}

func (v *Vector) Add(a *Vector) float64 {
	if len(v.Values) != len(a.Values) {
		return math.Inf(1)
	}
	for i := range v.Values {
		v.Values[i] = v.Values[i] + a.Values[i]
	}
	return float64(0)
}

func (v *Vector) Max() (float64, int) {
	if len(v.Values) == 0 {
		return math.Inf(1), -1
	}
	max := v.Values[0]
	maxIndex := 0
	for idx, element := range v.Values {
		if element > max {
			max = element
			maxIndex = idx
		}
	}
	return max, maxIndex
}

func (v *Vector) Min() (float64, int) {
	if len(v.Values) == 0 {
		return math.Inf(-1), -1
	}
	min := v.Values[0]
	minIndex := 0
	for idx, element := range v.Values {
		if element < min {
			min = element
			minIndex = idx
		}
	}
	return min, minIndex
}

func (v *Vector) Unique() *Vector {
	var uniqueMap = make(map[float64]bool, len(v.Values)/3) // why divided by a constant '3' ??
	var keys []float64
	for _, e := range v.Values {
		uniqueMap[e] = true
	}
	for key := range uniqueMap {
		keys = append(keys, key)
	}
	return &Vector{Values: keys}
}

func (v *Vector) Mean() float64 {
	sum := float64(0)
	var (
		idx  int
		elem float64
	)

	for idx, elem = range v.Values {
		sum += elem
	}
	if sum == 0 {
		return float64(0)
	}
	return sum / float64(idx+1)
}
