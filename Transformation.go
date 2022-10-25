package main

import (
	"fmt"
	"math"
)

func transform() {
	fmt.Println(math.MaxInt8)   // 127
	fmt.Println(math.MaxUint8)  // 255
	fmt.Println(math.MaxInt16)  // 32767
	fmt.Println(math.MaxUint16) // 65535
}
