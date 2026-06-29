/**
  @author:
  @date:
  @note:
**/

package slice

import "fmt"

// ErrIndexOfRange 索引下标不能小于0，或不能超出范围
func ErrIndexOfRange(length int, index int) error {
	return fmt.Errorf("索引不正确，长度：%d，索引：%d", length, index)
}
