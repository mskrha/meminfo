package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	FILE = "/proc/meminfo"
)

type Info struct {
	Total      uint64
	Free       uint64
	Buffers    uint64
	Cached     uint64
	PageTables uint64
}

func main() {
	data, err := get()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(calc(data))
}

func get() (ret Info, err error) {
	f, err := ioutil.ReadFile(FILE)
	if err != nil {
		return
	}

	s := strings.Replace(string(f), ":", "", -1)

	for _, v := range strings.Split(s, "\n") {
		x := strings.Fields(v)
		if len(x) < 2 {
			continue
		}
		switch x[0] {
		case "MemTotal":
			ret.Total, err = strconv.ParseUint(x[1], 10, 64)
			if err != nil {
				return
			}
		case "MemFree":
			ret.Free, err = strconv.ParseUint(x[1], 10, 64)
			if err != nil {
				return
			}
		case "Buffers":
			ret.Buffers, err = strconv.ParseUint(x[1], 10, 64)
			if err != nil {
				return
			}
		case "Cached":
			ret.Cached, err = strconv.ParseUint(x[1], 10, 64)
			if err != nil {
				return
			}
		case "PageTables":
			ret.PageTables, err = strconv.ParseUint(x[1], 10, 64)
			if err != nil {
				return
			}
		}
	}

	return
}

func calc(d Info) string {
	used := d.Total - d.Free - d.Buffers - d.Cached + d.PageTables
	return fmt.Sprintf("%.2f GB (%.1f%%) / %d MB", float64(used)/1048576, 100*float64(used)/float64(d.Total), d.Free/1024)
}
