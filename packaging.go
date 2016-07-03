package main

import ("github.com/adesokanayo/goclass/toy"
 "fmt"
)


int main(){

p:= toy.New("Ayo",20)
}

p.UpdateOnHand(343)
p.UpdateSold(49)


fmt.Println("Name:",p.Name)
fmt.Println("Weight:",p.Weight)
fmt.Println("OnHand:",p.OnHand)
fmt.Println("Sold:",p.Onsold)
}