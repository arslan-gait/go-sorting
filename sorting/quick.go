package sorting

// QuickSort partition input array into 2 parts and sort them recursively
// in-place
func QuickSort(A []int) []int {
	length := len(A)

	if length < 2 {
		return A
	}

	L, R := 0, length-1
	pivot := (L + R) / 2

	// swap pivot element with the rightmost one
	A[R], A[pivot] = A[pivot], A[R]

	// swap elements around pivot
	// such that el on left are smaller
	// and on right are larger
	for i := range A {
		if A[i] < A[R] {
			A[L], A[i] = A[i], A[L]
			L++
		}
	}

	// move pivot on its right position
	A[L], A[R] = A[R], A[L]

	// call quicksort recursively on left and right parts
	QuickSort(A[:L])
	QuickSort(A[L+1:])

	return A
}
