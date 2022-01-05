package utils

import (
	"constraints"
	"fmt"
	"sort"
)

type Item struct {
	Key any
	Val any
}
type MapSorter []Item

func NewMapSorter[K comparable, V constraints.Ordered](m map[K]V) []K {

	ms := make(MapSorter, 0, len(m))
	for k, v := range m {
		ms = append(ms, Item{k, v})
	}
	// sort ms
	sort.Slice(ms, func(i, j int) bool {
		return ms[j+1].Val.(V) > ms[j].Val.(V)
	})
	var ret []K
	for _, v := range ms {
		ret = append(ret, v.Key.(K))
	}
	return ret
}

func TestMapSort() {
	m := map[string]int8{
		"e": 10,
		"a": 2,
		"d": 15,
		"c": 12,
		"f": 1,
		"b": 12,
	}
	ms := NewMapSorter(m)

	for _, item := range ms {
		fmt.Printf("%s:%d\n", item, m[item])
	}
}
