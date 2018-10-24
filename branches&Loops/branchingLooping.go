package main

func main() {
	// branches
	foo := 5

	if foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	switch {
	case foo == 1:
		println("one")
	case foo > 2:
		println("greater than 2")
	}

	// loops
	for i := 0; i < 5; i++ {
		println(i)
	}

	i := 0
	for {
		i++
		println(i)
		if i > 5 {
			break
		}
	}

	// collections
	s := []string{"foo", "bar", "buz"}

	for i, v := range s {
		println(i, v)
	}

	myMap := make(map[int]string, 0)
	myMap[1] = "foo"
	myMap[2] = "bar"
	myMap[3] = "buz"

	for k, v := range myMap {
		println(k, v)
	}

}
