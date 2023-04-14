func runningSum(nums []int) []int {
    r := make([]int, len(nums))
    r[0] = nums[0]
    for i:=1; i<len(nums); i++ {
        r[i] = nums[i] + r[i-1]
    }
    return r
}
