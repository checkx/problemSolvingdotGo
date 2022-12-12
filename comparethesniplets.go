func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	var aliceCount, bobCount int32
	for index := range a {
		if a[index] > b[index] {
			aliceCount++
		} else if b[index] > a[index] {
			bobCount++
		}
	}

	return []int32{aliceCount, bobCount}
}