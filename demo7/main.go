package main

import (
	"fmt"
	"reflect"
)

// 综合案例：遍历Map
func main() {
	a := map[string]int{"xx": 5}
	i := interface{}(a)

	var m reflect.Value = reflect.ValueOf(i)
	if m.Type().Kind() == reflect.Map {
		fmt.Println("大小: ", m.Len())
		// 迭代
		iter := m.MapRange()
		for iter.Next() {
			k := iter.Key()	// key
			v := iter.Value()	//value
			fmt.Println("反射：", k.String(), v.Int())
			m.SetMapIndex(k, reflect.ValueOf(int(v.Int()) + 1))
		}
	}

	fmt.Println(a)
}