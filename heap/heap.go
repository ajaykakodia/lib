package heap

type heapNode struct {
	value    string
	priority int
}

type Heap struct {
	arr []*heapNode
}

func NewPriorityQueue() *Heap {
	return &Heap{
		arr: []*heapNode{},
	}
}

func (h *Heap) Size() int {
	return len(h.arr)
}

func (h *Heap) IsEmpty() bool {
	return h.Size() == 0
}

func (h *Heap) Insert(value string, priority int) {
	node := &heapNode{
		value:    value,
		priority: priority,
	}
	h.arr = append(h.arr, node)
	heapify(h)
}

func (h *Heap) Get() string {
	if h.IsEmpty() {
		return ""
	}
	return h.arr[0].value
}

func (h *Heap) Remove() string {
	if h.IsEmpty() {
		return ""
	}
	removedEle := h.arr[0]

	h.arr[0] = h.arr[h.Size()-1]
	h.arr = h.arr[:h.Size()-1]
	downHeapify(h)
	return removedEle.value
}

func (h *Heap) Replace(value string, priority int) string {
	if h.IsEmpty() {
		return ""
	}
	node := &heapNode{
		value:    value,
		priority: priority,
	}
	removedEle := h.arr[0]
	h.arr[0] = node
	downHeapify(h)
	return removedEle.value
}

// Place root node to its proper place as per heap order property
func downHeapify(h *Heap) {
	if h.IsEmpty() {
		return
	}
	size := h.Size()
	parentIndex, minIndex := 0, 0
	lChild := parentIndex*2 + 1
	rChild := parentIndex*2 + 2

	for lChild < size {
		minIndex = parentIndex
		if h.arr[minIndex].priority > h.arr[lChild].priority {
			minIndex = lChild
		}
		if rChild < size && h.arr[minIndex].priority > h.arr[rChild].priority {
			minIndex = rChild
		}
		if minIndex == parentIndex {
			break
		}
		h.arr[minIndex], h.arr[parentIndex] = h.arr[parentIndex], h.arr[minIndex]
		parentIndex = minIndex
		lChild = parentIndex*2 + 1
		rChild = parentIndex*2 + 2
	}
}

// Place last array index to its proper index as per min heap order property
func heapify(h *Heap) {
	size := h.Size()
	childNode := size - 1

	for childNode > 0 {
		parentNodeIndex := (childNode - 1) / 2
		if h.arr[parentNodeIndex].priority > h.arr[childNode].priority {
			h.arr[parentNodeIndex], h.arr[childNode] = h.arr[childNode], h.arr[parentNodeIndex]
			childNode = parentNodeIndex
			continue
		}
		break
	}
}
