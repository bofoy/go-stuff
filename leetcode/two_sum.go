package main

func twoSum(nums []int, target int) [2]int {
	complements := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, ok := complements[complement]; ok {
			return [2]int{complements[complement], i}
		}
		complements[nums[i]] = i
	}
	return [2]int{}
}

func main() {
	twoSum([]int{2, 7, 11, 14}, 9)
}
