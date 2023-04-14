func pivotIndex(nums []int) int {
    i := 0
    j := 0
    for _, v := range nums {
        i +=v
    }
    for x,v := range nums {
        i -= v
        if i == j {
            return x
        }
        j += v
    }
    return -1
    
}
