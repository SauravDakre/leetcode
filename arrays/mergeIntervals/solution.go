package mergeIntervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Sort(it(intervals))
	merged := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		// fmt.Println(intervals[i])
		if len(merged) == 0 || merged[len(merged)-1][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
			// fmt.Println("added", merged)
		} else {
			if merged[len(merged)-1][1] < intervals[i][1] {
				merged[len(merged)-1][1] = intervals[i][1]
			}
			// fmt.Println("edited", merged)
		}
	}
	return merged
}

type it [][]int

func (a it) Len() int           { return len(a) }
func (a it) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a it) Less(i, j int) bool { return a[i][0] < a[j][0] }
