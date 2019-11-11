package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/slice"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	a := slice.IntSlice(make([]int, 0))
	a = append(a, 1)
	a = append(a, 4)
	a = append(a, 3)
	a = append(a, 6)
	fmt.Printf("%v\n", a)

	sort.Sort(a.Begin(), a.End())
	fmt.Printf("%v\n", a)

	sort.Sort(a.Begin(), a.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
	fmt.Printf("%v\n", a)
}
