package pool

import (
	"context"
	"errors"
)

/*
	处理返回结果
*/

// resultPair 任务返回的结果
type resultPair[T any] struct {
	val T
	err error
}

type Future[T any] struct {
	ch chan resultPair[T]
}

// Result 返回结果
func (f *Future[T]) Result(ctx context.Context) (T, error) {
	select {
	case <-ctx.Done():
		//取消，返回默认值
		var zero T
		return zero, ctx.Err()
	case res, ok := <-f.ch:
		if !ok {
			var zero T
			return zero, errors.New("result channel closed")
		}
		return res.val, res.err
	}
}
