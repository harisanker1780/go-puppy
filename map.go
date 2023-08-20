package puppy

import "fmt"

func Map() {
	am := map[string]int{
		"Test1": 1,
		"Test2": 2,
	}

	fmt.Println("Test 1 is ", am["Test 1"])

	an := make(map[string]int)
	an["Test3"] = 3
	an["Test4"] = 4
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	for k, v := range an {
		fmt.Println(k, v)
	}

	delete(an, "Test3")
	for k, v := range an {
		fmt.Println(k, v)
	}
}

func MapDelete() {
	an := make(map[string]int)
	an["Test3"] = 3
	an["Test4"] = 4
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	for k, v := range an {
		fmt.Println(k, v)
	}

	delete(an, "Test3")

	fmt.Println("------------ After Delete -----------")
	for k, v := range an {
		fmt.Println(k, v)
	}
}

func MapIdiom() {
	an := make(map[string]int)
	an["Test3"] = 3
	an["Test4"] = 4
	if v, ok := an["Test5"]; ok {
		fmt.Println("Value is ", v)
	} else {
		fmt.Println("Value doesn't exist")
	}
}
