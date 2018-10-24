package main

import "fmt"

func main() {
	// constructing an object as a struct
	// 1. declaration and initialization
	foo := myStruct{}
	foo.myField = "bar"

	// 2. Go's composite literal format (like a slice)
	foo1 := myStruct{"bar"}

	// 3. using the built-in new function
	// in this case foo is a simple memory address
	// go takes care of dereference the pointer when we use new function
	foo2 := new(myStruct)
	foo2.myField = "bar"

	foo3 := &myStruct{"bar"}

	fmt.Println(foo)
	fmt.Println(foo1)
	fmt.Println(foo2)
	fmt.Println(foo3)

	fmt.Println(foo.myField)
	fmt.Println(foo1.myField)
	fmt.Println(foo2.myField)
	fmt.Println(foo3.myField)

	// using constructors
	// foo4 := &myStruct2{}
	foo4 := myStruct2Constructor()
	foo4.myMap["bar"] = "baz"

	fmt.Println(foo4.myMap)

	// using methods
	mp := messagePrinter{"bar"}
	mp.printMessage()

	// using object composition
	// 1. one way of composition call
	emc := enhanceMessagePrinter{}
	emc.message = "bar from enhancement"
	emc.printMessage()

	// 2. another way of composition call
	emc1 := enhanceMessagePrinter{messagePrinter{"bar from enhancement"}}
	emc1.printMessage()

}

type myStruct struct {
	myField string
}

// constructor
type myStruct2 struct {
	myMap map[string]string
}

func myStruct2Constructor() *myStruct2 {
	result := myStruct2{}
	result.myMap = map[string]string{}
	return &result
}

// methods
type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	fmt.Println("from a method: " + mp.message)
}

// Object Composition (the inheritance way in Go)
type enhanceMessagePrinter struct {
	messagePrinter
}
