package LeetCode

import "strings"

func groupAnagrams1(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, str := range strs {
		byteArr := make([]int, 26)
		bytes := []byte(str)
		for _, b := range bytes {
			byteArr[b-'a']++
		}
		builder := strings.Builder{}
		for i, v := range byteArr {
			for j := 0; j < v; j++ {
				builder.WriteByte(byte(i + 'a'))
			}
		}
		strMap[builder.String()] = append(strMap[builder.String()], str)
	}
	res := make([][]string, 0)
	for _, v := range strMap {
		res = append(res, v)
	}
	return res
}

// 看了别人的代码第二版 直接用int数组当map的key不用重新用组合用string当key
func groupAnagrams2(strs []string) [][]string {
	strMap := make(map[[26]int][]string)
	for _, str := range strs {
		byteArr := [26]int{}
		bytes := []byte(str)
		for _, b := range bytes {
			byteArr[b-'a']++
		}

		strMap[byteArr] = append(strMap[byteArr], str)
	}
	res := make([][]string, 0)
	for _, v := range strMap {
		res = append(res, v)
	}
	return res
}
