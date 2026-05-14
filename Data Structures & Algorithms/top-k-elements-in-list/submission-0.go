type Item struct {
    Num int
    Freq int
    
}
type MaxHeap []Item

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].Freq > h[j].Freq
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
 
    freq := make(map[int]int)
    for _, x := range nums {
        freq[x]++
    }
    fmt.Println(freq)
    h := &MaxHeap{}
    heap.Init(h)
  
    for key, val := range freq {
        heap.Push(h, Item{Freq: val, Num : key})
    }
    var result []int
    for x := range k {
        fmt.Println(x)
        item := heap.Pop(h)
        result = append(result,item.(Item).Num)
    }

    
    return result


}
