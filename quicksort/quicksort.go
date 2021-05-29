package quicksort

func QuickSort(a []int) {
	//exit the recursion if the given range have less than 2 elements
	if len(a) <= 1 {
		return
	}
    	//find the correct position of the pivot element (all element in the left part of the pivot - smaller, in the right - greater)
	pivotPosition := partition(a)

	//recursively call the save process for both sides (pivot will not be affected)
	QuickSort(a[:pivotPosition])
	QuickSort(a[pivotPosition+1:])
}

func partition(a []int) int {
	//choose like a pivot the last element of the array
	pivot := a[len(a)-1]
	//index of the most left element that will be swapped with the current one (at index 'i'), in the case if the current element is smaller that pivot
	k := -1
	for i:=0;i<len(a);i++ {
		if a[i] < pivot {
			k++
			a[k], a[i] = a[i], a[k]
		}
	}

	//swap the pivot with the element at the index k+1 (because all the elements till index 'k' are smaller than the pilot)
	a[k+1], a[len(a)-1] = a[len(a)-1], a[k+1]
	return k+1
}
