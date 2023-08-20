package puppy

import "fmt"

func Channels() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func ChannelsBuffer() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}

func DirectionalChannel() {
	c := make(chan int)

	//send
	go drc_foo(c)

	// receive
	drc_bar(c)

	fmt.Println("About to exit")
}

func drc_foo(c chan<- int) {
	c <- 42
}

func drc_bar(c <-chan int) {
	fmt.Println(<-c)
}

func ChannelRange() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

func ChannelSelect() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Send
	go send(even, odd, quit)

	// Receive
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func ChannelOkIdiom() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel", v)
		case v := <-o:
			fmt.Println("from the odd channel", v)
		case v := <-q:
			fmt.Println("from the quit chhanel", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
