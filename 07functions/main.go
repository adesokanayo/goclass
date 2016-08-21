//functions

package main

import "fmt"

type user struct {
	username string
	serial   int
}

func createuser() (*user, error) {

	return &user{"truethinker", 123}, nil

}

func main() {

	newsignup, err := createuser()

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(newsignup)

	//to test for the errors, I dont understand why he used equals = and not :=

	_, err = createuser()
	if err != nil {
		fmt.Println(err)
	}

}
