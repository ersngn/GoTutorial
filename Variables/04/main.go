//Conversion, Not Casting
//Conversion means we take the value of one type and convert it to another type.
//thanks for Todd McLood

package main

import (
	"fmt"
)

var a int

type lahmacun int

var b lahmacun

func main() {
	a = 42
	b = lahmacun(a)

	fmt.Println(a)
	fmt.Printf("%T \n", a)
	fmt.Println(b)
	fmt.Printf("%T \n", b)
}
