package puppy

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type person struct {
	firstname string
	lastname  string
	age       int
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func Struct() {
	sa1 := secretAgent{
		person: person{
			firstname: "James",
			lastname:  "Bond",
			age:       32,
		},
		ltk: false,
	}

	fmt.Println(sa1)
	fmt.Printf("%T\n", sa1)
	fmt.Printf("%#v\n", sa1)

	fmt.Println(sa1.firstname, sa1.firstname, sa1.age, sa1.person)
	saySomething(sa1)

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	var b bytes.Buffer

	sa1.person.writeOut(f)
	sa1.person.writeOut(&b)

	fmt.Println(b.String())
}

func AnonymousStruct() {
	p2 := struct {
		firstname string
		lastname  string
		age       int
	}{
		firstname: "James",
		lastname:  "Bond",
		age:       32,
	}

	fmt.Printf("%T\n", p2)
}
