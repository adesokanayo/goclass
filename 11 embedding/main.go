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
func (s sport) match(searchTerm string) bool {
	return strings.Contains(s.team, searchTerm) || strings.Contains(s.city, searchTerm)
}

// Declare a struct type named hockey that represents specific
// hockey information. Have it embed the sport type first.

type hockey struct {
	sport
	grade string
}

// match checks the value for the specified term.
func (h hockey) match(searchTerm string) bool {

	return h.sport.match(searchTerm) || strings.Contains(h.grade, searchTerm)

	// Make sure you call into match method for the embedded sport type.

}

func main() {

	// Define the term to match.
	searchTerm := "bombers"

	data := []matcher{
		hockey{sport{"bombers", "lagos"}, "one"},
		hockey{sport{"Slayers", "Ibadan"}, "two"},
	}
	// Create a slice of matcher values and assign values
	// of the concrete hockey type.

	fmt.Println("looking for :", searchTerm)

	// Display what we are trying to match.

	for _, m := range data {
		if m.match(searchTerm) {
			fmt.Printf("%s is present", m)
		}
		//return err

	}

	// Range of each matcher value and check the term.
}
