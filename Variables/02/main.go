package main

import "fmt"

func main() {
	var (
		name     string  = "Ersin"
		age      int8    = 25
		isMariad bool    = false
		weight   float64 = 70.4
		height   uint8   = 174
	)

	name2, age2, ismaried2, weight2, height2 := "Ali", 50, true, 80.56, 175

	var name3 string = "Veli"
	var age3 int8 = 52
	var ismaried3 bool = false
	var weight3 float32 = 65.76
	var height3 uint8 = 180

	fmt.Println()
	fmt.Println("Person 1")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMariad)
	fmt.Println(weight)
	fmt.Println(height)

	fmt.Println()
	fmt.Println("Person 2")
	fmt.Println(name2)
	fmt.Println(age2)
	fmt.Println(ismaried2)
	fmt.Println(weight2)
	fmt.Println(height2)

	fmt.Println()
	fmt.Println("Person 3")
	fmt.Println(name3)
	fmt.Println(age3)
	fmt.Println(ismaried3)
	fmt.Println(weight3)
	fmt.Println(height3)
}
