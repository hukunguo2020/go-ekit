/**
  查找切片的元素
  by hukg，2023
**/

package slice

// Find 查找匹配到的第一个元素（仅支持可比较类型）
func Find[T comparable](arr []T, value T) (T, bool) {
	for _, v := range arr {
		if v == value {
			return v, true
		}
	}
	var zore T //使用零值
	return zore, false
}

// FindFunc 查找匹配到的第一个元素（支持任何类型，自定义查找条件）
func FindFunc[T any](arr []T, find func(T) bool) (T, bool) {
	for _, v := range arr {
		if find(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// FindAllFunc 查找匹配到的所有元素（支持任何类型，自定义查找条件）
func FindAllFunc[T any](arr []T, find func(T) bool) ([]T, bool) {
	//创建一个切片，长度是原长度/8(>>3表示位运算右移3位)
	var src []T = make([]T, 0, len(arr)>>3+1)
	for _, v := range arr {
		if find(v) {
			src = append(src, v)
		}
	}
	if len(src) > 0 {
		return src, true
	}
	return src, false
}
