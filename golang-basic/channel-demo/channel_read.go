package channeldemo

import (
	"fmt"
	"time"
)

func ReadChannelNotBlock(c chan int, timeout time.Duration) {
	select {
	case t := <-c:
		fmt.Println("received: ", t)
	case <-time.After(timeout):
		fmt.Println("receive timeout")
	}
}

func ReadChannelNotBlock2(c chan int) {
	select {
	case t := <-c:
		fmt.Println("received: ", t)
	default:
		fmt.Println("no data")
	}
}
