package sorting


// MergeSort recursive solution
func MergeSort(A []int) []int {

	if len(A) < 2 {
		return A
	}

	mid := (len(A)) / 2

	// recursively merge left and right parts
	return Merge(MergeSort(A[:mid]), MergeSort(A[mid:]))
}

// Merge helper for main func
func Merge(L, R []int) []int {

	n, i, j := len(L)+len(R), 0, 0
	A := make([]int, n, n)

	for k := 0; k < n; k++ {
		// left is empty
		if i > len(L)-1 && j <= len(R)-1 {
			A[k] = R[j]
			j++
			// right is empty
		} else if j > len(R)-1 && i <= len(L)-1 {
			A[k] = L[i]
			i++
			// pick from left
		} else if L[i] < R[j] {
			A[k] = L[i]
			i++
			// pick from right
		} else {
			A[k] = R[j]
			j++
		}
	}
	return A
}