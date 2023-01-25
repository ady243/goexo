// // package main

// // import "fmt"

// // func main(){
// // 	var(
// // 		b bool
// // 		s string
// // 		i int
// // 		u uint
// // 		u8 uint8
// // 		i8 int8
// // 		i16 int16
// // 		f float32

// // )
// // b = true
// // s  = "Ady"
// // i = -15
// // u = 10 // ne prend que le positif
// // u8 = 255 // o - 255
// // i8 = -4 // -128 - 127
// // i16 = 32767 // limite 32764
// // f = 4  //ne prend pas des virgules

// // 	// y:=8
// // 	fmt.Println(b,s,i,u,u8,i8,i16,f)

// // }

// // package main

// // import "fmt"

// // func main(){
// // 	var(
// // x int
// // y int
// // 	)
// // 	x = 7
// // 	y = 3
// // 	fmt.Println(x + y)
// // 	fmt.Println(x % y)

// // // Checking if x is not equal to y.

// // 	fmt.Println(x != y)

// // // Comparing the values of x and y and returning a boolean value.
// // 	fmt.Println(x < y)

// // 	// Comparing the values of x and y and returning a boolean value.
// // 	fmt.Println(x <= y)

// // // Bitwise OR
// // 	fmt.Println(x | y)

// // 	// Comparing the values of x and y.
// // 	fmt.Println(x == y)

// // // Comparing the values of x and y and returning a boolean value.
// // 	fmt.Println(x >= y)

// // }
// package main

// import "fmt"

// func main(){
// 	var(
// x int
// y int
// x1 int
// y1 int
// 	)
// 	x = 7
// 	y = 7
// 	x1=3
// 	y1=8
// 	fmt.Println(x + y)
// 	fmt.Println(x % y)

// // Checking if x is not equal to y.

// 	fmt.Println(x != y)

// // Comparing the values of x and y and returning a boolean value.
// 	fmt.Println(x < y)

// 	// Comparing the values of x and y and returning a boolean value.
// 	fmt.Println(x <= y)

// // Bitwise OR
// 	fmt.Println(x | y)

// 	// Comparing the values of x and y.
// 	fmt.Println(x == y)

// // Comparing the values of x and y and returning a boolean value.
// 	fmt.Println(x >= y)

// // Checking if x is not equal to y and y1 is greater than or equal to x1.
// 	fmt.Println(x != y && y1 != x1)

// // // Checking if x is greater than or equal to y or y is greater than x.

// // 	fmt.Println(x >= y || y > x)

// }

// //make database repository 	with go

package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	
	rand.Seed(time.Now().UnixNano())
fmt.Println(rand.Int())
if x := rand.Int();
 x%2 == 0{
fmt.Println(" x est un nombre paire")
}else{
	fmt.Println("est un nombre impaire")
}

}