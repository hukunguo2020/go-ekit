package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDelete(t *testing.T) {
	//定义测试用例结构体
	testCases := []struct {
		name   string
		input  []string
		Index  int
		output []string
		err    error
	}{
		//准备数据
		{
			name:   `删除 索引 3`,
			input:  []string{"a", "a11", "b22", "c33", "d44"},
			Index:  3,
			output: []string{"a", "a11", "b22", "d44"},
			err:    nil,
		},
		{
			name:   `删除 索引 10，报错索引超出长度`,
			input:  []string{"a", "a11", "b22", "c33", "d44"},
			Index:  10,
			output: nil,
			err:    ErrIndexOfRange(5, 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Delete(tc.input, tc.Index)
			assert.Equal(t, tc.output, res)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestDelete_NoSort(t *testing.T) {
	//定义测试用例结构体
	testCases := []struct {
		name   string
		input  []string
		Index  int
		output int //因为顺序会打乱，直接判断元素个数
		err    error
	}{
		//准备数据
		{
			name:   `删除 索引 3`,
			input:  []string{"a", "a11", "b22", "c33", "d44"},
			Index:  3,
			output: 4,
			err:    nil,
		},
		{
			name:   `删除 索引 10，报错索引超出长度`,
			input:  []string{"a", "a11", "b22", "c33", "d44"},
			Index:  10,
			output: 0,
			err:    ErrIndexOfRange(5, 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Delete_NoSort(tc.input, tc.Index)

			if err != nil {
				assert.Equal(t, tc.output, len(res))
			}
			assert.Equal(t, tc.err, err)
		})
	}
}
