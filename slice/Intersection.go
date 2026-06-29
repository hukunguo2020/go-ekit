/**
  求两个切片的交集
  by hukg，2023
**/

package slice

import "fmt"

// Intersect 求切片的交集（去重,comparable）
func Intersect[T comparable](arr, arr2 []T) []T {

	src := make([]T, 0, len(arr))
	mp := make(map[T]struct{}, len(arr))

	for _, v := range arr {
		// 先检查是否已经处理过（去重）
		_, isOk := mp[v]
		if isOk {
			continue
		}

		for _, v2 := range arr2 {
			if v == v2 {
				//两个切片里面都有的元素，添加
				src = append(src, v)
				mp[v] = struct{}{}
			}
		}
	}
	return src
}

// IntersectFunc 求切片的交集（去重, any）
func IntersectFunc[T any, K comparable](arr, arr2 []T, getKey func(T) K) []T {

	src := make([]T, 0, len(arr))
	mp := make(map[K]struct{}, len(arr2))

	// 将 arr2 转换为 mapx，提高查找效率
	arr2Map := make(map[K]struct{}, len(arr2))
	for _, v := range arr2 {
		key := getKey(v)
		arr2Map[key] = struct{}{}
	}

	for _, v := range arr {
		key := getKey(v)
		// 先检查是否已经处理过（去重）
		_, isOk := mp[key]
		if isOk {
			continue
		}

		if _, exixt := arr2Map[key]; exixt {
			//两个数组里都存在
			src = append(src, v)
			mp[key] = struct{}{}
		}
	}
	return src
}

func Intersect_Demo() {
	arr := []string{"aa", "ab", "ac", "ad", "aa"}
	arr2 := []string{"ab", "ae", "dg", "aa"}
	src := Intersect(arr, arr2)
	fmt.Printf("切片的交集 %v \n", src)

	type Persion struct {
		Name    string
		Age     int
		Address []string
	}
	arr3 := []Persion{Persion{Name: "aaa", Age: 10}, Persion{Name: "bbb", Age: 12}, Persion{Name: "aaa", Age: 14}, Persion{Name: "aaa", Age: 14}, Persion{Name: "bbb", Age: 16}}
	arr4 := []Persion{{Name: "aaa", Age: 14}, {Name: "acc", Age: 14}}
	src3 := IntersectFunc(arr3, arr4, func(p Persion) string {
		return fmt.Sprintf("%v_%v", p.Name, p.Age)
	})
	fmt.Printf("切片的交集 %+v \n", src3)

}
