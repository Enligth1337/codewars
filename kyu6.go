package main

import (
	"math"
	"math/bits"
	"strconv"
	"strings"
	"unicode"
)

// https://www.codewars.com/kata/52b757663a95b11b3d00062d
func toWeirdCase(str string) string {
	// Your code here and happy coding!
	var res string
	for i, _ := range str {
		if i == 0 || string(rune(res[i-1])) == " " || unicode.IsLower(rune(res[i-1])) {
			res += strings.ToUpper(string(str[i]))
		} else {
			res += strings.ToLower(string(str[i]))
		}
	}
	return res
}

// https://www.codewars.com/kata/52c31f8e6605bcc646000082
func TwoSum(numbers []int, target int) [2]int {
	a := map[int]int{}

	for i, b := range numbers {
		p := target - b

		if _, ok := a[p]; ok {
			return [2]int{a[p], i}
		}
		a[b] = i

	}
	return [2]int{}
}

//func Decode(roman string) int {
//	if len(roman) == 0 {
//		return 0
//	}
//	first := decoder[rune(roman[0])]
//	if len(roman) == 1 {
//		return first
//	}
//	next := decoder[rune(roman[1])]
//	if next > first {
//		return (next - first) + Decode(roman[2:])
//	}
//	return first + Decode(roman[1:])
//}

// https://www.codewars.com/kata/5ce399e0047a45001c853c2b
func PartsSums(ls []uint64) []uint64 {
	// your code here
	res := make([]uint64, len(ls)+1)
	res[len(res)-1] = 0
	for i := len(ls); i >= 1; i-- {
		res[i-1] = ls[i-1] + res[i]
	}
	return res
}

// https://www.codewars.com/kata/556deca17c58da83c00002db
func Tribonacci(signature [3]float64, n int) []float64 {
	// your code here
	if n < 3 {
		return signature[:n]
	}
	res := []float64{signature[0], signature[1], signature[2]}
	var a float64
	for i := 3; i < n; i++ {
		a = res[i-1] + res[i-2] + res[i-3]
		res = append(res, a)
	}
	return res
}

// https://www.codewars.com/kata/514b92a657cdc65150000006
func Multiple3And5(number int) int {
	if number < 0 {
		return 0
	}
	var res int
	for i := number - 1; i >= 0; i-- {
		if i%3 == 0 || i%5 == 0 {
			res = i + res
		}
	}
	return res
}

// https://www.codewars.com/kata/526571aae218b8ee490006f4
func CountBits(x uint) int {
	//res := bits.OnesCount(x)
	return bits.OnesCount(x)
}

// https://www.codewars.com/kata/54da5a58ea159efa38000836
func FindOdd(seq []int) int {
	// your code here
	res := make(map[int]int)
	for _, b := range seq {
		target := b
		if _, ok := res[target]; ok {
			res[target]++
		} else {
			res[b] = 1
		}
	}
	for i, b := range res {
		if b%2 != 0 {
			return i
		}
	}
	return 0
}

// https://www.codewars.com/kata/541c8630095125aba6000c00
func DigitalRoot(n int) int {
	// ...
	if n < 10 {
		return n
	}
	var res int
	arr := strconv.Itoa(n)
	for i, _ := range arr {
		a, _ := strconv.Atoi(string(arr[i]))
		res = a + res
	}
	if res < 10 {
		return res
	}
	return DigitalRoot(res)
}

// https://www.codewars.com/kata/523f5d21c841566fde000009
func ArrayDiff(a, b []int) (res []int) {
	// your code here
	if len(b) == 0 || len(a) == 0 {
		return a
	}
	mp := make(map[int]int, len(b))
	for k, v := range b {
		mp[v] = k
	}
	for _, v := range a {
		if _, ok := mp[v]; ok {
			continue
		} else {
			res = append(res, v)
		}
	}
	return res
}

// https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1
func duplicate_count(s1 string) (res int) {
	//your code goes here
	mp := make(map[rune]int, len(s1))
	for _, b := range strings.ToLower(s1) {
		if mp[b] == 1 {
			res++
		}
		mp[b]++
	}
	return res
}

func duplicate_countcel(s string) (c int) {
	h := map[rune]int{}
	for _, r := range strings.ToLower(s) {
		if h[r]++; h[r] == 2 {
			c++
		}
	}
	return
}

// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001
func Solution(str string) (res []string) {
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}
	return res
}

