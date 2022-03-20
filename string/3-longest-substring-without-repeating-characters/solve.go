package main

import "fmt"

/**
  题目链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 */

// "abcabcbb"
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0{
		return 0
	}
	left := 0
	right := -1
	res := 1
	for left < len(s) {
		m := make(map[uint8]struct{})
		m[s[left]] = struct{}{}
		right = left + 1
		max := 1
		for right < len(s) {
			if _,ok := m[s[right]]; ok{
				break
			}
			m[s[right]] = struct{}{}
			right++
			max++
		}
		res = getMax(res, max)
		left++
	}
	return res
}

func getMax(src,target int) int {
	if src > target{
		return src
	}
	return target
}


func lengthOfLongestSubstring2(s string) int {
	// 记录字符是否出现过
	m := make(map[byte]int)
	length := len(s)
	res,right := 0, -1
	for index := 0; index < length; index++{
		if index != 0{
			delete(m, s[index-1])
		}
		for right + 1 < length && m[s[right+1]] == 0{
			m[s[right+1]]++
			right++
		}
		res = getMax(res, right - index + 1)
	}
	return res
}


func main() {
	s := "au"
	fmt.Println(lengthOfLongestSubstring2(s))
}