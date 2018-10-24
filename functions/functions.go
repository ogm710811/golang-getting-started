package main

func main() {
	message := "hello"
	sayHello(message)
	println(message)

	sayHelloRefer(&message)
	println(message)

	sayHelloMultiParams("hello", "from", "go")

	println(add(1, 3, 5, 9))
	lengthOfTerms, sum := addMultiples(1, 3, 5, 9)
	println("Numbers in Terms: ", lengthOfTerms, "total of terms:", sum)

	numTerms, sumTerms := addNamedParams(1, 3, 5, 9)
	println("Numbers in Terms: ", numTerms, "total of terms:", sumTerms)

	addAnonymouesFunc := func(terms ...int) (termsLength int, sum int) {
		for _, term := range terms {
			sum += term
		}

		termsLength = len(terms)

		return
	}

	numTermsAnonymous, sumTermsAnonymous := addAnonymouesFunc(1, 3, 5, 9)
	println("Numbers in Terms: ", numTermsAnonymous, "total of terms:", sumTermsAnonymous)

}

// passing parameters by value
func sayHello(message string) {
	println(message)

	// only create a copy of the variable
	message = "Hello Go"
}

// passing parameters by reference
func sayHelloRefer(message *string) {
	// this print out the memory address of the pointer
	println(message)

	// this print out the value of the pointer [this is call DEREFERENCE]
	println(*message)

	// only create a copy of the variable
	*message = "Hello Go"
}

// passing multiple parameters
func sayHelloMultiParams(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}

// returning values
// 1. single return value
func add(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}

	return result
}

// 2. multiple return values
func addMultiples(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}

	return len(terms), result
}

// 3. named return values
func addNamedParams(terms ...int) (termsLength int, sum int) {
	for _, term := range terms {
		sum += term
	}

	termsLength = len(terms)

	return
}

// anonymous functions
