// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan int, 1)
// 	ch2 := make(chan int, 1)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch1 <- 1
// 	}()

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch2 <- 2
// 	}()

// 	select {
// 	case <-ch1:
// 		fmt.Println("Received from ch1")
// 	case <-ch2:
// 		fmt.Println("Received from ch2")
// 	case <-time.After(3 * time.Second):
// 		fmt.Println("Timeout")
// 	}
// }

// # 2. Non-blocking channel operations.

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	// Writing data to a channel when it is not full
	// will not block the operation.
	ch <- 1
	fmt.Println("Data written to channel")

	select {
	// Attempting to write data to a channel that is already full
	// will result in an immediate return, as the channel is unable to accommodate the new data.
	case ch <- 2:
		fmt.Println("Data written to channel")
	default:
		fmt.Println("Channel is full, unable to write data")
	}

	// Trying to read data from a channel when it already contains data
	// will not block the operation.
	data, ok := <-ch
	if ok {
		fmt.Println("Data read from channel:", data)
	}

	data, ok = <-ch
	if ok {
		fmt.Println("Data read from channel:", data)
	}

	select {
	// Attempting to read data from an empty channel will
	// result in an immediate return,
	// as the channel is empty and cannot fulfill the read operation.
	case data, ok := <-ch:
		if ok {
			fmt.Println("Data read from channel:", data)
		} else {
			fmt.Println("Channel is empty, unable to read data")
		}
	default:
		fmt.Println("Default : Channel is empty, unable to read data")
	}
}
