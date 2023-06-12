package main

import "sort"

//https://www.codewars.com/kata/559a28007caad2ac4e000083
func Perimeter(n int) int {
	// your code
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1 * 4
	}
	if n == 2 {
		return 2 * 4
	}

	res := 0
	fibb := map[int]int{1: 1, 2: 1}

	for i := 2; i < n+1; i++ {
		fibb[i+1] = fibb[i-1] + fibb[i]
	}
	for _, edge := range fibb {
		res += edge * 4
	}
	return res
}

//https://www.codewars.com/kata/525c65e51bf619685c000059
func Cakes(recipe, available map[string]int) int {
	var cnt []int
	for i, _ := range recipe {
		if j, ok := available[i]; ok {
			j = j / recipe[i]
			cnt = append(cnt, j)
		} else {
			return 0
		}

	}
	sort.Ints(cnt)
	return cnt[0]
}
