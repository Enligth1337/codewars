package main

import (
	"sort"
	"strconv"
	"strings"
)

//https://www.codewars.com/kata/565f5825379664a26b00007c
func GetSize(w, h, d int) [2]int {
	// your code here
	s := 2 * (w*h + w*d + h*d)
	v := w * h * d
	return [2]int{s, v}
}

//https://www.codewars.com/kata/56cd44e1aa4ac7879200010b
type MyString string

func IsUpperCase(s MyString) bool {
	return s == MyString(strings.ToUpper(string(s)))
}

//https://www.codewars.com/kata/57a5c31ce298a7e6b7000334
func BinToDec(bin string) int {
	// your code here
	res, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return 0
	}
	return int(res)
}

//https://www.codewars.com/kata/5a6663e9fd56cb5ab800008b
func CalculateYears(years int) [3]int {
	// Write your solution here
	if years == 1 {
		return [3]int{1, 15, 15}
	}
	if years == 2 {
		return [3]int{2, 24, 24}
	}
	result := [3]int{2, 24, 24}
	for i := 2; i < years; i++ {
		result[0] += +1
		result[1] += +4
		result[2] += +5
	}
	return result
}

//https://www.codewars.com/kata/55ecd718f46fba02e5000029
func Between(a, b int) []int {
	// your code here
	res := []int{}
	for i := 0; i < b; i++ {
		res = append(res, a+i)
	}
	return res
}

//https://www.codewars.com/kata/53da3dbb4a5168369a0000fe
func EvenOrOdd(number int) string {
	// your code here
	if number == 0 || number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

//https://www.codewars.com/kata/596fba44963025c878000039
func Contamination(text string, char string) string {
	var res string
	for _, _ = range text {
		res += char
	}
	return res
}

//https://www.codewars.com/kata/5a34b80155519e1a00000009
func multipleOfIndex(ints []int) []int {
	// good luck
	var res []int
	for i := 1; i < len(ints); i++ {
		if ints[i]%i == 0 {
			res = append(res, ints[i])
		}
	}
	return res
}

//https://www.codewars.com/kata/55a2d7ebe362935a210000b2
func SmallestIntegerFinder(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0] // your code here
}

//https://www.codewars.com/kata/56fa3c5ce4d45d2a52001b3c/train/go
func Xor(a, b bool) bool {
	// your code here:
	if a != b {
		return true
	} else {
		return false
	}
}
