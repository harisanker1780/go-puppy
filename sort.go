package puppy

import (
	"fmt"
	"sort"
)

func Sort() {
	xi := []int{4, 3, 23, 8, 9}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("---------------------------")
	xs := []string{"XI", "FG", "PQ", "AV", "LN", "LA"}
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

func CustomSort() {
	people := []Person{
		{Firstname: "Test 23", Age: 23},
		{Firstname: "Test 22", Age: 22},
		{Firstname: "Test 32", Age: 32},
		{Firstname: "Test 43", Age: 43},
		{Firstname: "Test 18", Age: 18},
	}

	fmt.Println(people)
	fmt.Println("-------------------------")
	sort.Sort(ByAge(people))
	fmt.Println(people)

	fmt.Println("-------------------------")
	sort.Sort(ByName(people))
	fmt.Println(people)
}

type ByAge []Person

func (ba ByAge) Len() int {
	return len(ba)
}

func (ba ByAge) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

func (ba ByAge) Less(i, j int) bool {
	return ba[i].Age < ba[j].Age
}

type ByName []Person

func (bn ByName) Len() int {
	return len(bn)
}

func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

func (ba ByName) Less(i, j int) bool {
	return ba[i].Firstname < ba[j].Firstname
}
