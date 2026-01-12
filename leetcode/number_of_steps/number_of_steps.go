package numberofsteps

func numberOfSteps(num int) int {
	// O(1) memory, O(number of bits) time using bit-count formula:
	if num == 0 {
		return 0
	}
	n := num
	bitLen := 0
	ones := 0
	for n > 0 {
		ones += n & 1
		bitLen++
		n >>= 1
	}
	return (bitLen - 1) + ones
}
