package main

import (
	"fmt"
	"strings"
)

// matcher defines the behavior required for performing matching.
type matcher interface {
	match(searchTerm string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// match checks the value for the specified term.

//  hockey struct  that embed sport type

func (s sport) match(searchTerm string) bool {
	return strings.Contains(s.team, searchTerm) || strings.Contains(s.city, searchTerm)
}

type hockey struct {
	sport
	grade string
}

// match checks the value for the specified term.
func (h hockey) match(searchTerm string) bool {

	return h.sport.match(searchTerm) || strings.Contains(h.grade, searchTerm)

}

// Define the term to match.

//IDEA: converted the search item  to lowercase
// Create a slice of matcher values and display matching

func main() {

	search := "Bombers"
	searchTerm := strings.ToLower(search)

	data := []matcher{
		hockey{sport{"bombers", "lagos"}, "one"},
		hockey{sport{"Slayers", "Ibadan"}, "two"},
	}

	fmt.Println("looking for :", searchTerm)

	for _, m := range data {
		if m.match(searchTerm) {
			fmt.Printf("%s is present", m)
		}

	}

}
