package main

import (
	"fmt"
	"math/bits"
	"sort"
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

func main() {
	//fmt.Println(Decode("IX"))
	//fmt.Println(Decode("MMVIII"))
	//fmt.Println(Decode("MDCLXVI"))
	//fmt.Println(IsUpperCase("aFDSvvszxsdgersDSDFSFsvbgdefre"))
	//fmt.Println(IsUpperCase("JHIJDNHFSJ"))
	//fmt.Println(BinToDec("10001"))
	//fmt.Println(CalculateYears(2))
	//fmt.Println(TwoSum([]int{1, 3, 2}, 4))
	//fmt.Println(toWeirdCase("This is a test Looks like you passed"))
	//fmt.Println(Perimeter(5))
	//fmt.Println(PartsSums([]uint64{0, 1, 3, 6, 10}))
	//fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
	//fmt.Println(Zeros(30))
	//fmt.Println(Tribonacci([3]float64{1, 1, 1}, 1))
	//fmt.Println(Between(1, 5))
	//fmt.Println(EvenOrOdd(1))
	//fmt.Println(GetCount("abracadabra"))
	//fmt.Println(Contamination("abc", "z"))
	//fmt.Println(multipleOfIndex([]int{22, -6, 32, 82, 9, 25}))
	//fmt.Println(Accum("ZpglnRxqenU"))
	//fmt.Println(Multiple3And5(10))
	//fmt.Println(Disemvowel("This website is for losers LOL!"))
	//fmt.Println(CountBits(1234))
	fmt.Println(FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
}

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

func CountBits(x uint) int {
	res := bits.OnesCount(x)
	return res
}

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

func Accum(s string) string {
	// your code
	var res []string
	for i := 0; i < len(s); i++ {
		res = append(res, strings.ToUpper(string(s[i]))+strings.ToLower(strings.Repeat(string(s[i]), i)))
	}
	return strings.Join(res, "-")
}

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

func Contamination(text string, char string) string {
	var res string
	for _, _ = range text {
		res += char
	}
	return res
}

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

func EvenOrOdd(number int) string {
	// your code here
	if number == 0 || number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func Between(a, b int) []int {
	// your code here
	res := []int{}
	for i := 0; i < b; i++ {
		res = append(res, a+i)
	}
	return res
}

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

func Zeros(n int) int {
	// your code here
	count := 0
	num := 0
	for i := 1; i <= n; i++ {
		num = i
		for num%5 == 0 {
			count++
			num /= 5
		}
	}
	return count
}

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

func PartsSums(ls []uint64) []uint64 {
	// your code here
	res := make([]uint64, len(ls)+1)
	res[len(res)-1] = 0
	for i := len(ls); i >= 1; i-- {
		res[i-1] = ls[i-1] + res[i]
	}
	return res
}

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

func BinToDec(bin string) int {
	// your code here
	res, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return 0
	}
	return int(res)
}

type MyString string

func IsUpperCase(s MyString) bool {
	// Your code here!
	//var a string
	//for i, _ := range s {
	//	a = strings.ToUpper(string(s[i]))
	//	if a == string(s[i]) {
	//		continue
	//	} else {
	//		return false
	//	}
	//}
	//return true

	return s == MyString(strings.ToUpper(string(s)))
}

func GetSize(w, h, d int) [2]int {
	// your code here
	s := 2 * (w*h + w*d + h*d)
	v := w * h * d
	return [2]int{s, v}
}

func Decode(roman string) int {
	if len(roman) == 0 {
		return 0
	}
	first := decoder[rune(roman[0])]
	if len(roman) == 1 {
		return first
	}
	next := decoder[rune(roman[1])]
	if next > first {
		return (next - first) + Decode(roman[2:])
	}
	return first + Decode(roman[1:])
}
