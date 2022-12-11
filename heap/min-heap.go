package heap

import "errors"

// Heapify => take array as input and create min priority queue
func Heapify(arr []int) {
	size := len(arr)
	nonLeafNodes := size / 2
	for i := nonLeafNodes - 1; i >= 0; i-- {
		downHeapifyArray(arr, i)
	}
}

// HeapPush => insert data into PQ and return PQ array
func HeapPush(arr []int, data int) []int {
	arr = append(arr, data)
	heapifyArray(arr)
	return arr
}

// HeapPop => remove data from PQ
func HeapPop(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array is empty")
	}
	popData := arr[0]
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	downHeapifyArray(arr, 0)
	return popData, nil
}

// HeapReplace => remove data from PQ and replace with existing one
func HeapReplace(arr []int, data int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array is empty")
	}
	popData := arr[0]
	arr[0] = data
	downHeapifyArray(arr, 0)
	return popData, nil
}

func heapifyArray(arr []int) {
	cIndex := len(arr) - 1
	var pIndex int
	for cIndex > 0 {
		pIndex = (cIndex - 1) / 2
		if arr[pIndex] > arr[cIndex] {
			arr[pIndex], arr[cIndex] = arr[cIndex], arr[pIndex]
			cIndex = pIndex
			continue
		}
		return
	}
}

func downHeapifyArray(arr []int, startIndex int) {
	size := len(arr)
	if size == 0 {
		return
	}
	pIndex, minIndex, lcIndex, rcIndex := startIndex, startIndex, startIndex*2+1, startIndex*2+2
	for lcIndex < size {
		minIndex = pIndex
		if arr[minIndex] > arr[lcIndex] {
			minIndex = lcIndex
		}
		if rcIndex < size && arr[minIndex] > arr[rcIndex] {
			minIndex = rcIndex
		}
		if minIndex == pIndex {
			return
		}
		arr[pIndex], arr[minIndex] = arr[minIndex], arr[pIndex]
		pIndex = minIndex
		lcIndex = pIndex*2 + 1
		rcIndex = pIndex*2 + 2
	}
}
