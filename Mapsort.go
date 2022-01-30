package utils
import (
	"constraints"
	"fmt"
	"sort"
)

type Item[T comparable, V constraints.Ordered] struct {
	Key T
	Val V
}

func NewMapSorter[K comparable, V constraints.Ordered](m map[K]V) []K {
	//var ms []Item[K, V]
	ms := make([]Item[K, V], 0, len(m))
	for k, v := range m {
		var t Item[K, V]
		t.Key = k
		t.Val = v
		ms = append(ms, t)
	}
	// sort ms
	sort.Slice(ms, func(i, j int) bool {
		return ms[j+1].Val > ms[j].Val
	})
	var ret []K
	for _, v := range ms {
		ret = append(ret, v.Key)
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
