package binary_system

import "fmt"

// 利用异或特点
// a = a ^ 0 = 0 ^ a
// 0 = a ^ a  -> a = a ^ 0 = a ^ b ^ b
func XOR() {
	a := -100
	b := 200
	fmt.Println("a: ", a)
	fmt.Printf("a^0: %v \n0^a: %v  \n", a^0, 0^a)
	fmt.Println("a^a: ", a^a)
	fmt.Println("a^b^b: ", a^b^b)
}

func ChangeValueByXOR(num1 int, num2 int) {
	num1 = num1 ^ num2
	num2 = num1 ^ num2 // new_num1 ^ old_num2 -> (old_num1 ^ old_num2) ^ old_num2 -> old_num1
	num1 = num1 ^ num2 // new_num1 ^ new_num2 -> (old_num1 ^ old_num2) ^ old_num1 -> old_num2
	fmt.Printf("Chang value: %v, %v\n", num1, num2)
}

// 利用按与操作的特性，减去1以后，最后一个1前面的与操作以后不变，最后一个1以及到最后，与操作以后都会为0
func RemoveLastOne(n int) {
	fmt.Println("Remove value's binary_system last 1:", n & (n - 1))
}

// 通过与操作去出最后一个1以后，进行异或可以排除仅剩最后一个1的值
func GetLastOne(n int) {
	fmt.Println("Get value's binary_system last 1:", (n & (n - 1)) ^ n)
}
