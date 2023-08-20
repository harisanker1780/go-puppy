package puppy

import "fmt"

func Slice() {

	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Println(xs)
	fmt.Printf("%T\n", xs)

	for i, v := range xs {
		fmt.Printf("%v - %v", i, v)
	}
}

func SliceAppend() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)

	xi = append(xi, 45, 46, 47)
	fmt.Println(xi)
}

func SliceDelete() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//Delete 4
	xi = append(xi[:4], xi[5:]...)
	fmt.Println(xi)
}

func SliceMake() {
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println("----------------------------")
	xi = append(xi, 10, 11, 12, 13, 14)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
}

func SliceCopy() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 6)

	copy(b, a)

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("-----------------")

	a[0] = 7

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("-----------------")
}
