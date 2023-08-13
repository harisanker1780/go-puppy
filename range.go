package puppy

import "fmt"

func Range() {

	xi := []int{1, 2, 3, 4, 5}
	for i, v := range xi {
		fmt.Println("Range over slice", i, v)
	}

	m := map[string]int{
		"Test":  1,
		"Test1": 2,
	}
	for i, v := range m {
		fmt.Println("Range over map", i, v)
	}

}
