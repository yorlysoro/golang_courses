package runningsum1darray

import "errors"

func runningsum(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("input slice is empty")
	}

	if len(nums) > 1000 {
		return nil, errors.New("input slice exceeds maximum length of 1000")
	}
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
		if nums[i] < -10^6 || nums[i] > 10^6 {
			return nil, errors.New("element out of allowed range [-10^6, 10^6]")
		}
	}
	return nums, nil
}
