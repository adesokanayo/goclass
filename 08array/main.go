package main



import "fmt"



func main() {

var first[5]string

second:= [5]string{"one", "two", "three", "four", "five"}

    
    first=second
    fmt.Println(first)
     fmt.Println(&first)

}