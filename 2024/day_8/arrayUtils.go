package main

import "strconv"

func toIntArray(arr []string) []int {
	var res []int

	for _, val := range arr {
		num := toInt(val)
		res = append(res, num)

	}

	return res
}

func toInt(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}

func toString(n int) string {
	return strconv.Itoa(n)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
