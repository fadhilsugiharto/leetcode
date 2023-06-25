package main

import "fmt"

type SnapshotArray struct {
	snapshot map[int][]int
}

func Constructor(length int) SnapshotArray {
	snapshotArray := SnapshotArray{
		snapshot: map[int][]int{
			0: make([]int, length),
		},
	}
	return snapshotArray
}

func (this *SnapshotArray) Set(index int, val int) {
	recentArray := this.snapshot[len(this.snapshot)-1]
	if index < len(recentArray) {
		recentArray[index] = val
	}
}

func (this *SnapshotArray) Snap() int {
	this.snapshot[len(this.snapshot)] = make([]int, len(this.snapshot[len(this.snapshot)-1]))
	copy(this.snapshot[len(this.snapshot)], this.snapshot[len(this.snapshot)-1])

	return len(this.snapshot) - 2
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	return this.snapshot[snap_id][index]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

func main() {
	arr := Constructor(3)
	arr.Set(0, 5)
	arr.Snap()
	fmt.Println(arr)
	arr.Set(0, 6)
	fmt.Println(arr)
	fmt.Println(arr.Get(0, 0))
}
