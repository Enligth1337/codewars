package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

//https://www.codewars.com/kata/559590633066759614000063
func MinMax(arr []int) [2]int {
	//mm := map[string]int{"min": arr[0], "max": arr[0]}
	//for _, a := range arr {
	//	if a < mm["min"] {
	//		mm["min"] = a
	//	}
	//	if a > mm["max"] {
	//		mm["max"] = a
	//	}
	//}
	//return [2]int{mm["min"], mm["max"]}
	sort.Ints(arr)
	return [2]int{arr[0], arr[len(arr)-1]}
}

//https://www.codewars.com/kata/54ff3102c1bad923760001f3
func GetCount(str string) (count int) {
	// Enter solution here
	var res int
	vowels := "aeiouAEIOU"
	for i := 0; i < len(str); i++ {
		if strings.ContainsRune(vowels, rune(str[i])) {
			res++
		}
	}
	return res
}

//https://www.codewars.com/kata/5667e8f4e3f572a8f2000039
func Accum(s string) string {
	// your code
	var res []string
	for i := 0; i < len(s); i++ {
		res = append(res, strings.ToUpper(string(s[i]))+strings.ToLower(strings.Repeat(string(s[i]), i)))
	}
	return strings.Join(res, "-")
}

//https://www.codewars.com/users/Enligth1337/completed_solutions
func NbYear(p0 int, percent float64, aug int, p int) int {
	// your code
	curr := p0
	res := 0
	for curr < p {
		curr = curr + aug + (int(float64(curr) * percent / 100))
		res++
	}
	return res
}

//https://www.codewars.com/kata/56541980fa08ab47a0000040
func PrinterError(s string) string {
	var cnt int
	for _, b := range s {
		if unicode.IsLetter(b) && (b >= 'n' && b <= 'z') {
			cnt++
		}
	}
	return fmt.Sprintf("%d/%d", cnt, len(s))
}

//https://www.codewars.com/kata/52fba66badcd10859f00097e
func Disemvowel(comment string) string {
	var res string
	vowels := "aeiouAEIOU"
	for i := 0; i < len(comment); i++ {
		if strings.ContainsRune(vowels, rune(comment[i])) {
			continue
		}
		res += string(comment[i])
	}
	return res
}

//https://www.codewars.com/kata/59cfc09a86a6fdf6df0000f1
func Capitalize(st string, arr []int) string {
	s := []rune(st)
	for _, b := range arr {
		if b >= 0 && b <= len(st) {
			s[b] = unicode.ToUpper(s[b])
		}
	}
	return string(s)
}
