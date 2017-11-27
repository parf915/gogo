package main 

import (
	"fmt"
)

func proc(id int, in, out chan int) {
	nums := make([]int, 0, 100)
	for num := range in {
		nums = append(nums, num)
	}
	lowerIndex := id
	for i := id+1; i < len(nums); i++ {
		if nums[i] < nums[lowerIndex] {
			lowerIndex = i
		}
	}
	for i := 0; i < id; i++ {
		out <- nums[i]
	}
	out <- nums[lowerIndex]
	for i := id; i < len(nums); i++ {
		if i != lowerIndex {
			out <- nums[i]
		}
	}
	close(out)
}

func main() {
	nums := []int{ 10, 20, 5, 7, 13, 17, 1, 15, 8, 12 }
	channels := make([]chan int, len(nums)+1)
	for i := 0; i < len(channels); i++ {
		channels[i] = make(chan int)
	}
	for i := range nums {
		go proc(i, channels[i], channels[i+1])
	}
	for _, num := range nums {
		channels[0] <- num
	}
	close(channels[0])
	for x := range channels[len(channels)-1] {
		fmt.Printf("%d ", x)
	}
	fmt.Println("")
}