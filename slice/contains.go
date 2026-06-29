/**
  判断切片里是否有某个元素
  by hukg，2023
**/

package slice

import "fmt"

// Contains 判断切片里是否有某个元素(仅支持可比较类型)
func Contains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsFunc 判断切片里是否有某个元素(自定义查找，支持任何类型)
func ContainsFunc[T any](arr []T, find func(T) bool) bool {
	for _, v := range arr {
		if find(v) {
			return true
		}
	}
	return false
}

func Contains_Demo() {
	//arr := []string{"a", "b", "c", "d", "a11", "b22", "c33", "d44"}
	//value := "a11"
	//
	//if Contains(arr, value) {
	//	fmt.Printf("%v 切片里面存在 %v \n", arr, value)
	//} else {
	//	fmt.Printf("%v 切片里面不存在 %v \n", arr, value)
	//}

	//arr := []string{"a", "b", "c", "d", "a11", "b22", "c33", "d44"}
	//value := "a11"
	//if ContainsFunc(arr, func(s string) bool {
	//	return s == value
	//}) {
	//	fmt.Printf("%v 切片里面存在 %v \n", arr, value)
	//} else {
	//	fmt.Printf("%v 切片里面不存在 %v \n", arr, value)
	//}
	//
	//arr1 := []int{111, 222, 333, 444, 555, 666}
	//value1 := 123
	//if ContainsFunc(arr1, func(s int) bool {
	//	return s == value1
	//}) {
	//	fmt.Printf("%v 切片里面存在 %v \n", arr1, value1)
	//} else {
	//	fmt.Printf("%v 切片里面不存在 %v \n", arr1, value1)
	//}

	type Persion struct {
		Name    string
		Age     int
		Address []string
	}
	arr2 := []Persion{Persion{Name: "aaa", Age: 10}, Persion{Name: "bbb", Age: 12}, Persion{Name: "aaa", Age: 14}, Persion{Name: "bbb", Age: 16}}
	value2 := Persion{Name: "aaa", Age: 14}
	if ContainsFunc(arr2, func(s Persion) bool {
		return s.Name == value2.Name && s.Age == value2.Age
	}) {
		fmt.Printf("%v 切片里面存在 %v \n", arr2, value2)
	} else {
		fmt.Printf("%v 切片里面不存在 %v \n", arr2, value2)
	}

}
