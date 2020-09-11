package main

import "fmt"

// 空的空接口
func main() {
	var i interface{} = nil
	if i == nil {
		fmt.Println("空接口为空")
	}

	var x interface{} = (*int)(nil)
	if x == nil {	// 不成立
		fmt.Println("指针为空")
	}
}

/*
	var x interface{} = (*int)(nil)
	if v, ok := x.(*int); ok && v == nil {	// 不成立
		fmt.Println("指针为空")
	}
 */