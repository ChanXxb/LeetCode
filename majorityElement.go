package main

import "fmt"

func majorityElement1(nums []int) int {
	mapA := make(map[int]int)
	length := len(nums)
	var result int

	for _ ,i := range nums{
		mapA[i] = mapA[i]+1
	}

	for i ,j := range mapA{
		if j >length/2{
			result = i
		}
	}
	return result
}

func majorityElement2(nums []int) int{
	x := nums[0]// 当前count的major数
	counts := 1 // 当前的权重

loop:
	for i:= 0 ;i<len(nums)-1;i++{
		if nums[i] != nums[i+1]{
			counts = counts-1
			if counts<0{
				x = nums[i+1]
			}
		}else {
			counts =counts+1
			if counts>len(nums)/2{
				break loop
			}
		}
	}
	return x
}


func main(){
	var nums = []int{7,7,7,3,4,1,7}

	fmt.Println(majorityElement1(nums))

	fmt.Println(majorityElement2(nums))
}
