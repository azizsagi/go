package main

import (
	"fmt"
	"slices"
)

func main() {

	// var prices = [5]int{1, 2, 3, 4, 5}
	// var profit int = 0
	// var buy int = prices[0]

	// for i := 0; i < len(prices); i++ {
	// 	if prices[i] < buy {
	// 		buy = prices[i]
	// 	}

	// 	if profit < (prices[i] - buy) {
	// 		profit = prices[i] - buy
	// 	}
	// }

	// fmt.Println(profit)

	// var prices = [...]int{1, 2, 3, 4, 5}
	// buy1 := 1
	// buy2 := 1
	// sell1 := 0
	// sell2 := 0

	// for i := 0; i < len(prices); i++ {

	// 	buy1 = min(buy1, prices[i])
	// 	sell1 = max(sell1, prices[i]-buy1)

	// 	buy2 = min(buy2, prices[i]-sell1)
	// 	sell2 = max(sell2, prices[i]-buy2)
	// }

	// fmt.Println(sell2)

	//var s string = "MAziziz"
	//sort.Strings(s)

	// var strs = []string{"adads", "dasdad", "dasdasd"}

	// slices.Sort(strs)

	// fmt.Print("value is", strs[len(strs)-1])

	// fmt.Println(strings.Count(s, "z"))

	// for i, c := range s {
	// 	if strings.Count(s, string(c)) == 1 {
	// 		fmt.Println("found at", i)
	// 	}
	// }

	// Write a function to find the longest common prefix string amongst an array of strings.

	// If there is no common prefix, return an empty string "".

	// Sort > get first and last item from array > compare > return common string

	strs := []string{"flower", "flow", "flight"}

	ans := ""
	slices.Sort(strs)
	first := strs[0]
	last := strs[len(strs)-1]

	min_range := min(len(first), len(last))

	for i := 0; i < min_range; i++ {
		if first[i] != last[i] {
			fmt.Println(ans)
		}
		ans = ans + string(first[i])
	}

	fmt.Println(ans)

}
