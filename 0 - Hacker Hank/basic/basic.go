package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var oddChan []string
	var evenChan []string
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr := []int{238, 447, 122, 200, -543, 0}
	for _, v := range arr {
		if v%2 == 0 {
			text := strconv.Itoa(v)
			oddChan = append(oddChan, text)
			continue
		}
		text := strconv.Itoa(v)
		evenChan = append(evenChan, text)
	}
	evenChan = append(evenChan, oddChan...)
	result := strings.Join(evenChan, "\n")
	fmt.Println(result)
}
