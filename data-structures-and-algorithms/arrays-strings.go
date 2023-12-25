package main

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	var sq int
	newNums := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[right]) >= abs(nums[left]) {
			sq = nums[right]
			right--
		} else {
			sq = nums[left]
			left++
		}
		newNums[i] = sq * sq
	}

	return newNums
}
