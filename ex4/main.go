package main 
import "fmt"


func main(){
 type name struct{
        
       name string
       age int
         }
    
    myguy:=name{
        name:"Nameless",
        age: 25,
    }


// value before 

fmt.Printf("Name Before :%s\n",myguy.name)
//value after
fmt.Printf("Name After: %s\n",changename(myguy.name))
    
}

// function to change the name by concatenating ab 
    
   func changename( name string)string{
       
       return name+"ab"
   }


