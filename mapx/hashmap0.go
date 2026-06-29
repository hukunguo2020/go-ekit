package mapx

//// Hashable 定义Hashable接口
//type Hashable interface {
//	// Code 返回该元素的哈希值
//	Code() uint64
//	// Equals 比较两个元素是否相等。如果返回 true，那么我们会认为两个键是一样的
//	Equals(key any) bool
//}
//
//// node 节点结构
//type node[T Hashable, ValType any] struct {
//	key     T
//	value   ValType
//	next    *node[T, ValType] //下一个节点的指针
//	visited bool              //用于标记遍历状态
//}
//
//// bucket 桶结构，封装链表操作
//type bucket[T Hashable, ValType any] struct {
//	head *node[T, ValType] //头指针
//	tail *node[T, ValType] //尾指针，便于尾部插入
//}
//
//type SimpleHashMap[T Hashable, ValType any] struct {
//	buckets  []*bucket[T, ValType] //使用切片而非map，便于控制容量
//	size     int                   //实际元素数量
//	capacity int                   //桶数量
//	nodePool *syncx.Pool[*node[T, ValType]]
//}
//
//// NewSimpleHashMap 创建新的HashMap
//func NewSimpleHashMap[T Hashable, ValType any](capacity int) *SimpleHashMap[T, ValType] {
//	if capacity <= 0 {
//		capacity = 16 //默认容量
//	}
//
//	//确保容量是2的幂，提高哈希分布性
//	capacity = roundUpToPowerOfTwo(capacity)
//
//	buckets := make([]*bucket[T, ValType], capacity)
//	for i := range buckets {
//		buckets[i] = &bucket[T, ValType]{}
//	}
//
//	return &SimpleHashMap[T, ValType]{
//		buckets:  buckets,
//		capacity: capacity,
//		nodePool: syncx.NewPool[*node[T, ValType]](func() *node[T, ValType] {
//			return &node[T, ValType]{}
//		}),
//	}
//}
//
//// bucketIndex 计算桶索引的辅助函数（把一个大的hash分配到有限的map里面，均匀分配）
//func (m *SimpleHashMap[T, ValType]) bucketIndex(hash uint64) int {
//	return int(hash % uint64(m.capacity))
//}
//
//// Put 添加/更新键值对
//func (m *SimpleHashMap[T, ValType]) Put(key T, val ValType) error {
//	hash := key.Code()
//	idx := m.bucketIndex(hash)
//	bkt := m.buckets[idx]
//
//	// 查找是否存在相同key
//	if node := m.findNodeInBucket(bkt, key); node != nil {
//		//存在，更新值
//		node.value = val
//		return nil
//	}
//
//	//不存在，创建新节点
//	newNode := m.newNode(key, val)
//
//	//插入到桶中
//	if bkt.head == nil {
//		//如果桶链表头为空，表示 newNode 是第一个节点
//		bkt.head = newNode
//		bkt.tail = newNode
//	} else {
//		//否则就是最后一个节点
//		bkt.tail.next = newNode
//		bkt.tail = newNode
//	}
//
//	m.size++
//	return nil
//}
//
//// Get 获取值
//func (m *SimpleHashMap[T, ValType]) Get(key T) (ValType, bool) {
//	hash := key.Code()
//	idx := m.bucketIndex(hash)
//	bkt := m.buckets[idx]
//
//	if node := m.findNodeInBucket(bkt, key); node != nil {
//		return node.value, true
//	}
//
//	var zero ValType
//	return zero, false
//}
//
//// Keys 返回所有键
//func (m *SimpleHashMap[T, ValType]) Keys() []T {
//	keys := make([]T, 0, m.size)
//
//	for _, bkt := range m.buckets {
//		current := bkt.head
//		for current != nil {
//			keys = append(keys, current.key)
//			current = current.next
//		}
//	}
//
//	return keys
//}
//
//// Values 获取所有值
//func (m *SimpleHashMap[T, ValType]) Values() []ValType {
//	values := make([]ValType, 0, m.size)
//
//	for _, bkt := range m.buckets {
//		current := bkt.head
//		for current != nil {
//			values = append(values, current.value)
//			current = current.next
//		}
//	}
//
//	return values
//}
//
//// Delete 删除键值对
//func (m *SimpleHashMap[T, ValType]) Delete(key T) (ValType, bool) {
//	hash := key.Code()
//	idx := m.bucketIndex(hash)
//	bkt := m.buckets[idx]
//
//	//处理桶为空的情况
//	if bkt.head == nil {
//		var zero ValType
//		return zero, false
//	}
//
//	// 处理头节点
//	if bkt.head.key.Equals(key) {
//		val := bkt.head.value
//		deleteNode := bkt.head
//
//		bkt.head = bkt.head.next
//		if bkt.head == nil {
//			bkt.tail = nil
//		}
//
//		m.recycleNode(deleteNode)
//		m.size--
//		return val, true
//	}
//
//	// 遍历链表查找要删除的节点
//	prev := bkt.head
//	current := bkt.head.next
//	for current != nil {
//		if current.key.Equals(key) {
//			val := current.value
//
//			prev.next = current.next
//			if current == bkt.tail {
//				bkt.tail = prev
//			}
//
//			m.recycleNode(current)
//			m.size--
//			return val, true
//		}
//
//		prev = current
//		current = current.next
//	}
//
//	var zero ValType
//	return zero, false
//}
//
//// Len 返回元素数量
//func (m *SimpleHashMap[T, ValType]) Len() int64 {
//	return int64(m.size)
//}
//
//// Iterate 遍历所有键值对
//func (m *SimpleHashMap[T, ValType]) Iterate(cb func(key T, val ValType) bool) {
//	// 重置所有节点的访问标记
//	m.resetVisitedFlags()
//
//	for _, bkt := range m.buckets {
//		current := bkt.head
//		for current != nil {
//			if current.visited {
//				break //防止循环链表导致的无限循环
//			}
//
//			current.visited = true
//			if !cb(current.key, current.value) {
//				return
//			}
//			current = current.next
//		}
//	}
//}
//
//// ContainsKey 检查是否包含指定键
//func (m *SimpleHashMap[T, ValType]) ContainsKey(key T) bool {
//	_, exists := m.Get(key)
//	return exists
//}
//
//// Clear 清空所有元素
//func (m *SimpleHashMap[T, ValType]) Clear() {
//	for i := range m.buckets {
//		//回收所有节点
//		current := m.buckets[i].head
//		for current != nil {
//			next := current.next
//			m.recycleNode(current)
//			current = next
//		}
//
//		m.buckets[i].head = nil
//		m.buckets[i].tail = nil
//	}
//
//	m.size = 0
//}
//
//// resetVisitedFlags 重置所有节点访问标记的辅助函数
//func (m *SimpleHashMap[T, ValType]) resetVisitedFlags() {
//	for _, bkt := range m.buckets {
//		current := bkt.head
//		for current != nil {
//			current.visited = false
//			current = current.next
//		}
//	}
//}
//
//// recycleNode 回收节点的辅助函数
//func (m *SimpleHashMap[T, ValType]) recycleNode(n *node[T, ValType]) {
//	var zoreKey T
//	var zoreVal ValType
//
//	n.key = zoreKey
//	n.value = zoreVal
//	n.next = nil
//	n.visited = false
//
//	m.nodePool.Put(n)
//}
//
//// newNode 创建新节点的辅助函数
//func (m *SimpleHashMap[T, ValType]) newNode(key T, val ValType) *node[T, ValType] {
//	n := m.nodePool.Get()
//	n.key = key
//	n.value = val
//	n.next = nil
//	n.visited = false
//	return n
//}
//
//// findNodeInBucket 在桶中查找节点的辅助函数
//func (m SimpleHashMap[T, ValType]) findNodeInBucket(bkt *bucket[T, ValType], key T) *node[T, ValType] {
//	current := bkt.head
//	for current != nil {
//		if current.key.Equals(key) {
//			return current
//		}
//		//不存在，找下一个；直到下一个为 nil
//		current = current.next
//	}
//	return nil
//}
//
//// 辅助函数：将数字向上取整到最近的2的幂
//func roundUpToPowerOfTwo(n int) int {
//	n--
//	n |= n >> 1
//	n |= n >> 2
//	n |= n >> 4
//	n |= n >> 8
//	n |= n >> 16
//	n++
//	return n
//}
