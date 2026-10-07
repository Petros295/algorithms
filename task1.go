package main

import "fmt"

func main() {
	nums := []int{122, 4, 1, 6, 23}

	length := len(nums)
	fmt.Println("Длина", length)

	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

	}
	fmt.Println("По возрастанию", nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println("По убыванию", nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
