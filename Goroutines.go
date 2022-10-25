package main

import (
	"fmt"
	"time"
)

func out(from, to int) {
	for i := from; i < to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}

}
