/**
  根据索引删除切片里面的元素
  by hukg，2023
**/

package slice

// Delete 删除切片中索引为 index 的元素-泛型
// comparable 支持比较（== 或 !=）的类型：int、float、string、bool等
// 复制次数：n-1次复制
func Delete[T any](arr []T, index int) ([]T, error) {
	//判断索引
	length := len(arr)
	if index < 0 || index >= length {
		return nil, ErrIndexOfRange(length, index)
	}

	//s := append(arr[:index], arr[index+1:]...)
	//Shrink(s)
	//return s[:length-1], nil
	for i := index; i+1 < length; i++ {
		arr[i] = arr[i+1]
	}
	Shrink(arr)
	return arr[:length-1], nil
}

// Delete_Generic_NoSort 高性能方法，顺序会打乱
// 不管长度多少，只需要复制一次
func Delete_NoSort[T any](arr []T, index int) ([]T, error) {
	//判断索引
	length := len(arr)
	if index < 0 || index >= length {
		return nil, ErrIndexOfRange(length, index)
	}
	//把index索引赋值成最后一个元素
	arr[index] = arr[length-1]
	//返回 0至length-1
	return arr[:length-1], nil
}

