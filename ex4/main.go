package main 
import "fmt"


func main(){
 type name struct{
        
       fname string
       age int
         }
    
    myguy:=name{
        fname:"Nameless",
        age: 25,
    }


// value before 

fmt.Printf("FName Before :%s\n",myguy.fname)
//value after
fmt.Printf("FName After: %s\n",changename(myguy.fname))
    
}

// function to change the name by concatenating ab 
    
   func changename( fname string)string{
       
       return fname+"ab"
   }


