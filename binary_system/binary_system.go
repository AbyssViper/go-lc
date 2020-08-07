package main

import (
	"fmt"
	"go-lc/binary_system"
	"go-lc/binary_system/base"
)

func main() {
	base.XOR()
	base.ChangeValueByXOR(10, 100)
	base.RemoveLastOne(10)
	base.GetLastOne(10)
	topic()
}

func topic() {
	fmt.Println(binary_system.SingleNumber([]int{2, 1, 3, 3, 2}))
	fmt.Println(binary_system.SingleNumberII([]int{2, 2, 1, 3, 3, 3, 2}))
}
