package utils

import (
	"fmt"
	"os"
	"runtime"
	"sort"
)

type MapSorter []Item

var filemap = make(map[string]int64)

type Item struct {
	Key string
	Val int64
}

func NewMapSorter(m map[string]int64) MapSorter {
	ms := make(MapSorter, 0, len(m))
	for k, v := range m {
		ms = append(ms, Item{k, v})
	}
	return ms
}

func (ms MapSorter) Len() int {
	return len(ms)
}

func (ms MapSorter) Less(i, j int) bool {
	return ms[i].Val > ms[j].Val // sort by value, high to low
	//return ms[i].Key > ms[j].Key // sort by key
}

func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func FolderSort(src string) {
	/*
		if len(os.Args) < 2 {
			fmt.Println("need path parameters ")
			os.Exit(1)
		}
		src := strings.TrimRight(os.Args[1], "/")
	*/
	//src := "c:/wuhao"
	runtime.GOMAXPROCS(runtime.NumCPU())
	GetAllFile(src, len(src)+1)
	ms := NewMapSorter(filemap)
	sort.Sort(ms)

	for _, item := range ms {
		fmt.Printf("%s:%d\n", item.Key, item.Val)
	}
}

func GetAllFile(pathname string, strip int) error {
	rd, err := os.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			GetAllFile(pathname+"/"+fi.Name(), strip)
		} else {
			fn := pathname + "/" + fi.Name()
			s, _ := fi.Info()

			filemap[fn[strip:]] = s.Size()
			//fmt.Println(fn[strip:])
		}
	}
	return err
}
