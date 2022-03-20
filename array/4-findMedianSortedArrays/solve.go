package main

import "fmt"

/**
 题目链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
 */


// 时间复杂度O(m+n)做法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	index1, index2 := 0, 0
	length1, length2 := len(nums1), len(nums2)
	sumArray := make([]int, 0, length1 + length2)
	var ans float64
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			sumArray = append(sumArray, nums1[index1])
			index1++
		}else {
			sumArray = append(sumArray, nums2[index2])
			index2++
		}
	}
	for index1 < length1 {
		sumArray = append(sumArray, nums1[index1])
		index1++
	}
	for index2 < length2 {
		sumArray = append(sumArray, nums2[index2])
		index2++
	}
	mindex := len(sumArray) / 2
	if len(sumArray) % 2 == 0 {
		ans = (float64(sumArray[mindex-1]) + float64(sumArray[mindex])) / 2
	}else {
		ans = float64(sumArray[mindex])
	}
	return ans
}


func main()  {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
