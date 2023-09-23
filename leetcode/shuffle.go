package leetcode 

func restoreString(s string, indices []int) string {
	result := make([]rune, len(s))
	for idx, val := range s {
		result[indices[idx]] = val
	}
	return string(result)
}
