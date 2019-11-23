package main

import . "sort"

type LessFunc func(c1, c2 *cacheValueObj) bool

// MultiSorter implements the Sort interface, sorting the changes within.
type MultiSorter struct {
	caches []*cacheValueObj
	less   []LessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *MultiSorter) Sort(caches []*cacheValueObj) {
	ms.caches = caches
	Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...LessFunc) *MultiSorter {
	return &MultiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *MultiSorter) Len() int {
	return len(ms.caches)
}
func (ms *MultiSorter) Less(i, j int) bool {
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
func (ms *MultiSorter) Swap(i, j int) {
	ms.caches[i], ms.caches[j] = ms.caches[j], ms.caches[i]
}
