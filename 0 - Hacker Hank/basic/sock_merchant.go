package basic

import "fmt"

func main() {
	var n int32 = 9
	ar := []int32{10, 20, 20, 10, 10, 10, 30, 50, 10, 20}
	fmt.Println(n, ar)
}

func sockMerchant(n int32, ar []int32) int32 {
	fmt.Println(n, ar)
	return n
}
