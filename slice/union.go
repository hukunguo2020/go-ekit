/**
  求两个切片的并集
  by hukg，2023
**/

package slice

import "fmt"

// Union 计算两个切片的并集
func Union[T comparable](arr, arr2 []T) []T {
	length := len(arr)
	length2 := len(arr2)
	if length == 0 {
		return arr2
	}
	if length2 == 0 {
		return arr
	}

	src := make([]T, 0, length+length2) //创建一个切片，长度 = 两个切片长度之和
	mp := make(map[T]struct{}, length+length2)

	appendSrc := func(v T) {
		if _, exist := mp[v]; !exist {
			mp[v] = struct{}{}
			src = append(src, v)
		}
	}

	for _, v := range arr {
		appendSrc(v)
	}
	for _, v := range arr2 {
		appendSrc(v)
	}
	return src
}

// Union 计算两个切片的并集(去重+并集)
func UnionFunc[T any, K comparable](arr, arr2 []T, getkey func(T) K) []T {
	length := len(arr) + len(arr2)

	src := make([]T, 0, length)        //创建一个切片，长度 = 两个切片之和
	mp := make(map[K]struct{}, length) // mapx 里面使用 mapx[K]struct{}，而不是 mapx[K]bool，理论上struct{}是零值，不占内容bool占一个字节

	appendKey := func(v T) {
		key := getkey(v) //获取key
		if _, exists := mp[key]; exists {
			return
		}
		mp[key] = struct{}{}
		src = append(src, v)
	}

	for _, v := range arr {
		appendKey(v)
	}
	for _, v := range arr2 {
		appendKey(v)
	}

	return src
}

func Union_Demo() {
	//arr := []int{1, 2, 3, 4, 5, 2}
	//arr2 := []int{1}
	//src := Union(arr, arr2)
	//fmt.Printf("切片的并集 %v \n", src)

	//arr := []int{1, 2, 3, 4, 5, 2}
	//arr2 := []int{1, 9}
	//src := UnionFunc(arr, arr2, func(value, value2 int) bool {
	//	return value2 == value
	//})
	//fmt.Printf("切片的并集 %v \n", src)

	//arr := []string{"aa", "ab", "ac", "ad", "aa"}
	//arr2 := []string{"ab", "ae", "dg"}
	//src := UnionFunc(arr, arr2, func(value, value2 string) bool {
	//	return value2 == value
	//})
	//fmt.Printf("切片的并集 %v \n", src)

	type Persion struct {
		Name    string
		Age     int
		Address []string
	}
	arr := []Persion{Persion{Name: "aaa", Age: 10}, Persion{Name: "bbb", Age: 12}, Persion{Name: "aaa", Age: 14}, Persion{Name: "aaa", Age: 14}, Persion{Name: "bbb", Age: 16}}
	arr2 := []Persion{{Name: "aaa", Age: 14}, {Name: "acc", Age: 14}}
	src := UnionFunc(arr, arr2, func(persion Persion) string {
		return fmt.Sprintf("%v_%v", persion.Name, persion.Age)
	})
	fmt.Printf("切片的并集 %+v \n", src)

}
