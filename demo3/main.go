package main

import (
	"fmt"
	"reflect"
)

// 反射类型
func main() {
	a := 1
	var i interface{} = &a

	var t reflect.Type = reflect.TypeOf(i)
	fmt.Println(t.Kind())	// 获取类型枚举

	var t1 reflect.Type = t.Elem()	//  解引用
	fmt.Println(t1.Kind())
}