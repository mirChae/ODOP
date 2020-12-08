package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	var raw = map[int]map[int]int{}

	var r sort.IntSlice
	r = append(r, nums...)
	r.Sort()

	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			for k := len(r) - 1; k > j; k-- {
				if r[i]+r[j]+r[k] == 0 {
					if _, ok := raw[r[i]]; !ok {
						raw[r[i]] = map[int]int{}
						raw[r[i]][r[j]] = r[k]
					} else {
						if _, ok := raw[r[i]][r[j]]; !ok {
							raw[r[i]][r[j]] = r[k]
						}
					}
				}
			}
		}
	}

	for x, v := range raw {
		for y, z := range v {
			result = append(result, []int{x, y, z})
		}
	}

	return result
}

func main() {
	//	var nums = []int{-1, 0, 1, 2, -1, -4}
	var nums = make([]int, 500)
	fmt.Println(threeSum(nums))
}