// https://www.codewars.com/kata/5526fc09a1bbd946250002dc
func rev(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func SpinWords(str string) (res string) {
	arr := strings.Split(str, " ")
	for a, b := range arr {
		if len(b) >= 5 {
			arr[a] = rev(b)
		}
	}
	res = strings.Join(arr, " ")

	return res
}

// https://www.codewars.com/kata/51b62bf6a9c58071c600001b
var romanDict = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func intToRoman(num int) (res string) {
	sortArr := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, b := range sortArr {
		for num >= b {
			res += romanDict[b]
			num -= b
		}
	}
	return res
}

// https://www.codewars.com/kata/52fb87703c1351ebd200081f
func WhatCentury(year string) (res string) {
	y, _ := strconv.Atoi(year)
	s := int(math.Ceil(float64(y) / 100))
	res = strconv.Itoa(s)
	l := len(res) - 1
	if res[l-1:] != "12" && res[l:] == "2" {
		res += "nd"
	} else if res[l-1:] != "11" && res[l:] == "1" {
		res += "st"
	} else if res[l-1:] != "13" && res[l:] == "3" {
		res += "rd"
	} else {
		res += "th"
	}
	return res
}

// https://www.codewars.com/kata/5526fc09a1bbd946250002dc
func FindOutlier(integers []int) int {
	var arr []int
	odd := 0
	for i := range integers {
		if integers[i]%2 == 0 {
			arr = append(arr, integers[i])
		} else {
			odd = integers[i]
		}
	}
	if len(arr) == 1 {
		return arr[0]
	}
	return odd
}

// https://www.codewars.com/kata/587731fda577b3d1b0001196
func CamelCase(s string) (res string) {
	if s == "" {
		return ""
	}
	arr := strings.Split(strings.Trim(s, " "), " ")
	for _, b := range arr {
		res += strings.ToUpper(string(b[0])) + b[1:]
	}
	return res
}

// https://www.codewars.com/kata/54da539698b8a2ad76000228
func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}
	m := map[rune][]int{'n': {0, 1}, 's': {0, -1}, 'w': {1, 0}, 'e': {-1, 0}}
	curr := [2]int{0, 0}
	for i := 0; i < 10; i++ {
		step := m[walk[i]]
		curr[0], curr[1] = step[0]+curr[0], step[1]+curr[1]

	}
	return curr == [2]int{0, 0}
}

// https://www.codewars.com/kata/54b42f9314d9229fd6000d9c
func DuplicateEncode(word string) (res string) {
	m := make(map[rune]int)
	word = strings.ToLower(word)
	for _, b := range word {
		m[b]++
	}
	for _, b := range word {
		if m[b] > 1 {
			res += ")"
			continue
		}
		res += "("
	}
	return res
}

// https://www.codewars.com/kata/5262119038c0985a5b00029f
func IsPrime(n int) bool {
	//your code here
	if n == 2 {
		return true
	}
	if n < 2 || n%2 == 0 {
		return false
	}
	sq := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// https://www.codewars.com/kata/55bf01e5a717a0d57e0000ec
func Persistence(n int) (res int) {
	// your code
	cur := n
	var s string
	for i := 1; cur >= 10; i++ {
		s = strconv.Itoa(n)
		n = 1
		res = i
		for j := 0; j < len(s); j++ {
			n *= int(s[j] - '0')
			cur = n
		}
	}
	return res
}

// https://www.codewars.com/kata/55c45be3b2079eccff00010f
func Order(sentence string) string {
	by := strings.Split(sentence, " ")
	l := len(by)
	res := make([]string, l)
	for _, b := range by {
		for j, _ := range b {
			if b[j] > '0' && b[j] <= '9' {
				res[int(b[j]-'1')] = b
				break
			}
		}
	}
	return strings.Join(res, " ")
}

// https://www.codewars.com/kata/585d7d5adb20cf33cb000235
func FindUniq(arr []float32) float32 {
	// Do the magic
	m := make(map[float32]int)
	for i := range arr {
		m[arr[i]]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

// https://www.codewars.com/kata/5839edaa6754d6fec10000a2
func FindMissingLetter(chars []rune) rune {
	if len(chars) < 1 {
		return chars[0]
	}
	cur := chars[0]
	for i := range chars {
		if cur != chars[i] {
			return cur
		}
		cur += 1
	}
	return chars[0]
}

// https://www.codewars.com/kata/5a58ca28e626c55ae000018a
func AreaOfPolygonInsideCircle(circleRadius float64, numberOfSides int) float64 {
	// your code here
	return math.Round(((float64(numberOfSides)*math.Pow(2*circleRadius*math.Sin(math.Pi/float64(numberOfSides)), 2))/(4*math.Tan(math.Pi/float64(numberOfSides))))*1000) / 1000
}
