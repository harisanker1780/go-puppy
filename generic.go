package puppy

import "fmt"

type myNumbers interface {
	int | float64
}

func Generic() {
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.363, 2.4647))

	fmt.Println(addT1(1, 2))
	fmt.Println(addT1(1.363, 2.4647))
}

func addT[T int | float64](a, b T) T {
	return a + b
}

func addT1[T myNumbers](a, b T) T {
	return a + b
}
