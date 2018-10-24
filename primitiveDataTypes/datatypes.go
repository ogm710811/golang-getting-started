package main

import (
	"fmt"
)

// constants declaration
const (
	first0  = "the first"
	second1 = "the second"
	third2  = "the third"

	// the iota value. how to use it?
	first3 = iota
	second4
	third5
)

const (
	first6 = iota
	second7
	third8
)

const (
	first9 = 1 << (10 * iota)
	second10
	third11
)

func main() {
	var myInt int
	myInt = 42

	println(myInt)

	var myFloat32 float32 = 42.
	println(myFloat32)

	myString := "Hello Go"
	println(myString)

	myComplex := complex(2, 3)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))

	// printing constants
	println(first0)
	println(second1)
	println(third2)

	println(first3)
	println(second4)
	println(third5)

	println(first6)
	println(second7)
	println(third8)

	println(first9)
	println(second10)
	println(third11)

	// collections
	// array declaration 1
	myArray := [3]int{}
	myArray[0] = 42
	myArray[1] = 27
	myArray[2] = 99
	// print element between brackets
	fmt.Println(myArray)

	// array declaration 2
	myArray2 := [...]int{42, 27, 99}
	fmt.Println(myArray2)
	// get length of the array
	fmt.Println(len(myArray2))

	// create slice of array
	mySlice0 := myArray[:]
	mySlice1 := myArray[1:]
	mySlice2 := myArray[:1]
	fmt.Println(mySlice0)
	fmt.Println(mySlice1)
	fmt.Println(mySlice2)

	// declaring slice: NOT need dots iniside brackets
	mySliceDec := []float32{14., 15., 19.}
	fmt.Println(mySliceDec)
	mySliceDec = append(mySliceDec, 100.)
	fmt.Println(mySliceDec)
	fmt.Println(len(mySliceDec))

	// declaring slice with huge length
	myHugeSlice := make([]float32, 100)
	myHugeSlice[0] = 12.
	myHugeSlice[5] = 15.
	myHugeSlice[10] = 18.
	fmt.Println(myHugeSlice)

	// creating Maps
	myMap := make(map[int]string, 0)
	fmt.Println(myMap)
	myMap[42] = "Foo"
	myMap[12] = "Bar"
	fmt.Println(myMap)
	fmt.Println(myMap[99])

}
