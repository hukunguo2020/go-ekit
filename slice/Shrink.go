/**
  切片的缩容
**/

package slice

// Shrink 缩容
func Shrink[T any](arr []T) []T {
	c, l := cap(arr), len(arr) //获取容量和长度
	n, changed := calCapacity(c, l)
	//如果没有缩容
	if !changed {
		return arr
	}
	//缩容之后，创建新的切片并把数据追加
	s := make([]T, 0, n)
	s = append(s, arr...)
	return s
}

// 缩容算法
func calCapacity(c, l int) (int, bool) {
	// 容量 <= 64 缩不缩都无所谓，内存浪费不了多少
	// 你可以考虑调大这个阈值，或者调小这个阈值
	if c <= 64 {
		return c, false
	}
	//如果容量打印2048，但元素不足一半
	//降低为 0.625，也就是 5/8
	//也就比一般多一点，和正向扩容的 1.25 倍相呼应
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}

	//如果在 2048以内，并且元素不足 1/4，直接缩减为一半
	if c < 2048 && (c/l >= 4) {
		return c / 2, false
	}

	//整个实现的核心是希望在后续少触发扩容的前提下，一次性释放尽可能多的内存
	return c, false
}
