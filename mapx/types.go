package mapx

type mapi[K any, V any] interface {
	// Put 更新或添加值
	Put(key K, value V) error
	// Get 获取值
	Get(key K) (V, bool)
	// Keys 返回所有的键
	// 注意，当你调用多次拿到的结果不一定相等
	// 取决于具体实现
	Keys() []K
	// Values 返回所有的值
	// 注意，当你调用多次拿到的结果不一定相等
	// 取决于具体实现
	Values() []V
	// Delete 删除
	// 第一个返回的是
	Delete(key K) (V, bool)

	Len() int64

	Iterate(cb func(key K, v V) bool)
}
