package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//fmt.Println(Decode("IX"))
	//fmt.Println(Decode("MMVIII"))
	//fmt.Println(Decode("MDCLXVI"))
	//https://www.codewars.com/kata/56cd44e1aa4ac7879200010b
	//fmt.Println(IsUpperCase("aFDSvvszxsdgersDSDFSFsvbgdefre"))
	//fmt.Println(IsUpperCase("JHIJDNHFSJ"))
	//https://www.codewars.com/kata/57a5c31ce298a7e6b7000334
	//fmt.Println(BinToDec("10001"))
	//https://www.codewars.com/kata/5a6663e9fd56cb5ab800008b
	//fmt.Println(CalculateYears(2))
	//https://www.codewars.com/kata/52c31f8e6605bcc646000082
	//fmt.Println(TwoSum([]int{1, 3, 2}, 4))
	//https://www.codewars.com/kata/52b757663a95b11b3d00062d
	//fmt.Println(toWeirdCase("This is a test Looks like you passed"))
	//https://www.codewars.com/kata/559a28007caad2ac4e000083
	//fmt.Println(Perimeter(5))
	//https://www.codewars.com/kata/5ce399e0047a45001c853c2b
	//fmt.Println(PartsSums([]uint64{0, 1, 3, 6, 10}))
	//https://www.codewars.com/kata/559590633066759614000063
	//fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
	//https://www.codewars.com/kata/556deca17c58da83c00002db
	//fmt.Println(Tribonacci([3]float64{1, 1, 1}, 1))
	//https://www.codewars.com/kata/55ecd718f46fba02e5000029
	//fmt.Println(Between(1, 5))
	//https://www.codewars.com/kata/53da3dbb4a5168369a0000fe
	//fmt.Println(EvenOrOdd(1))
	//https://www.codewars.com/users/Enligth1337/completed_solutions
	//fmt.Println(NbYear(1500, 5, 100, 5000))
	//https://www.codewars.com/kata/54ff3102c1bad923760001f3
	//fmt.Println(GetCount("abracadabra"))
	//https://www.codewars.com/kata/596fba44963025c878000039
	//fmt.Println(Contamination("abc", "z"))
	//https://www.codewars.com/kata/5a34b80155519e1a00000009
	//fmt.Println(multipleOfIndex([]int{22, -6, 32, 82, 9, 25}))
	//https://www.codewars.com/kata/5667e8f4e3f572a8f2000039
	//fmt.Println(Accum("ZpglnRxqenU"))
	//https://www.codewars.com/kata/514b92a657cdc65150000006
	//fmt.Println(Multiple3And5(10))
	//https://www.codewars.com/kata/52fba66badcd10859f00097e
	//fmt.Println(Disemvowel("This website is for losers LOL!"))
	//https://www.codewars.com/kata/526571aae218b8ee490006f4
	//fmt.Println(CountBits(1234))
	//https://www.codewars.com/kata/54da5a58ea159efa38000836
	//fmt.Println(FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
	//https://www.codewars.com/kata/565f5825379664a26b00007c
	//fmt.Println(GetSize(5, 7, 8))
	//https://www.codewars.com/kata/56541980fa08ab47a0000040
	//fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
	//https://www.codewars.com/kata/541c8630095125aba6000c00
	//fmt.Println(DigitalRoot(16))
	//https://www.codewars.com/kata/55a2d7ebe362935a210000b2
	//fmt.Println(SmallestIntegerFinder([]int{1, 23, 481}))
	//https://www.codewars.com/kata/525c65e51bf619685c000059
	//fmt.Println(Cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}))
	//fmt.Println(xor([]int{2, 1, 2, 1, 3, 5, 3}))
	//https://www.codewars.com/kata/523f5d21c841566fde000009
	//fmt.Println(ArrayDiff([]int{1, 2, 3}, []int{1, 2}))
	//https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1
	//fmt.Println(duplicate_count("indivisibility"))
	//fmt.Println(RGB(0, 275, 125))
	//fmt.Println(CountBits(1234))
	//fmt.Println(Capitalize("abcdef", []int{1, 2, 5}))
	//fmt.Println(seatsInTheater(16, 11, 5, 3))
	//fmt.Println(StringToArray("Robin Singh"))
	//fmt.Println(Derive(7, 8))
	//fmt.Println(RemoveChar("eloquent"))
	////fmt.Println(CountSheeps([]bool{
	//	true, true, true, false,
	//	true, true, true, true,
	//	true, false, true, false,
	//	true, false, false, true,
	//	true, true, true, true,
	//	false, false, true, true,
	//}))
	//fmt.Println(Greet("Oleg"))
	//fmt.Println(PositiveSum([]int{1, -2, 3, 4, 5}))
	//fmt.Println(RepeatStr(4, "a"))
	//fmt.Println(Solve([]int{7, 8, 7}))

	//fmt.Println(Solution("abc"))
	//fmt.Println(GetMiddle("abc"))

	//fmt.Println(GrowingPlant(100, 10, 910))

	//fmt.Println(Decode("MMVIII"))

	//fmt.Println(solution(" ", ""))

	//fmt.Println(NextBigger(5275))

	fmt.Println(SpinWords("Burgers are my favorite fruit"))

}

func NextBigger(n int) (res int) {
	// your code here
	s := strconv.Itoa(n)
	var arr []int
	for _, b := range s {
		curr, _ := strconv.Atoi(string(b))
		arr = append(arr, curr)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	var str string
	for _, b := range arr {
		str += strconv.Itoa(b)
	}
	res, _ = strconv.Atoi(str)
	if res == n {
		return -1
	}
	return res
}

func solution(str, ending string) bool {
	// Your code here!
	if len(str) > 1 || str == "" {
		return str[(len(str)-len(ending)):] == ending
	}
	return false
}

var mp = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func Decode(roman string) (res int) {
	if len(roman) == 1 {
		return mp[rune(roman[0])]
	}
	s := []rune(roman)
	res = mp[s[0]]
	for i := 1; i < len(s); i++ {
		if mp[s[i]] > mp[s[i-1]] {
			res += mp[s[i]] - mp[s[i-1]]
		} else {
			res += mp[s[i]]
		}
	}
	return res
}

func Solve(arr []int) (res int) {
	sort.Ints(arr)
	switch {
	case arr[2] > arr[0]+arr[1]:
		return arr[0] + arr[1]
	default:
		return arr[2]
	}
}

//func RGB(r, g, b int) string {
//	// Your code here
//	bt := []byte{}
//	r := hex.EncodeToString(bt[r])
//
//	res := hex.EncodeToString(bt)
//	return strings.ToUpper(res)
//}

func Binarray(a []int) int {
	//mp := map[string]int{"curr?": a[0], "range": 1, "max": 0}
	var res []int

	for i := 1; i < len(a); i++ {
		res = append(res, a[i-1]|a[i])
	}
	fmt.Println(res)
	return 0
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
