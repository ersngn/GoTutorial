package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := 123
	b := "golang"
	c := 3.1415
	d := true
	e := "Hi"
	f := "Hey There"
	g := 'A'
	h := "A"
	var i byte = 255
	var j rune = 2147483647

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)
	fmt.Printf("%v \n", i)
	fmt.Printf("%v \n", j)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
	fmt.Printf("%T \n", h)
	fmt.Printf("%T \n", i)
	fmt.Printf("%T \n", j)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.ReadTrace())

}
