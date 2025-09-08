package LeetCode

import (
	"fmt"
	"slices"
)

type Router struct {
	packageCount   int
	maxCount       int
	packageList    []*packageInfo
	destinationMap map[int][]int //destination -> []timeStamp
	distinctMap    map[string]bool
}
type packageInfo struct {
	source      int
	destination int
	timestamp   int
}

func Constructor(memoryLimit int) Router {
	return Router{
		packageCount:   0,
		maxCount:       memoryLimit,
		packageList:    make([]*packageInfo, 0),
		destinationMap: make(map[int][]int),
		distinctMap:    make(map[string]bool),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	complexName := getComplexString(source, destination, timestamp)
	if this.distinctMap[complexName] {
		return false
	}
	if this.packageCount >= this.maxCount {
		this.ForwardPacket()
	}
	info := &packageInfo{
		source:      source,
		destination: destination,
		timestamp:   timestamp,
	}
	this.distinctMap[complexName] = true
	this.packageList = append(this.packageList, info)
	this.destinationMap[destination] = append(this.destinationMap[destination], timestamp)
	this.packageCount++
	return true
}

func (this *Router) ForwardPacket() []int {
	if this.packageCount == 0 {
		return []int{}
	}
	first := this.packageList[0]
	destination := first.destination
	this.destinationMap[destination] = this.destinationMap[destination][1:]
	this.packageList = this.packageList[1:]
	this.packageCount--
	complexName := getComplexString(first.source, destination, first.timestamp)
	delete(this.distinctMap, complexName)
	return []int{first.source, destination, first.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	arr := this.destinationMap[destination]
	l, _ := slices.BinarySearch(arr, startTime)
	r, _ := slices.BinarySearch(arr, endTime+1)
	return r - l + 1
}
func getComplexString(source int, destination int, timestamp int) string {
	return fmt.Sprint("%d-%d-%d", source, destination, timestamp)
}
