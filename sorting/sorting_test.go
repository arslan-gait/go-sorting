package sorting

import (
	"math/rand"
	"testing"
	"time"
)

// number of elements
var N int = 1000

func generateSlice(N int) []int {

	slice := make([]int, N, N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		slice[i] = rand.Intn(N) - rand.Intn(N)
	}
	return slice
}

func isSorted(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	A := BubbleSort(generateSlice(N))
	if result := isSorted(A); !result {
		t.Errorf("error \n%v\n", A)
	}
}

func TestShakerSortRandom(t *testing.T) {
	A := ShakerSort(generateSlice(N))
	if result := isSorted(A); !result {
		t.Errorf("error \n%v\n", A)
	}
}

func TestShakerSortSorted(t *testing.T) {
	A := ShakerSort([]int{1,2,3,4,5})
	if result := isSorted(A); !result {
		t.Errorf("error \n%v\n", A)
	}
}

func TestInsertionSort(t *testing.T) {
	A := InsertionSort(generateSlice(N))
	if result := isSorted(A); !result {
		t.Errorf("error \n%v\n", A)
	}
}

func TestSelectionSort(t *testing.T) {
	A := SelectionSort(generateSlice(N))
	if result := isSorted(A); !result {
		t.Errorf("error \n%v\n", A)
	}
}

func TestMergeSort(t *testing.T) {
	A := MergeSort(generateSlice(N))
	if result := isSorted(A); !result {
		t.Errorf("error \n%v\n", A)
	}
}

func TestQuickSort(t *testing.T) {
	A := QuickSort(generateSlice(N))
	if result := isSorted(A); !result {
		t.Errorf("error \n%v\n", A)
	}
}
