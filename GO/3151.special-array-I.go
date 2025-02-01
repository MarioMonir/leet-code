package main

func isArraySpecial(nums []int) bool {
	isEven := func(n int) bool {
		return n%2 == 0
	}

	isTwoSuccessiveEvenNumbers := func(n1 int, n2 int) bool {
		return isEven(n1) && isEven(n2)
	}

	isTwoSuccessiveOddNumbers := func(n1 int, n2 int) bool {
		return !isEven(n1) && !isEven(n2)
	}

	for i := 0; i < len(nums)-1; i++ {
		if isTwoSuccessiveEvenNumbers(nums[i], nums[i+1]) ||
			isTwoSuccessiveOddNumbers(nums[i], nums[i+1]) {
			return false
		}
	}
	return true
}
