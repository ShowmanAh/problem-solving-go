package main

import "fmt"

func main() {
	//sum := twoSum([]int{2, 7, 11, 15}, 9)
	sum := twoSum2([]int{2, 7, 11, 15}, 9)
	fmt.Println(sum)

}

// big o of n notation
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0}
}

// use pointer big o notation
func twoSum2(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	fmt.Printf("left: %d, right: %d\n", left, right)
	for left < right {
		sum := nums[left] + nums[right]
		fmt.Println(sum)
		if sum == target {
			fmt.Println(sum, target)
			return []int{left, right}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	fmt.Println("end")
	return []int{0, 1}
}
