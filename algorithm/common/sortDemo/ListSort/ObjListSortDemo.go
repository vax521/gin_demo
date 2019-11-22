package main

import (
	"fmt"
	"sort"
	"time"
)

type cacheValueObj struct {
	value    int
	hitCount int
	lastTime time.Time
}

func (c cacheValueObj) String() string {
	return fmt.Sprintf("(%v, %v,%v)", c.value, c.hitCount, c.lastTime.Format("2006/01/02 15:04:05"))
}

func NewCacheObj(value int, hitcount int, lastTime time.Time) *cacheValueObj {
	return &cacheValueObj{value: value, hitCount: hitcount, lastTime: lastTime}
}

type lessFunc func(c1, c2 *cacheValueObj) bool

// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
	caches []*cacheValueObj
	less   []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(caches []*cacheValueObj) {
	ms.caches = caches
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.caches)
}
func (ms *multiSorter) Less(i, j int) bool {
	p, q := ms.caches[i], ms.caches[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.caches[i], ms.caches[j] = ms.caches[j], ms.caches[i]
}

func main() {
	var testCacheList []*cacheValueObj
	for i := 0; i < 10; i++ {
		tempObj := NewCacheObj(i, 0, time.Now())
		fmt.Println(tempObj)
		time.Sleep(time.Second)
		testCacheList = append(testCacheList, tempObj)
	}
	fmt.Println(testCacheList)
	sort.Slice(testCacheList, func(i, j int) bool {
		return testCacheList[i].value > testCacheList[j].value
	})
	fmt.Println(testCacheList)
	sort.Slice(testCacheList, func(i, j int) bool {
		return testCacheList[i].lastTime.Unix() < testCacheList[j].lastTime.Unix()
	})
	fmt.Println(testCacheList)

	testCacheList[0].hitCount = 10
	testCacheList[1].hitCount = 8
	testCacheList[2].hitCount = 8
	testCacheList[3].hitCount = 6
	testCacheList[4].hitCount = 6
	testCacheList[5].hitCount = 6
	testCacheList[7].hitCount = 6
	sort.Slice(testCacheList, func(i, j int) bool {
		return testCacheList[i].hitCount < testCacheList[j].hitCount
	})
	fmt.Println(testCacheList)

	increHitCountSort := func(c1, c2 *cacheValueObj) bool {
		return c1.hitCount < c2.hitCount
	}
	decreLastTimeSort := func(c1, c2 *cacheValueObj) bool {
		return c1.lastTime.Unix() > c2.lastTime.Unix()
	}
	OrderedBy(increHitCountSort, decreLastTimeSort).Sort(testCacheList)
	fmt.Println(testCacheList)

}
