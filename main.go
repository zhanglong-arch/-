package main

import (
	"fmt"
	"math/big"
)

func main(){
	a := 1
	b := 3
	fmt.Println(a + b)
	var a1 int
	a1 = 3
	var b1 int
	b1 = 5
	fmt.Println(a1 + b1)

	//new内置函数：新建一个big.Int类型，并初始化，返回一个指针
	a2 := big.NewInt(1)
	b2 := big.NewInt(3)
	//a2 = a2.Add(a2, b2)
	//比较：compire
	//一共有三个返回值：-1，0，1
	/**
	 *  x < y  -1
	 *  x = y  0
	 *  x > y  1
	 */
	rs := a2.Cmp(b2)
	fmt.Println("a2和b2的大小关系：", rs)
	fmt.Println("两个大整数相加：", a2)

	a3 := big.NewInt(4)
	a3 = a3.Lsh(a3,4)
	fmt.Println("a3左移1位后的值：", a3)
}
