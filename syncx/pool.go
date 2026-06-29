package syncx

import "sync"

// Pool 是对 sync.Pool 的简单封装
type Pool[T any] struct {
	p sync.Pool
}

// NewPool 创建一个 Pool实例
// factory 类型必须返回 T类型 的值，不能返回 nil
func NewPool[T any](factory func() T) *Pool[T] {
	return &Pool[T]{
		p: sync.Pool{
			New: func() any {
				return factory()
			},
		},
	}
}

// Get 取出一个元素
func (p *Pool[T]) Get() T {
	return p.p.Get().(T)
}

// Put 放回一个元素
func (p *Pool[T]) Put(x T) {
	p.p.Put(x)
}
