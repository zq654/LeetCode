package LeetCode

import "strconv"

var strType = 1
var intType = 2

func decodeString(s string) string {
	//用两个栈来维护
	//遇到数字就入栈
	//遇到 “[” 就将 “[” 与 “]” 中间的字符串入栈
	//遇到 “]” 就取字符串栈中的字符串for 数字栈顶次
	//如果字符串栈中还有东西就入栈 否则就拼接到结果中
	//如果数字栈里面没有东西就直接拼接到结果中
	runes := []rune(s)
	res := ""
	numStack := make([]int, 0)
	stringStack := make([]string, 0)
	leftIndex := 0  //标记左括号的下一位
	rightIndex := 0 //标记右括号的下一位
	for i := 0; i < len(s); i++ {
		//“[”之前肯定是int
		if runes[i] == '[' {
			num, _ := strconv.Atoi(s[rightIndex:i])
			numStack = append(numStack, num)
			leftIndex = i + 1
			continue
		}
		//"]"之前肯定是string
		if runes[i] == ']' {
			subStr := s[leftIndex:i]
			if len(numStack) > 1 {
				subStr = stringStack[len(stringStack)-1]
				stringStack = stringStack[:len(stringStack)-1]
			}
			tempStr := getRepeatString(numStack[len(numStack)-1], subStr)
			numStack = numStack[:len(numStack)-1]
			rightIndex = i + 1
			if len(numStack) == 0 {
				res = res + tempStr
			} else {
				stringStack = append(stringStack, tempStr)
			}
			continue
		}

		if getRuneTypeEnum(runes[leftIndex]) == strType && len(numStack) == 0 {
			res += string(s[i])
		}
	}
	return res
}

func checkRuneIsInt(char rune) bool {
	return char-'0' >= 0 && char-'0' <= 9
}
func getRuneTypeEnum(char rune) int {
	if checkRuneIsInt(char) {
		return intType
	}
	return strType
}

func getRepeatString(repeat int, subStr string) string {
	res := ""
	for i := 0; i < repeat; i++ {
		res += subStr
	}
	return res
}
