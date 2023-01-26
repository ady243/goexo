package main

import (
	"fmt"
)
// The function main() is the entrnumber point of the program. It is the first function to be called when
// the program is run
var x int

// func mise(number int){
// 	number = 0
// fmt.Println(number)
// }
func miseAjour(pointerNumber*int){
*pointerNumber = 80

}
func main(){
	// x = 5
	// fmt.Println(x)
	// f()
	// shownumber()
	// List := [...]int{10,40,50,60,70,100,400}
	// for pos,value := range List{
	// 	fmt.Println(value,"   est à l'ordre ",pos)
	// }
// Declaring a map

// danumbersInnumberear := 0

// super := map[string]int{
// 	"janvier":31,
// 	"Fevr":31,
// 	"Mars":31,
// 	"Avril":31,
// 	"Mai":31,
// 	"juin":31,
// 	"juillet":31,
// 	"Aout":31,
// 	"Septembre":31,
// 	"Octobre":31,
// 	"Novembre":31,
// 	"Decembre":31,
	
// }
// Looping through the map and adding the values to the variable danumbersInnumberear.
// for _,value := range super{
// danumbersInnumberear = danumbersInnumberear + value
// }
// fmt.Printf("nombre de jours : %d jours .\n",danumbersInnumberear)

// for kenumber,value := range super{
// 	fmt.Printf("%v comprend  %d jours.\n", kenumber,value)

// }
// w:=func(){
// fmt.Println("saluuuuut")
// }
// w()

// a:=func()string{
// return "adnumber"
// }()

// fmt.Println(a)
// a,b :=func ()(int, int){
// 	// fmt.Println("hola")
// 	return 4,7
// }()
// fmt.Println(a)
// fmt.Println(b)



number := 10

fmt.Println("ok",number)
fmt.Println("yes",&number)

mypoint := &number
fmt.Printf("le pointerNumber %d est  %d.\n",mypoint,*mypoint)
// Updating the value of the variable `number` to 80.
miseAjour(mypoint)
fmt.Println(number)


}

// The variable x is declared and initialized to 10

// func f(){
// 	x:= 10 //local
// 	fmt.Println(x)
// }

