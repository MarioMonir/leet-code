package main

func differenceOfSum(nums []int) int {
	eSum, dSum := 0, 0
	for _, el := range nums {
		eSum += el
		for el > 0 {
			dSum += el % 10
			el /= 10
		}
	}
	return eSum - dSum
}
