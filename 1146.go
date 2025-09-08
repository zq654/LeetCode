package LeetCode

type SnapshotArray struct {
	infoArr [][]snapshotInfo
	nowSnap int
}

type snapshotInfo struct {
	snapID int
	value  int
}

func Constructor(length int) SnapshotArray {
	temp := make([][]snapshotInfo, length)
	for i := 0; i < length; i++ {
		temp[i] = make([]snapshotInfo, 0)
	}
	res := SnapshotArray{
		infoArr: temp,
		nowSnap: -1,
	}
	return res
}

func (this *SnapshotArray) Set(index int, val int) {
	length := len(this.infoArr[index])
	if length > 0 && this.infoArr[index][length-1].snapID == this.nowSnap {
		this.infoArr[index][length-1].value = val
	} else {
		this.infoArr[index] = append(this.infoArr[index], snapshotInfo{this.nowSnap, val})
	}
}

func (this *SnapshotArray) Snap() int {
	this.nowSnap++
	return this.nowSnap
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	info := this.infoArr[index]
	infoIndex := findValue(info, snap_id)
	return info[infoIndex].value
}

func findValue(arr []snapshotInfo, target int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2

		if arr[mid].snapID < target {
			l = mid + 1
		} else if arr[mid].snapID >= target {
			r = mid - 1
		}
	}
	return r
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
