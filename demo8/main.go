package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode"
)

// 综合案例：结构体 转 map
type Data struct {
	Username string
	SmzdmId int
	unused string
}

func struct2Map(v interface{}) (retval map[string]string) {
	retval = make(map[string]string)

	structsVal := reflect.ValueOf(v)
	if structsVal.Type().Kind() == reflect.Ptr {	// 解引用
		structsVal = structsVal.Elem()
	}

	for i := 0; i < structsVal.NumField(); i ++ {
		fieldName := structsVal.Type().Field(i).Name	// 字段名
		fieldVal := structsVal.Field(i)	// 字段值
		if unicode.IsUpper(([]rune(fieldName))[0]) {	// 只导出大写字段
			if fieldVal.Type().Kind() == reflect.Int {
				retval[fieldName] = strconv.Itoa(int(fieldVal.Int()))
			} else if fieldVal.Type().Kind() == reflect.String {
				retval[fieldName] = fieldVal.String()
			}
		}
	}
	return
}

func main()  {
	data := &Data{"test", 12345, "xxxx"}

	m := struct2Map(data)
	fmt.Println(m)
}
