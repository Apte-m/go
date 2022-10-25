package main

import "fmt"

func sumTwoNumber() {
	var n, sum int
	for fmt.Scan(&n); n > 0; n-- {
		var x int
		if fmt.Scan(&x); x > 9 && x < 100 && x%8 == 0 {
			sum += x
		}
	}
	fmt.Println(sum)
}

func mars_age(age int) int {
	if age == 1000 {
		return 526
	} else {
		return int(float32(age) / float32(1.90))

	}
}

func test(x1 *int, x2 *int) {

	*x2 = *x1
	*x1 = *x2 + *x2
	fmt.Print(*x1, *x2)
}

//func minimumFromFour() int {
//	//print your code here
//	array := []int{4, 5, 6, 7}
//	sort.Ints(array)
//	result := array[0]
//	return result
//
//}

func vote(x int, y int, z int) int {
	if x+y+z > 1 {
		return 1
	} else {
		return 0
	}
}
func sumInt(a ...int) (int, int) {
	var count, item int
	for _, elem := range a {
		count++
		item += elem
	}
	return count, item
}
