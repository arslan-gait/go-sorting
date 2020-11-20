package sorting

// InsertionSort pick one element and find its place
// in-place
func InsertionSort(A []int) []int {
	for i := 1; i < len(A); i++ {
		j := i
		for j > 0 {
			if A[j-1] > A[j] {
				A[j-1], A[j] = A[j], A[j-1]
			}
			j = j - 1
		}
	}
	return A
}