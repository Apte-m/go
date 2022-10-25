package main

import "fmt"

type Timer struct {
	name  string
	count int
}

func (t *Timer) tick() {
	for i := 0; i < t.count; i++ {
		fmt.Println(i + 1)
	}
}
