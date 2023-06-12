package main

import (
	"fmt"
)

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
	//fmt.Println(FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))

	//fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))

	//fmt.Println(DigitalRoot(16))
	fmt.Println(SmallestIntegerFinder([]int{1, 23, 481}))

}

func Zeros(n int) int {
	// your code here
	var mp map[int]int
	count := 0
	num := 0
	for i := 1; i <= n; i++ {
		if _, ok := mp[i]; ok {
			mp[i]++
		}

		num = i
		for num%5 == 0 {
			count++
			num /= 5
		}
	}
	return count

}

//func HighAndLow(in string) string {
//	// Code here or
//	if len(in) == 0 {
//		return "0 0"
//	}
//	mp := map[string]int{}
//
//	for i, _ := range in {
//		target, _ := strconv.Atoi(string(in[i]))
//	}
//
//	return "throw towel"
//}
