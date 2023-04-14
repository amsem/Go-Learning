func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    
    for idx, num := range nums {
        complement := target - num
        
        if c, ok := m[complement]; ok {
            return []int{c, idx}
        }
        m[num] = idx
    }
    return []int{}    
}
