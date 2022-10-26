package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func main() {
	
	 

      
	 



}
// example
type SmartList struct {
	words []string
}

func New(newWords []string) *SmartList {
	return &SmartList{newWords}
}

type Count int

// Sort words and printed it
func (sl *SmartList) GetAnswer() {
	count := []rune(sl.words[0])
	strconv.Atoi(string(count))
	fmt.Print(strconv.Atoi(string(count)))

}
