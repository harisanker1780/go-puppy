package puppy

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Sum(ii ...int) int {
	r := 0
	for _, v := range ii {
		r += v
	}
	return r
}

func WriteToFile() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	s := []byte("Test")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

func AnonymousFunction() {
	x := func(s string) {
		fmt.Println("Anonymous function", s)
	}

	x("Test 1")
}

func ReturnFucntion() {
	x := bar()
	fmt.Println(x())
}

func CallbackFunction() {
	x := doMath(20, 10, add)
	fmt.Println(x)

	y := doMath(20, 10, subtract)
	fmt.Println(y)
}

func ClosureFunction() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func bar() func() int {
	return func() int {
		return 42
	}
}

func (p person) speak() {
	fmt.Println("I'am ", p.firstname)
}

func (s secretAgent) speak() {
	fmt.Println("I'am secret agent", s.firstname)
}

func saySomething(h human) {
	h.speak()
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.firstname))
	return err
}
