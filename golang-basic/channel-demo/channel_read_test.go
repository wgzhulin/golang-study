package channeldemo

import (
	"testing"
	"time"
)

func TestReadChannelNotBlock(t *testing.T) {
	ch := make(chan int)

	waitTime := time.Second * 2
	go func() {
		time.Sleep(time.Second * 1)
		ch <- 1
	}()

	ReadChannelNotBlock(ch, waitTime)
}

func TestReadChannelNotBlock_Timeout(t *testing.T) {
	ch := make(chan int)

	waitTime := time.Second * 1
	go func() {
		time.Sleep(waitTime + time.Second)
		ch <- 1
	}()

	ReadChannelNotBlock(ch, waitTime)
}

func TestReadChannelNotBlock2_Timeout(t *testing.T) {
	ch := make(chan int)

	waitTime := time.Second * 1
	go func() {
		time.Sleep(waitTime + time.Second)
		ch <- 1
	}()

	ReadChannelNotBlock2(ch)
}
