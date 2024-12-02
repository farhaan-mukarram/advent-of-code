package main

import (
	"strconv"
)

func toIntArray(arr []string) []int {
	var res []int

	for _, val := range arr {
		num, _ := strconv.Atoi(val)
		res = append(res, num)

	}

	return res
}

// TODO: Figure out how this works, copied from SO
func remove(arr []int, i int) []int {
	var temp []int
	temp = append(temp, arr[:i]...)
	temp = append(temp, arr[i+1:]...)

	return temp
}
