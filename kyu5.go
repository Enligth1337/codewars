package main

import (
	"sort"
	"unicode"
)

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

//https://www.codewars.com/kata/52223df9e8f98c7aa7000062
func findRune(curr, idx int) rune {
	target := curr + idx
	if unicode.IsLower(rune(curr)) {
		if target > 'z' {
			return rune(target - 'z' + 'a' - 1)
		} else {
			return rune(target)
		}
	} else {
		if target > 'Z' {
			return rune(target - 'Z' + 'A' - 1)
		} else {
			return rune(target)
		}
	}
}

func Rot13(msg string) (res string) {
	// Your code here
	s := []rune(msg)
	for a, b := range s {
		if s[a] >= 'a' && s[a] <= 'z' || s[a] >= 'A' && s[a] <= 'Z' {
			res += string(findRune(int(b), 13))
		} else {
			res += string(b)
		}
	}
	return res
}

//https://www.codewars.com/kata/52f787eb172a8b4ae1000a34
func Zeros(n int) (res int) {
	res = 0
	for n >= 5 {
		n /= 5
		res += n
	}
	return res
}
