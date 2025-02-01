package main

import "fmt"

func main() {
	fmt.Println(isArraySpecial([]int{1, 2, 3, 4, 5}))
	fmt.Println(isArraySpecial([]int{2, 1, 4}))
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}))
}
