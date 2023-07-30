package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// https://www.codewars.com/kata/559590633066759614000063
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

// https://www.codewars.com/kata/54ff3102c1bad923760001f3
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

// https://www.codewars.com/kata/5667e8f4e3f572a8f2000039
func Accum(s string) string {
	// your code
	var res []string
	for i := 0; i < len(s); i++ {
		res = append(res, strings.ToUpper(string(s[i]))+strings.ToLower(strings.Repeat(string(s[i]), i)))
	}
	return strings.Join(res, "-")
}

// https://www.codewars.com/users/Enligth1337/completed_solutions
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

// https://www.codewars.com/kata/56541980fa08ab47a0000040
func PrinterError(s string) string {
	var cnt int
	for _, b := range s {
		if unicode.IsLetter(b) && (b >= 'n' && b <= 'z') {
			cnt++
		}
	}
	return fmt.Sprintf("%d/%d", cnt, len(s))
}

// https://www.codewars.com/kata/52fba66badcd10859f00097e
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

// https://www.codewars.com/kata/59cfc09a86a6fdf6df0000f1
func Capitalize(st string, arr []int) string {
	s := []rune(st)
	for _, b := range arr {
		if b >= 0 && b <= len(st) {
			s[b] = unicode.ToUpper(s[b])
		}
	}
	return string(s)
}

// https://www.codewars.com/kata/56747fd5cb988479af000028
func GetMiddle(s string) string {
	//if len(s)%2 == 0 {
	//	return s[len(s)/2-1 : len(s)/2+1]
	//}
	//return s[int(len(s)/2):int(len(s)/2+1)]
	n := len(s)
	if n == 0 {
		return s
	}
	return s[(n-1)/2 : n/2+1]
}

// https://www.codewars.com/kata/56606694ec01347ce800001b
func IsTriangle(a, b, c int) bool {
	m := []int{a, b, c}
	sort.Ints(m)
	return m[2] < m[0]+m[1]
}

// https://www.codewars.com/kata/526c7363236867513f0005ca
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

// https://www.codewars.com/kata/58941fec8afa3618c9000184
func GrowingPlant(upSpeed, downSpeed, desiredHeight int) (res int) {
	curr := upSpeed
	for curr < desiredHeight {
		res++
		curr += upSpeed - downSpeed
	}
	return res + 1
}

// https://www.codewars.com/kata/554b4ac871d6813a03000035
func HighAndLow(in string) string {
	// Code here or
	s := strings.Split(in, " ")
	arr := []int{}
	for idx := range s {
		curr, _ := strconv.Atoi(s[idx])
		arr = append(arr, curr)
	}
	sort.Ints(arr)
	return fmt.Sprintf("%d %d", arr[len(arr)-1], arr[0])
}

// https://www.codewars.com/kata/571c1e847beb0a8f8900153d
func RakeGarden(garden string) string {
	s := strings.Split(garden, " ")
	for i := range s {
		if s[i] == "rock" || s[i] == "gravel" {
			continue
		}
		s[i] = "gravel"
	}
	return strings.Join(s, " ")
}

// https://www.codewars.com/kata/5259b20d6021e9e14c0010d4
func reverse(s string) (res string) {
	for _, v := range s {
		res = string(v) + res
	}
	return
}

func ReverseWords(str string) string {
	s := strings.Split(str, " ")
	for i, b := range s {
		s[i] = reverse(b)
	}
	return strings.Join(s, " ")
}
