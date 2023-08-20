package puppy

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func JsonMarshal() {
	p1 := Person{
		Firstname: "Test",
		Lastname:  "L1",
		Age:       22,
	}

	p2 := Person{
		Firstname: "Test",
		Lastname:  "L2",
		Age:       32,
	}

	persons := []Person{p1, p2}
	bs, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

func JsonUnMarshal() {
	s := `[{"Firstname":"Test","Lastname":"L1","Age":22},{"Firstname":"Test","Lastname":"L2","Age":32}]`
	bs := []byte(s)
	var people []Person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range people {
		fmt.Printf("%d) %s %s, %d \n", i+1, v.Firstname, v.Lastname, v.Age)
	}
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Firstname, p.Age)
}
