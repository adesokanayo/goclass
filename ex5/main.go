package main

import ("fmt"
"reflect") 

type counter int  

func main() {
    
 var d counter
 d=0
 
 fmt.Println(d)
 fmt.Println("D is a type of:",reflect.TypeOf(d))
    
    
  var e counter
  e=10
  fmt.Println(e)  
    
    var f int
 // at compile time , the below assignment was not allowed.    
    d=f
}