package main

import (
	"fmt"
	"reflect"
)

// 更多反射类型
func main() {
	a := map[string]int{"xx": 5}
	i := interface{}(a)

	var t reflect.Type = reflect.TypeOf(i)
	fmt.Println(t.Kind())	// 获取类型枚举

	var kt reflect.Type = t.Key()	// key的类型
	fmt.Println(kt.Kind())

	// 怎么获取value的类型呢？
}