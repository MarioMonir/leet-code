package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}

	merge(nums1, 3, []int{2, 5, 6}, 3)

	fmt.Println(nums1)

}
