func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var mainDiagonal, secondaryDiagonal int32
	// Get the size of the array
	size := len(arr)
	// Iterate through the array
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			// Check if the current element is on the main diagonal
			if i == j {
				mainDiagonal += arr[i][j]
			}
			// Check if the current element is on the secondary diagonal
			if i+j == size-1 {
				secondaryDiagonal += arr[i][j]
			}
		}
	}
	// Return the absolute difference between the sums of the main and secondary diagonals
	return int32(math.Abs(float64(mainDiagonal - secondaryDiagonal)))
}