package LeetCode

type TimeMap map[string][]storeData

type storeData struct {
	timeStamp int
	data      string
}

func Constructor() TimeMap {
	return map[string][]storeData{}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	(*this)[key] = append((*this)[key], storeData{
		timeStamp: timestamp,
		data:      value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	//二分查找 <=timestamp 的最大值 即找 <timestamp+1 的最大值
	dataArr := (*this)[key]
	index := findValue(dataArr, timestamp+1)
	if index == -1 {
		return ""
	}
	return dataArr[index].data
}

func findValue(arr []storeData, target int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid].timeStamp >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
