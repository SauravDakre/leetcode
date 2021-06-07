package courseScheduleII

func findOrder(numCourses int, prerequisites [][]int) []int {

	set := make(map[int]bool)
	res := make([]int, 0)
	adjacencyList := make(map[int]*[]int)

	// fmt.Println("pr", prerequisites)
	for _, edge := range prerequisites {
		// fmt.Println("e", edge)
		if ar, ok := adjacencyList[edge[0]]; ok {
			*ar = append(*ar, edge[1])
		} else {
			adjacencyList[edge[0]] = &[]int{edge[1]}
		}
		set[edge[0]] = true
		set[edge[1]] = true
	}
	// fmt.Println("adjacencyList",adjacencyList)
	// fmt.Println("set", set)
	indegree := make(map[int]int)
	for k, _ := range set {
		indegree[k] = 0
	}
	for _, ar := range adjacencyList {
		for _, v := range *ar {
			indegree[v]++
		}

	}
	// fmt.Println("inde", indegree)
	stack := []int{}
	zeroInDeg := []int{}
	for k, v := range indegree {
		if v == 0 {
			zeroInDeg = append(zeroInDeg, k)
			delete(indegree, k)
		}
	}
	// fmt.Println("zi",zeroInDeg)
	for len(zeroInDeg) != 0 {
		firstEl := zeroInDeg[0]
		stack = append(stack, firstEl)
		zeroInDeg = zeroInDeg[1:]
		if ar, ok := adjacencyList[firstEl]; ok {
			for _, t := range *ar {
				indegree[t]--
				if indegree[t] == 0 {
					zeroInDeg = append(zeroInDeg, t)
					delete(indegree, t)
				}
			}
		}

		// fmt.Println("inde", indegree)
		// fmt.Println("zi",zeroInDeg)
	}
	if len(indegree) != 0 {
		return res
	}
	if numCourses != len(stack) {
		diff := numCourses - len(set)
		for i := numCourses - 1; diff > 0; i = i - 1 {
			if _, ok := set[i]; !ok {
				res = append(res, i)
				diff--
			}

		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}
	return res
}
