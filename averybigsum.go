func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var sum int64

	// Iterate through the array and add each element to the running total
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}

	return sum
}