package sorting

// ShakerSort is same as bubble sort, but operates from left to right and vice versa
// in-place
func ShakerSort(A []int) []int {
	swapped := false
	for {
		for i := 0; i < len(A)-1; i++ {
			if A[i] > A[i+1] {
				A[i], A[i+1] = A[i+1], A[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		swapped = false
		for i := len(A) - 2; i > 0; i-- {
			if A[i] > A[i+1] {
				A[i], A[i+1] = A[i+1], A[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return A
}
