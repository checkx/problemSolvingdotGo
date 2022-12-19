func plusMinus(arr []int32) {
	var pos, neg, zero float32
	n := float32(len(arr))

	for _, v := range arr {
		switch {
		case v > 0:
			pos++
		case v < 0:
			neg++
		default:
			zero++
		}
	}

	fmt.Printf("%.6f\n", pos/n)
	fmt.Printf("%.6f\n", neg/n)
	fmt.Printf("%.6f\n", zero/n)
}