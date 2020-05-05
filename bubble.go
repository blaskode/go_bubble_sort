
package main

import (
	"fmt"
)

var smp_sz int = 10
var prompt string = "Enter an integer: "

func getNums(nums []int) []int{
	for i := 0; i < smp_sz; i++ {
		fmt.Print(prompt);
		var in int
		fmt.Scan(&in)
		nums = append(nums, in)
	}
	return nums
}

func bubbleSort(nums []int) {
	swapped := true
	for swapped == true {
		swapped = swap(nums)
	}
}

func swap(nums []int) bool{
	swap := false
	var smaller int
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i+1] < nums[i]{
			smaller = nums[i+1]
			nums[i+1] = nums[i]
			nums[i] = smaller
			swap = true
		}
	}
	return swap
}


func prtSli(sli []int){
	for _, v := range sli {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}

func main(){
	fmt.Printf("You will now be prompted for %d integers.\n", smp_sz)
	nums := make([]int, 0)
	nums = getNums(nums)
	bubbleSort(nums)
	prtSli(nums)
}


