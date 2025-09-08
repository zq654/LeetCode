package LeetCode

import "slices"

// "aaaaaaasgddgfdgdgsgsssssssaaaaaa"
func maximumLength(s string) int {
	check := func(char rune, length int) bool {
		count := 0
		tempLength := 0
		for _, v := range s {
			if v == char {
				tempLength++
			} else {
				if tempLength >= length {
					count += tempLength - length + 1
					if count >= 3 {
						return true
					}
				}
				tempLength = 0
			}
		}
		if tempLength >= length {
			count += tempLength - length + 1
			if count >= 3 {
				return true
			}
		}
		return false
	}
	charMap := make([]int, 26)
	for _, v := range s {
		charMap[v-'a']++
	}
	res := -1
	for char, charNums := range charMap {
		if charNums >= 3 {
			l := 1
			r := charNums - 2
			for l <= r {
				mid := l + (r-l)>>1
				if check(rune('a'+char), mid) {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			res = max(res, l-1)
		}
	}
	return res
}

func maximumLength_2(s string) int {
	runeList := [26][]int{}
	tempLength := 1
	preRune := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == preRune {
			tempLength++
		}
		if s[i] != preRune {
			runeList[s[i-1]-'a'] = append(runeList[s[i-1]-'a'], tempLength)
			tempLength = 1
			preRune = s[i]
		}
	}
	runeList[s[len(s)-1]-'a'] = append(runeList[s[len(s)-1]-'a'], tempLength)
	res := -1
	for _, v := range runeList {
		if len(v) == 0 {
			continue
		}
		v = append(v, -1, -1)
		slices.SortFunc(v, func(a, b int) int {
			return b - a
		})
		res = max(res, v[0]-2, min(v[0]-1, v[1]), v[2])
	}
	if res == 0 {
		return -1
	}
	return res
}
