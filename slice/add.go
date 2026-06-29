/**
  根据索引，增加切片的元素
  by hukg，2023
**/

package slice

// Add 添加元素到指定位置
func Add[T any](arr []T, index int, value T) ([]T, error) {
	//判断索引
	length := len(arr)
	if index < 0 || index > length {
		//添加元素所以可以与长度相等
		return nil, ErrIndexOfRange(length, index)
	}

	arr = append(arr, value)
	newLength := len(arr)
	for i := newLength - 1; i > index; i-- {
		if i > 0 {
			arr[i] = arr[i-1]
		}
	}
	arr[index] = value

	return arr, nil
}
