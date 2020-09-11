package main

import "fmt"

// demo1：空接口
func main() {
	a := 1
	var i interface{} = a

	// 空接口由（类型，值）组成
	if v, ok := i.(int); ok {
		fmt.Println(v)
	}
}