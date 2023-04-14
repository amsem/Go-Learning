func ContainsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for idx := 0; idx < len(nums)-1; idx++ {
        if nums[idx] == nums[idx+1] {
            return true
        }
    }
    return false
}

//used the sort package, so make sure you import it when you use this code
