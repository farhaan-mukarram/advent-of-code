package main

import "strconv"

func toIntArray(arr []string) []int {
	var res []int

	for _, val := range arr {
		num, _ := strconv.Atoi(val)
		res = append(res, num)

	}

	return res
}
