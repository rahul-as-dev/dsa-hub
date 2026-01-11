package arraysstrings

// [12,1,4,10,7,3,2,6]
func NextGreaterElementToRight(arr []int) []int {
	size := len(arr)
	ans := make([]int, size)
	ans[size-1] = -1
	stack := make([]int, 0, size)
	stack = append(stack, arr[size-1])
	for i:=size-2; i>=0; i-- {
		if arr[i] < stack[len(stack)-1] {
			ans[i] = stack[len(stack)-1]
			stack = append(stack, arr[i])
			continue
		}
		for len(stack) > 0 && arr[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1]
		}else {
			ans[i] = -1
		}
		stack = append(stack, arr[i])
	}
	return ans
}