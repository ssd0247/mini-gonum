package main

import (
	"fmt"

	e "github.com/ssd0247/mini-gonum/elements"
)

func main() {
	s := e.NewSet(1, 2, 3, 4, 4, 4, 2, 4, 5, 6)

	s.Add(8, 2, 4, 6, 1)

	fmt.Println(s.Contains(100))
	fmt.Println()
	fmt.Println(s.Members())
	fmt.Println()
	fmt.Printf("SET-I : %s\n", s.String())

	s1 := e.NewSet(11, 12, 1, 10, 4, 2)
	fmt.Printf("SET-II : %s\n", s1.String())

	// Union of the two sets
	fmt.Printf("Union(SET-I, SET-II) : %s\n", s.Union(s1))
	// Intersection of the two sets
	fmt.Printf("Intersection(SET-I, SET-II) : %s\n", s.Intersection(s1))
}
