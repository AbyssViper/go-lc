package main

import (
	"fmt"
	"go-lc/data_structure/binary_system"
)

func main() {
	binary_system.XOR()
	binary_system.ChangeValueByXOR(10, 100)
	binary_system.RemoveLastOne(10)
	binary_system.GetLastOne(10)
	topic()
}

func topic() {
	fmt.Println(binary_system.SingleNumber([]int{2, 1, 3, 3, 2}))
	fmt.Println(binary_system.SingleNumberII([]int{2, 2, 1, 3, 3, 3, 2}))
}
