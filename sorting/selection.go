package sorting

// SelectionSort every time select min item and move to correct pos
// in-place
func SelectionSort(A []int) []int {
	for i := 0; i < len(A); i++ {
		minIdx := i
		for j := i; j < len(A); j++ {
			if A[j] < A[minIdx] {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
	return A
}