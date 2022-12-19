func staircase(n int32) {
	// Write your code here
	for i := int32(1); i <= n; i++ {
		for j := int32(1); j <= n; j++ {
			if j <= n-i {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}