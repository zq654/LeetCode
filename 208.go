package LeetCode

type Node struct {
	son [26]*Node
	end bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{&Node{}}
}

func (t Trie) Insert(word string) {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil { // 无路可走？
			cur.son[c] = &Node{} // 那就造路！
		}
		cur = cur.son[c]
	}
	cur.end = true
}

func (t Trie) find(word string) int {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil { // 道不同，不相为谋
			return 0
		}
		cur = cur.son[c]
	}
	// 走过同样的路（2=完全匹配，1=前缀匹配）
	if cur.end {
		return 2
	}
	return 1
}

func (t Trie) Search(word string) bool {
	return t.find(word) == 2
}

func (t Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != 0
}
