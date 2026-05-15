type Item struct {
    Frequency int
    Number int
}

func topKFrequent(nums []int, k int) []int {
    freqMap := make(map[int]int)
    for _, x := range nums{
        freqMap[x]++
    }

    var arr []Item
    for number , frequency := range freqMap{
        arr = append(arr, Item{
            Frequency : frequency,
            Number : number,
        })
    }

    sort.Slice(arr, func (i,j int)bool {
        return arr[i].Frequency > arr[j]. Frequency
    })
    
    fmt.Println (arr)
    var result []int
    for x := range k {
        result = append(result, arr[x].Number)
    }
    return result    
}
