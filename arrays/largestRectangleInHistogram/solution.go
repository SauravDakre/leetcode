package largestRectangleInHistogram

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	max := 0
	area := 0
	i := 0
	for i < len(heights) {
		// fmt.Println(stack, i)
		if len(stack) == 0 || heights[stack[len(stack)-1]] < heights[i] {
			stack = append(stack, i)
			i++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) != 0 {
				width = (i - stack[len(stack)-1] - 1)
			}
			area = heights[top] * width

			if max < area {
				max = area
			}
			// fmt.Println(stack, i, width, area, max)
		}
	}
	// fmt.Println(stack, i)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		width := i
		if len(stack) != 0 {
			width = (i - stack[len(stack)-1] - 1)
		}
		area = heights[top] * width

		if max < area {
			max = area
		}
		// fmt.Println(stack, i, width, area, max)
	}
	return max
}
