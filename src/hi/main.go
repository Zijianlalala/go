package main

import (
	"fmt"
	"math"
)

func main()  {
	nums := [4]int{1,2,3,4}
	mySlice := nums[:2]
	mySlice[0] = 100
	fmt.Println("mySlice length(before): ", len(mySlice))
	mySlice = append(mySlice, 1)
	fmt.Println("mySlice length(after): ", len(mySlice))
	fmt.Println("mySlice", mySlice[0])
	fmt.Println("nums", nums[0])
	fmt.Println(max(1,2,3,4))
}

func max(numbers ...float64) float64{
	maxValue := math.Inf(-1)
	for _, val := range numbers {
		if val > maxValue {
			maxValue = val
		}
	}
	return  maxValue
}