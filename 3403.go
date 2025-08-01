package LeetCode

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	//贪心找到字符序最大的的字符 然后截取他尽可能得到的最大字符串
	wordArr := []byte(word)
	length := len(wordArr)
	//找到最大的字符的索引值
	index := 0
	for i := 1; i < length; i++ {
		if wordArr[i] > wordArr[index] {
			index = i
		}
		//如果相等则比较一下后续
		if wordArr[i] == wordArr[index] {
			for j := 1; j+i < length; j++ {
				if wordArr[i+j] > wordArr[index+j] {
					index = i
					break
				} else if wordArr[i+j] < wordArr[index+j] {
					break
				}
			}
		}
	}
	//算出最大能得到的子串长度
	subStrLength := length + 1 - numFriends
	//算出一直到结尾的索引值
	endIndex := min(index+subStrLength, length)
	return string(wordArr[index:endIndex])

}
