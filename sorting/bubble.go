package sorting

// BubbleSort largest elements bubble to the top
// in-place
func BubbleSort(A []int) []int {
	n := len(A)
	for {
		if n <= 1 {
			break
		}
		newn := 0
		for i := 1; i < n; i++ {
			// if the pair is out of order
			if A[i-1] > A[i] {
				// swap
				A[i-1], A[i] = A[i], A[i-1]
				// remember pos
				newn = i
			}
		}
		n = newn
	}
	return A
}
