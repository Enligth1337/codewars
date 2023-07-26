package main

import (
	"fmt"
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

//https://www.codewars.com/kata/588417e576933b0ec9000045
func seatsInTheater(nCols int, nRows int, col int, row int) int {
	return (nCols - col + 1) * (nRows - row)
}

//https://www.codewars.com/kata/57e76bc428d6fbc2d500036d
func StringToArray(str string) (res []string) {
	res = strings.Split(str, " ")
	return
}

//https://www.codewars.com/kata/582cb0224e56e068d800003c
func Litres(time float64) int {
	return int(time * 0.5)
}

//https://www.codewars.com/kata/5963c18ecb97be020b0000a2
func Derive(a, b int) string {
	return fmt.Sprintf("%dx^%d", a*b, b-1)
}

//https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0
func RemoveChar(word string) string {
	return word[1:(len(word) - 1)]
}

//https://www.codewars.com/kata/57eae20f5500ad98e50002c5
func NoSpace(word string) string {
	var res []string
	res = strings.Split(word, " ")
	str := strings.Join(res, "")
	return str
}

//https://www.codewars.com/kata/54edbc7200b811e956000556
func CountSheeps(numbers []bool) (res int) {
	for _, b := range numbers {
		if b {
			res++
		}
	}
	return res
}

//https://www.codewars.com/kata/5545f109004975ea66000086
func IsDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

//https://www.codewars.com/kata/55a70521798b14d4750000a4
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

//https://www.codewars.com/kata/5715eaedb436cf560600038
func PositiveSum(numbers []int) (res int) {
	for _, b := range numbers {
		if b > 0 {
			res += b
		}
	}
	return res
}

//https://www.codewars.com/kata/5265326f5fda8eb1160004c8
func NumberToString(n int) string {
	return fmt.Sprintf("%d", n)
}

//https://www.codewars.com/kata/56dec885c54a926dcd001095
func Opposite(value int) int {
	return -value
}

//https://www.codewars.com/kata/57a0e5c372292dd76d000d7e
func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

//https://www.codewars.com/kata/515e271a311df0350d00000f
func SquareSum(numbers []int) (res int) {
	for _, b := range numbers {
		res += b * b
	}
	return res
}

//https://www.codewars.com/kata/55a70521798b14d4750000a4
func greet(p string) string {
	if p == "Johnny" {
		return "Hello, my love!"
	}
	return "Hello, " + p + "!"
}

//https://www.codewars.com/kata/5b853229cfde412a470000d0
func TwiceAsOld(dadYearsOld, sonYearsOld int) (a int) {
	a = 2 * (dadYearsOld - sonYearsOld)
	if a > dadYearsOld {
		return a - dadYearsOld
	}
	return dadYearsOld - a
}

//https://www.codewars.com/kata/5168bb5dfe9a00b126000018
func Solutionn(word string) (res string) {
	for i := len(word) - 1; i >= 0; i-- {
		res += string(word[i])
	}
	return res
}

//https://www.codewars.com/kata/568dcc3c7f12767a62000038
func FakeBin(x string) (res string) {
	// your code here
	for i, _ := range x {
		if x[i] >= '5' {
			res += "1"
		} else {
			res += "0"
		}
	}
	return res
}

func BonusTime(salary int, bonus bool) string {
	// Your code here
	if bonus {
		return fmt.Sprintf("£%d", salary*10)
	}
	return fmt.Sprintf("£%d", salary)
}
