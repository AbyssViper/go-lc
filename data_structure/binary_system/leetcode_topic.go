package binary_system

/*
Number: 136
URL: https://leetcode-cn.com/problems/single-number/
Content: 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
Point:	a^a = 0 ; a^0 = a
Process: result ^ same ^ same -> result ^ 0 -> result
*/
func SingleNumber(nums []int) int {
	result := 0
	for _, value := range nums {
		result ^= value
	}
	return result
}

/*
Number: 137
URL: https://leetcode-cn.com/problems/single-number-ii/
Content: 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
Point: a&a^a = 0
*/
func SingleNumberII(nums []int) int {
	result := 0
	
	return result
}
