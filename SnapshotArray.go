package main

import "fmt"

// NGGAE DEWE
//type SnapshotArray struct {
//	snapshot map[int][]int
//}
//
//func Constructor(length int) SnapshotArray {
//	snapshotArray := SnapshotArray{
//		snapshot: map[int][]int{
//			0: make([]int, length),
//		},
//	}
//	return snapshotArray
//}
//
//func (this *SnapshotArray) Set(index int, val int) {
//	recentArray := this.snapshot[len(this.snapshot)-1]
//	if index < len(recentArray) {
//		recentArray[index] = val
//	}
//}
//
//func (this *SnapshotArray) Snap() int {
//	this.snapshot[len(this.snapshot)] = make([]int, len(this.snapshot[len(this.snapshot)-1]))
//	copy(this.snapshot[len(this.snapshot)], this.snapshot[len(this.snapshot)-1])
//
//	return len(this.snapshot) - 2
//}
//
//func (this *SnapshotArray) Get(index int, snap_id int) int {
//	return this.snapshot[snap_id][index]
//}

// 5 Head Approach
type Snapshot struct {
	id  int
	val int
}

type SnapshotsHistory struct {
	snapshots *[]Snapshot
	currVal   int
}

type SnapshotArray struct {
	array      *[]*SnapshotsHistory
	snapshotId int
}

func Constructor(length int) SnapshotArray {
	arr := make([]*SnapshotsHistory, length)
	res := SnapshotArray{&arr, 0}

	for i := 0; i < length; i++ {
		hist := make([]Snapshot, 0)
		(*res.array)[i] = &SnapshotsHistory{&hist, 0}
	}

	return res
}

func (this *SnapshotArray) Set(index int, val int) {
	hist := (*this.array)[index]

	if this.snapshotId > 0 && (len(*hist.snapshots) == 0 || (*hist.snapshots)[len(*hist.snapshots)-1].id < this.snapshotId) {
		*hist.snapshots = append(*hist.snapshots, Snapshot{this.snapshotId - 1, hist.currVal})
	}

	hist.currVal = val
}

func (this *SnapshotArray) Snap() int {
	defer func() {
		this.snapshotId++
	}()
	return this.snapshotId
}

func (this *SnapshotArray) Get(index int, snapId int) int {
	hist := (*this.array)[index]

	if len(*hist.snapshots) == 0 || (*hist.snapshots)[len(*hist.snapshots)-1].id < snapId {
		return hist.currVal
	}

	l := 0
	r := len(*hist.snapshots)
	m := 0

	for l < r {
		m = (l + r) / 2

		mId := (*hist.snapshots)[m].id

		if m == l && mId >= snapId || m > l && mId >= snapId && (*hist.snapshots)[m-1].id < snapId {
			break
		}

		if mId < snapId {
			l = m + 1
		} else {
			r = m
		}
	}

	return (*hist.snapshots)[m].val
}

func main() {
	arr := Constructor(3)
	arr.Set(0, 5)
	arr.Snap()
	fmt.Println(arr)
	arr.Set(0, 6)
	fmt.Println(arr)
	fmt.Println(arr.Get(0, 0))
}
