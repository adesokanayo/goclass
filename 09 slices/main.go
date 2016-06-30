package main

import "fmt"

func main() {

	
	 var x []int

	for i:=0 ;i<10 ;i++{
        
        x=append(x, i+1)
    }
	
    
    for _,x:= range x{
        
        fmt.Println(x)
    }

fmt.Println("End of X\n","Start of Y")
y:= []int{11,12,13,14,15,}

 for _,y:= range y{
        
        fmt.Println(y)
    }


z:=y[0:2]


 for i,z:= range z{
        
       fmt.Printf("\nValue of slice:%d\n Position:%d",z, i )
    }

}