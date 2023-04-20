package leetcode

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(nums))
	for idx := range nums {
		movedIndx := index[idx]
		if idx != movedIndx {
			for mdx := idx; mdx > movedIndx; mdx-- {
				target[mdx] = target[mdx-1]
			}
		}
		target[movedIndx] = nums[idx]
	}
	return target
}
