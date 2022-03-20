package main

import "fmt"

func FindDiffNum(nums []int) int {
	length := len(nums)
	left := 0
	right := length - 1
	sum := 0
	for left <= right {
		if nums[left] == nums[right]{
			sum++
			left++
			right--
		}else if nums[left] < nums[right]{
			sum++
			left++
		}else {
			sum++
			right--
		}
	}
	return sum
}


func main()  {
	array1 := []int{1, 2, 3, 2, 1}
	array2 := []int{1,2,8,7,6,5,3,2,1}
	fmt.Println(FindDiffNum(array1))
	fmt.Println(FindDiffNum(array2))
}
