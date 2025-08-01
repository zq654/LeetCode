package LeetCode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	//自底向上归并排序
	length := 0
	for tempNode := head; tempNode != nil; tempNode = tempNode.Next {
		length++
	}
	step := 1
	resNode := &ListNode{}
	preNode := resNode
	for step < length {
		for leftNode := head; leftNode != nil; {
			rightNode := leftNode
			for i := 0; i < step-1 && rightNode.Next != nil; i++ {
				rightNode = rightNode.Next
			}
			rightNode, rightNode.Next = rightNode.Next, nil
			nextNode := rightNode
			for i := 0; i < step-1 && nextNode.Next != nil; i++ {
				nextNode = nextNode.Next
			}
			nextNode, nextNode.Next = nextNode.Next, nil
			preNode.Next = mergeSort(leftNode, rightNode)
			for preNode.Next != nil {
				preNode = preNode.Next
			}
			leftNode = nextNode
		}
		step = step + step
	}
	return resNode.Next
}

// 俩链表合并
func mergeSort(leftHead *ListNode, rightHead *ListNode) *ListNode {
	preNode := &ListNode{}
	headNode := preNode
	for leftHead != nil && rightHead != nil {
		if leftHead.Val < rightHead.Val {
			headNode.Next = leftHead
			leftHead = leftHead.Next
			headNode = headNode.Next
			headNode.Next = nil
		} else {
			headNode.Next = rightHead
			rightHead = rightHead.Next
			headNode = headNode.Next
			headNode.Next = nil
		}
	}
	if leftHead != nil {
		headNode.Next = leftHead
	}
	if rightHead != nil {
		headNode.Next = rightHead
	}
	return preNode.Next
}
