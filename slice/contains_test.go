package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestContains(t *testing.T) {
	//定义测试用例结构体
	testCases := []struct {
		name   string
		input  []string
		Val    string
		output bool
		err    error
	}{
		//准备数据
		{
			name:   "切片中是否有 “a11”",
			input:  []string{"a", "a11", "b22", "c33", "d44"},
			Val:    "a11",
			output: true,
		},
		{
			name:   "切片中是否有 “a13”",
			input:  []string{"a", "a11", "b22", "c33", "d44"},
			Val:    "a13",
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Contains[string](tc.input, tc.Val)
			assert.Equal(t, tc.output, res)
		})
	}
}

func TestContainsFunc(t *testing.T) {
	type Person struct {
		Name    string
		Age     int
		Address []string
	}

	//定义测试用例结构体
	testCases := []struct {
		name   string
		input  []Person
		Val    Person
		output bool
		err    error
	}{
		{
			name:   "Person 中是否有 {Name: \"aaa\", Age: 14}",
			input:  []Person{Person{Name: "aaa", Age: 10}, Person{Name: "bbb", Age: 12}, Person{Name: "aaa", Age: 14}, Person{Name: "bbb", Age: 16}},
			Val:    Person{Name: "aaa", Age: 14},
			output: true,
		},
		{
			name:   "Person 中是否有 {Name: \"aaa\", Age: 16}",
			input:  []Person{Person{Name: "aaa", Age: 10}, Person{Name: "bbb", Age: 12}, Person{Name: "aaa", Age: 14}, Person{Name: "bbb", Age: 16}},
			Val:    Person{Name: "aaa", Age: 16},
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsFunc[Person](tc.input, func(p Person) bool {
				return p.Name == tc.Val.Name && p.Age == tc.Val.Age
			})
			assert.Equal(t, tc.output, res)
		})
	}
}
