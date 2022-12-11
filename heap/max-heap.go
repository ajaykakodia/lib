package heap

import "errors"

// HeapifyMax => take array as input and create max priority queue
func HeapifyMax(arr []int) {
	size := len(arr)
	nonLeafNodes := size / 2
	for i := nonLeafNodes - 1; i >= 0; i-- {
		downHeapifyMaxArray(arr, i)
	}
}

// HeapMaxPush => insert data into PQ and return PQ array
func HeapMaxPush(arr []int, data int) []int {
	arr = append(arr, data)
	heapifyMaxArray(arr)
	return arr
}

// HeapMaxPop => remove data from PQ
func HeapMaxPop(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array is empty")
	}
	popData := arr[0]
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	downHeapifyMaxArray(arr, 0)
	return popData, nil
}

// HeapMaxReplace => remove data from PQ and replace with existing one
func HeapMaxReplace(arr []int, data int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array is empty")
	}
	popData := arr[0]
	arr[0] = data
	downHeapifyMaxArray(arr, 0)
	return popData, nil
}

func heapifyMaxArray(arr []int) {
	cIndex := len(arr) - 1
	var pIndex int
	for cIndex > 0 {
		pIndex = (cIndex - 1) / 2
		if arr[pIndex] < arr[cIndex] {
			arr[pIndex], arr[cIndex] = arr[cIndex], arr[pIndex]
			cIndex = pIndex
			continue
		}
		return
	}
}

func downHeapifyMaxArray(arr []int, startIndex int) {
	size := len(arr)
	if size == 0 {
		return
	}
	pIndex, maxIndex, lcIndex, rcIndex := startIndex, startIndex, startIndex*2+1, startIndex*2+2
	for lcIndex < size {
		maxIndex = pIndex
		if arr[maxIndex] < arr[lcIndex] {
			maxIndex = lcIndex
		}
		if rcIndex < size && arr[maxIndex] < arr[rcIndex] {
			maxIndex = rcIndex
		}
		if maxIndex == pIndex {
			return
		}
		arr[pIndex], arr[maxIndex] = arr[maxIndex], arr[pIndex]
		pIndex = maxIndex
		lcIndex = pIndex*2 + 1
		rcIndex = pIndex*2 + 2
	}
}
