package twosum

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i,t := range nums {
        if ind, ok := m[target - t]; ok {
            return []int{ind, i}
        }else{
            m[t] = i
        }
    }
    return []int{}
}