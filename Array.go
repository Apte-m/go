package main

import "fmt"

func arrayRange() {
	a := make([]int, 5)
	a[1] = 2
	a[2] = 7

	for i, v := range a {
		fmt.Println(i, v)

	}
}

func arrayString() {
	x := "hello"
	for _, v := range x {
		fmt.Printf("%v \n", v)

	}
}

func switchArray() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	var a int
	var b string
	fmt.Scanln(&b)
	//results = append (results, results[(len(results)-1)])
	results = append(results, b)
	for _, v := range results {
		switch {
		case v == "w":
			a += 3
		case v == "l":
			a += 0
		case v == "d":
			a += 1
		}
	}
	fmt.Println(a)
}
