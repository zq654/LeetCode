package LeetCode

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	//S是长串 t是待匹配短串
	firstChar := t[0]
	//先知道s中的一个为t开头的子串
	beginIndex := -1
	for i := 0; i < len(s); i++ {
		if s[i] == firstChar {
			beginIndex = i
			break
		}
	}
	if beginIndex == -1 {
		return 0
	}
	//用一个arr记录子串中每个字符出现的位置 字符包含大写字符
	locationArr := make([][]int, 52)
	for i := 0; i < len(t); i++ {
		tCharNum := t[i] - 'a'
		if tCharNum > 25 {
			tCharNum = t[i] - 'A' + 26
		}
		locationArr[tCharNum] = append(locationArr[tCharNum], i)
	}
	//用一个arr记录t字符串中[0:i]的字串出现的次数
	resultArr := make([]int, len(t))
	resultArr[0] = 1
	for i := beginIndex + 1; i < len(s); i++ {
		//s的当前字符转化成数字
		sCharNum := s[i] - 'a'
		if sCharNum > 25 {
			sCharNum = s[i] - 'A' + 26
		}
		//沦陷的当前字符在t中的位置
		locations := locationArr[sCharNum]
		//更新结果数组,倒叙更新是防止某一个值被重复计算
		for j := len(locations) - 1; j >= 0; j-- {
			if locations[j] == 0 {
				resultArr[locations[j]]++
			} else {
				resultArr[locations[j]] += resultArr[locations[j]-1]
			}
		}
	}
	return resultArr[len(t)-1]
}
