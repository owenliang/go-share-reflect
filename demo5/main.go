package main

import (
	"fmt"
	"reflect"
)

// 反射值
func main() {
	a := 5
	var i interface{} = &a
	fmt.Printf("%d\n", &a)

	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		var vp reflect.Value = reflect.ValueOf(i)
		fmt.Println(int(vp.Pointer()))	// 打印指针

		var v reflect.Value = vp.Elem()
		fmt.Println(v.Int())	// 指针解引用，打印值
	}
}