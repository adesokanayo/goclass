package main

import "fmt"

func main() {
       
          pay:=10
    
    fmt.Printf("Values of int:%d ,Address of int:%d\n" ,pay, &pay)
    
    
   
    p:= &pay
    
     fmt.Println("Address  of pointer:",p)
    fmt.Println("Value  of pointer:", *p)
    
    
}