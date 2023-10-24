package GoUtil

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
)

func List2Map[M interface{}, N constraints.Integer | constraints.Float | string](arr []M, filedName string, filed N) map[N]M {
	res := make(map[N]M)

	for idx := range arr {
		item := arr[idx]

		key := GetFieldValue(item, filedName)
		if key == nil {
			panic(fmt.Sprintf("not found field:%v", filedName))
		}

		res[key.(N)] = item
	}

	return res
}

func GetFieldValue(obj interface{}, fieldName string) interface{} {
	value := reflect.ValueOf(obj)

	// 检查 obj 是否为指针类型，如果是，获取指针指向的值
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// 检查 obj 是否为结构体类型
	if value.Kind() != reflect.Struct {
		return nil
	}

	// 获取字段值
	field := value.FieldByName(fieldName)
	if !field.IsValid() {
		return nil
	}

	return field.Interface()
}
