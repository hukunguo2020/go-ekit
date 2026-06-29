package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFind(t *testing.T) {
	//定义测试用例结构体
	testCases := []struct {
		name       string
		input      []string
		Value      string
		wantOutput string
		wantOk     bool
	}{
		//准备数据
		{
			name:       `查找元素 a11`,
			input:      []string{"a", "b", "c", "d", "a11", "b22", "c33", "d44"},
			Value:      `a11`,
			wantOutput: `a11`,
			wantOk:     true,
		},
		{
			name:       `查找元素 aaaaa11`,
			input:      []string{"a", "a11", "b22", "c33", "d44"},
			Value:      `aaaaa11`,
			wantOutput: ``,
			wantOk:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, ok := Find(tc.input, tc.Value)
			assert.Equal(t, tc.wantOutput, res)
			assert.Equal(t, tc.wantOk, ok)
		})
	}
}

func TestFindFunc(t *testing.T) {
	testCases := []struct {
		name  string
		input []Number
		match matchFunc[Number]

		wantVal Number
		found   bool
	}{
		{
			name: "找到了",
			input: []Number{
				{val: 123},
				{val: 234},
			},
			match: func(src Number) bool {
				return src.val == 123
			},
			wantVal: Number{val: 123},
			found:   true,
		},
		{
			name: "没找到",
			input: []Number{
				{val: 123},
				{val: 234},
			},
			match: func(src Number) bool {
				return src.val == 456
			},
		},
		{
			name: "nil",
			match: func(src Number) bool {
				return src.val == 123
			},
		},
		{
			name:  "没有元素",
			input: []Number{},
			match: func(src Number) bool {
				return src.val == 123
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, found := FindFunc[Number](tc.input, tc.match)
			assert.Equal(t, tc.found, found)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestFindAllFunc(t *testing.T) {
	testCases := []struct {
		name  string
		input []Number
		match matchFunc[Number]

		wantVals []Number
	}{
		{
			name:  "没有符合条件的",
			input: []Number{{val: 2}, {val: 4}},
			match: func(src Number) bool {
				return src.val%2 == 1
			},
			wantVals: []Number{},
		},
		{
			name:  "找到了",
			input: []Number{{val: 2}, {val: 3}, {val: 4}},
			match: func(src Number) bool {
				return src.val%2 == 1
			},
			wantVals: []Number{{val: 3}},
		},
		{
			name: "nil",
			match: func(src Number) bool {
				return src.val%2 == 1
			},
			wantVals: []Number{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			vals, _ := FindAllFunc[Number](tc.input, tc.match)
			assert.Equal(t, tc.wantVals, vals)
		})
	}
}

type Number struct {
	val int
}
