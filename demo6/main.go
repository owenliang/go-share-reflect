package main

import (
	"fmt"
	"reflect"
)

// 反射值并修改它
func main() {
	a := 1
	i := interface{}(&a)

	var v reflect.Value = reflect.ValueOf(i)
	v.Elem().SetInt(5)
	fmt.Println(a)
}
