package LeetCode

type LRUCache struct {
	IndexMap map[int]*DataNode
	DataList *DataList
	Capacity int
}

type DataNode struct {
	key   int
	value int
	next  *DataNode
	pre   *DataNode
}
type DataList struct {
	head *DataNode
	tail *DataNode
}

// 最先使用的放队尾  对头的是最旧的数据
func Constructor(capacity int) LRUCache {
	list := &DataList{
		head: &DataNode{},
		tail: &DataNode{},
	}
	list.head.next = list.tail
	list.tail.pre = list.head

	return LRUCache{
		IndexMap: make(map[int]*DataNode),
		DataList: list,
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	//1 获得index
	if d, ok := this.IndexMap[key]; !ok {
		return -1
	} else {
		//2 remove
		this.Remove(d)
		//3 移动
		this.AddTail(d)
		return d.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	//判断这个值是否存在 存在就移动到对位
	if d, ok := this.IndexMap[key]; ok {
		d.value = value
		this.Remove(d)
		this.AddTail(d)
	} else {
		//不存在就看看是否要淘汰 然后添加到队尾
		if len(this.IndexMap) == this.Capacity {
			oldNode := this.DataList.head.next
			delete(this.IndexMap, oldNode.key)
			this.Remove(oldNode)
			nowNode := &DataNode{
				key:   key,
				value: value,
			}
			this.IndexMap[key] = nowNode
			this.AddTail(nowNode)
		} else {
			nowNode := &DataNode{
				key:   key,
				value: value,
			}
			this.IndexMap[key] = nowNode
			this.AddTail(nowNode)
		}
	}
}

func (this *LRUCache) Remove(node *DataNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) AddTail(node *DataNode) {
	node.pre = this.DataList.tail.pre
	this.DataList.tail.pre.next = node
	node.next = this.DataList.tail
	this.DataList.tail.pre = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
