package main

import (
	"math/bits"
	"strconv"
	"strings"
	"unicode"
)

var decoder = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

//https://www.codewars.com/kata/52b757663a95b11b3d00062d
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

//https://www.codewars.com/kata/52c31f8e6605bcc646000082
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

//https://www.codewars.com/kata/5ce399e0047a45001c853c2b
func PartsSums(ls []uint64) []uint64 {
	// your code here
	res := make([]uint64, len(ls)+1)
	res[len(res)-1] = 0
	for i := len(ls); i >= 1; i-- {
		res[i-1] = ls[i-1] + res[i]
	}
	return res
}

//https://www.codewars.com/kata/556deca17c58da83c00002db
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

//https://www.codewars.com/kata/514b92a657cdc65150000006
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

//https://www.codewars.com/kata/526571aae218b8ee490006f4
func CountBits(x uint) int {
	//res := bits.OnesCount(x)
	return bits.OnesCount(x)
}

//https://www.codewars.com/kata/54da5a58ea159efa38000836
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

//https://www.codewars.com/kata/541c8630095125aba6000c00
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

//https://www.codewars.com/kata/523f5d21c841566fde000009
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

//https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1
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
