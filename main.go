package main

import (
	"fmt"

	"github.com/kvnbanunu/leetcode/leet"
)

func main() {
	arr := []int{2, 20, 4, 10, 3, 4, 5}
	ans := leet.LongestConsecutive(arr)
	fmt.Println(ans)
}
