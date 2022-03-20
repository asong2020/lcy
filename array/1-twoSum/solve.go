package __twoSum

/**
  题目链接：https://leetcode-cn.com/problems/two-sum/
 */

/*
 O(n2)解法
 */
func twoSum0(nums []int, target int) []int {
	res := make([]int,0)
	for index := 0; index < len(nums); index++{
		for j := index+1; j < len(nums); j++{
			if nums[index] + nums[j] == target{
				res = append(res, index, j)
			}
		}
	}
	return res
}

/*
	O(n)解法
    使用map存储，key：值  value: index
 */

func twoSum1(nums []int, target int) []int {
	m := make(map[int]int, 10)
	for i, v := range nums{
		if p, ok := m[target - v]; ok{
			return []int{i,p}
		}
		m[v] = i
	}
	return nil
}