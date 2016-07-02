package main

import "fmt"

func main() {

	type student struct {
		name  string
		email string
		age   int
	}

	post := student{
		name:  " odunayo",
		email: "adesokanayo@gmail.com",
		age:   29,
	}
    
    

	fmt.Println("Name: ", post.name)
	fmt.Println("Email: ", post.email)
	fmt.Println("Age: ", post.age)

	student2 := struct {
		name  string
		email string
		age   int
	}{
        // i struggled with a this part seriously due to format problem.
		name:  "Second Name",
		email: "secondemailgmail.com",
		age:   20,
	}

	fmt.Println("Name: ", student2.name)
	fmt.Println("Email: ", student2.email)
	fmt.Println("Age: ", student2.age)

}
