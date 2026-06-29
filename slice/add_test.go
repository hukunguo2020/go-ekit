package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdd(t *testing.T) {
	//定义测试用例结构体
	testCases := []struct {
		name   string   //用例名称
		input  []string //输入的切片数据
		addVal string   //添加的数据
		index  int      //添加的索引
		output []string //输出的切片数据
		Err    error    //预期的错误
	}{
		//准备数据
		{
			name:   "index 0",
			input:  []string{"aa", "bb", "cc"},
			addVal: "abc",
			index:  0,
			output: []string{"abc", "aa", "bb", "cc"},
		},
		{
			name:   "索引错误",
			input:  []string{"aa", "bb", "cc"},
			addVal: "abc",
			index:  10,
			Err:    ErrIndexOfRange(3, 10), //预期错误
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := Add(tc.input, tc.index, tc.addVal)

			assert.Equal(t, tc.Err, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.output, result)

		})
	}
}
