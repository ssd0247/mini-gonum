package elements

import (
	"fmt"
	"strings"
)

// using generics

type Set[E comparable] map[E]struct{}

func NewSet[E comparable](vals ...E) Set[E] {
	s := Set[E]{}
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[E]) Add(vals ...E) {
	for _, v := range vals {
		s[v] = struct{}{}
	}
}

func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

func (s Set[E]) Members() []E {
	result := make([]E, 0, len(s))
	for v := range s {
		result = append(result, v)
	}
	return result
}

func (s Set[E]) String() string {
	str := "{"
	for _, val := range s.Members() {
		str = str + fmt.Sprintf("%v ", val)
	}
	str = strings.TrimRight(str, " ") + "}"
	return str
}

func (s Set[E]) Union(s2 Set[E]) Set[E] {
	uni := NewSet(s.Members()...)
	uni.Add(s2.Members()...)
	return uni
}

func (s Set[E]) Intersection(other Set[E]) Set[E] {
	intersect := NewSet[E]()
	for _, v := range s.Members() {
		if other.Contains(v) {
			intersect.Add(v)
		}
	}
	return intersect
}
