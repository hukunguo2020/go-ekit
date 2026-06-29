package slice

// ToMap 构造map，返回空结构体
func ToMap[T any, K comparable](arr []T, getKey func(T) K) map[K]struct{} {
	mp := make(map[K]struct{}, len(arr))
	for _, v := range arr {
		mp[getKey(v)] = struct{}{} //使用空结构体，减少内存占用
	}
	return mp
}

// ToMapV 构造map，返回原结构体数据
func ToMapV[T any, K comparable](arr []T, getKey func(T) K) map[K]T {
	mp := make(map[K]T, len(arr))
	for _, v := range arr {
		mp[getKey(v)] = v //使用原结构体数据
	}
	return mp
}

// Map 是一个通用函数，用于将源切片转换为目标切片
// Src: 源元素类型
// Dst: 目标元素类型
// src: 源切片
// m: 转换函数，接收索引和源元素，返回目标元素
func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	// 创建与源切片长度相同的目标切片
	dst := make([]Dst, len(src))
	//遍历源切片，应用转换函数
	for i, v := range src {
		//调用转换函数，转换成目标元素
		dst[i] = m(i, v)
	}
	return dst
}
