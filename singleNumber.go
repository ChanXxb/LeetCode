package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	loop1:
	for {
		i := nums[0]
		loop:
		for n ,comp := range nums[1:]{
			if i == comp{
				nums = nums[1:]

				nums = append(nums[:n] ,nums[n+1:]...)
				if len(nums) == 1{
					break loop1
				}
				break loop
			}
		}
	}
	return nums[0]
}

func singleNumberA(nums []int)int{
	var result int
	mapA := make(map[int]int)

	for _ ,i := range nums{
		mapA[i] = mapA[i]+1
	}

	for i ,j:= range mapA{
		if j == 1{
			result = i
		}
	}
	return result
}


func main(){
	var nums  = []int{1,2,2,1,3,4,3}

	fmt.Println(singleNumberA(nums))
}
