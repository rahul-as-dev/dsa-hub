package arraysstrings

// [12,1,4,10,7,3,2,6]
func NextGreaterElementToRight(arr []int) []int {
	size := len(arr)
	if size == 0 {
		return []int{}
	}
	ans := make([]int, size)
	for i := range size {
		ans[i] = -1
	}
	stack := make([]int, 0, size)
	for i:=size-1; i>=0; i-- {
		for len(stack) > 0 && arr[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	return ans
}