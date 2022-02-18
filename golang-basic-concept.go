package study

import (
	"fmt"
	"sync"
	"time"
)

func study() {

	// POINTER
	b := 255
	// The & operator is used to get the address of a variable
	// *T is the type of the pointer variable which points to a value of type T.
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
}

// ----------------------------------------------------------------
// CONCURRENCY
// Dealing with a lot of things at once: eating and watching TV
// Breaking the code into smaller independent task that can run at the same time
func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	notRunAtTheSameTime := false
	if notRunAtTheSameTime == false {
		// it will run these at the same time if add "go" routine in it
		go count("sheep")
		count("fish")
	} else if notRunAtTheSameTime == true {
		count("sheep")
		count("fish")
	} else if 1 == 2 {
		go count("sheep")
		go count("fish")
		fmt.Scanln()
	} else if 1 == 3 {
		// in relatity this option is not useful because it requires money input
		go count("sheep")
		go count("fish")
		fmt.Scanln()
	} else if 1 == 4 {
		// create a go routine
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			count("sheep")
			wg.Done()
		}()

		wg.Wait()
	}
}

// channel type is string or anything even channel through channel
func countWithChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// send the value of thing over the channel
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	// finish sending we can close the channel
	close(c)
}

// call this in main

func main_concurrency() {
	// send and receive the message through the channel
	// use make function to create a new channel
	c := make(chan string)
	go countWithChannel("sheep", c)

	// this is to receive a msg through channel
	// for is infinate for loop
	for {
		msg, open := <-c

		// check if channel is still open so we can exit it
		if !open {
			break
		}

		fmt.Println(msg)
	}
}

func main_concurrency_2() {
	// make a channel of string, put the capacity of 2 will make it unblocks the channel
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}

func main_concurrency_3() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		// allow select statement to allow us to select a ready channel
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}
