package main

// getMaxSubarray
// The main point of the algorithm is:
// If we know max sum of arrays with ending on index i - 1,
// we can easily find max sum of arrays with ending on index i.
// If current element can `improve` sum - we add it to current sum (subSum).
// If not - we start new subSum from current element.
// `Improvement` is checked by comparison of current element and `possible` sum with subSum
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	rollingSum := nums[0]

	for idx, val := range nums {
		if idx == 0 { // skip first element cuz we already passed it on line 8
			continue
		}

		if rollingSum+val > val {
			rollingSum = rollingSum + val
		} else {
			rollingSum = val
		}
		if rollingSum > maxSum {
			maxSum = rollingSum
		}
	}

	return maxSum
}
