package puppy

import "fmt"

func Pointer() {
	a := 1
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)

	b := 1
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)
}

func addOne(v int) int {
	return v + 1
}

func addOneP(v *int) {
	*v += 1
}
