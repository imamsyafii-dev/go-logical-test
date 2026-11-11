package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Question 1")
	fmt.Println(SearchTheMissingOne([]int{-1, 1}))      // 0
	fmt.Println(SearchTheMissingOne([]int{-1, -7, -3})) // -2
	fmt.Println(SearchTheMissingOne([]int{1, 2, 1, 6})) // 3

	fmt.Println("Question 2")
	fmt.Println(ExcludeJoinAndReverse([]int{0, 1, 0}, 1000, 1010)) // 0 not found
	result, err := ExcludeJoinAndReverse([]int{0, 1, 0}, 1259, 26301)
	fmt.Println(result, err) // 62952 <nil>

	fmt.Println("Question 3")
	fmt.Println(SameJoinOf([]int{1, 2}, []int{1, 3}))       // [1]
	fmt.Println(SameJoinOf([]int{1, 2, 2}, []int{1, 2, 4})) // [2 1]

	fmt.Println("Question 4")
	fmt.Println(CountEachRepeatNumber(1223334)) // map[1:1 2:2 3:3 4:1]

	fmt.Println("Question 5")
	fmt.Println(MultiplyWithTheSide(1, 2, 3, 4))    // 14
	fmt.Println(MultiplyWithTheSide(1, 2, 3, 4, 5)) // 39
}

func SearchTheMissingOne(arr []int) int {
	hasPositive := false
	hasNegative := false
	positives := make(map[int]bool)
	negatives := make(map[int]bool)

	for _, v := range arr {
		if v > 0 {
			hasPositive = true
			positives[v] = true
		}

		if v < 0 {
			hasNegative = true
			negatives[v] = true
		}
	}

	if hasPositive && !hasNegative {
		for i := 1; ; i++ {
			if !positives[i] {
				return i
			}
		}
	}

	if !hasPositive && hasNegative {
		for i := -1; ; i-- {
			if !negatives[i] {
				return i
			}
		}
	}

	return 0
}

func ExcludeJoinAndReverse(exclude []int, nums ...int) (int, error) {
	excludeMap := make(map[int]bool)
	for _, v := range exclude {
		excludeMap[v] = true
	}

	allExclude := []string{}
	for _, num := range nums {
		numExclude := []string{}
		textNum := strconv.Itoa(num)
		for i := 0; i < len(textNum); i++ {
			digitNum := int(textNum[i] - '0')
			if !excludeMap[digitNum] {
				numExclude = append(numExclude, string(textNum[i]))
			}
		}
		allExclude = append(allExclude, numExclude...)
	}

	for i, j := 0, len(allExclude)-1; i < j; i, j = i+1, j-1 {
		allExclude[i], allExclude[j] = allExclude[j], allExclude[i]
	}

	resultJoin := strings.Join(allExclude, "")
	if resultJoin == "" {
		return 0, fmt.Errorf("not found")
	}

	data, err := strconv.Atoi(resultJoin)
	if err != nil {
		return 0, err
	}

	return data, nil
}

func SameJoinOf(numsA, numsB []int) []int {
	mapB := make(map[int]bool)
	mapIdentic := make(map[int]bool)
	for _, v := range numsB {
		mapB[v] = true
	}

	for _, v := range numsA {
		if mapB[v] {
			mapIdentic[v] = true
		}
	}

	resultSlice := make([]int, 0, len(mapIdentic))
	for m := range mapIdentic {
		resultSlice = append(resultSlice, m)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(resultSlice)))

	return resultSlice
}

func CountEachRepeatNumber(num int) map[int]int {
	result := make(map[int]int)
	numStr := strconv.Itoa(num)

	for i := 0; i < len(numStr); i++ {
		digit := int(numStr[i] - '0')
		result[digit]++
	}
	return result
}

func MultiplyWithTheSide(nums ...int) int {
	jumlah := 0

	for i := 0; i < len(nums); i += 2 {
		var hasil int

		if i+1 < len(nums) {
			hasil = nums[i] * nums[i+1]
		} else {
			hasil = nums[i] * nums[i]
		}

		jumlah += hasil
	}

	return jumlah
}
