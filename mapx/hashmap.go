package mapx

import "ekit/syncx"

/*
	扩展版HashMap实现：支持泛型、自定义哈希和相等比较、使用链表法解决哈希冲突的哈希表
核心功能：
（1）创建：初始化HashMap，指定初始容量(size)，并创建节点对象池
（2）插入/更新：插入键值对。使用 key.Code() 计算哈希，通过 key.Equals() 判断键相等，支持更新已有键的值
（3）查询：根据键查找值，返回值和是否存在标志
（4）获取集合：
	keys()：返回所有键的切片（随机顺序）
	Values()：返回所有值的切片（随机顺序）
（5）遍历：随机顺序遍历所有键值对，若回调函数返回false则停止遍历
（6）获取大小：返回哈希表中桶的数量（非键值对总数）
*/

type node[T Hashable, ValType any] struct {
	key   T                 // 键
	value ValType           // 值
	next  *node[T, ValType] // 指向下一个节点的指针（用于解决哈希冲突）
}

func (m *HashMap[T, ValType]) newNode(key T, val ValType) *node[T, ValType] {
	newNode := m.nodePool.Get()
	newNode.value = val
	newNode.key = key
	return newNode
}

type Hashable interface {
	// Code 返回该元素的哈希值
	// 注意：哈希值应该尽可能的均匀以避免冲突
	Code() uint64
	// Equals 比较两个元素是否相等。如果返回 true，那么我们会认为两个键是一样的。
	Equals(key any) bool
}

type HashMap[T Hashable, ValType any] struct {
	hashmap  map[uint64]*node[T, ValType]   // 哈希表：key是哈希值，value是链表头节点
	nodePool *syncx.Pool[*node[T, ValType]] // 对象池：复用节点，减少内存分配
}

// Put 插入/更新键值对
func (m *HashMap[T, ValType]) Put(key T, val ValType) error {
	// 1.计算哈希
	hash := key.Code()
	// 2.查找桶
	root, ok := m.hashmap[hash]
	if !ok {
		// 3.桶为空，创建新节点并作为链表头
		hash = key.Code()
		newNode := m.newNode(key, val)
		m.hashmap[hash] = newNode
		return nil
	}
	// 4.桶不为空
	pre := root
	for root != nil {
		// 5.遍历链表查找相同 key
		if root.key.Equals(key) {
			// 6.找到相同 key，更新值
			root.value = val
			return nil
		}
		// 7.没有找到相同 key，继续遍历
		pre = root
		root = root.next
	}
	// 8.没有找到相同 key，创建新节点并插入到链表末尾
	newNode := m.newNode(key, val)
	pre.next = newNode // 将新节点插入到链表末尾
	return nil
}

// Get - 获取值
func (m *HashMap[T, ValType]) Get(key T) (ValType, bool) {
	// 1.计算哈希，定位桶
	hash := key.Code()
	root, ok := m.hashmap[hash]
	var val ValType
	if !ok {
		// 2.如果桶为空，返回 (零值, false)
		return val, false
	}
	// 3.遍历链表，用 Equals 比较 key
	for root != nil {
		if root.key.Equals(key) {
			// 4.找到返回 (value, true)
			return root.value, true
		}
		root = root.next
	}
	// 5.找不到，返回 (零值, false)
	return val, false
}

// Keys 返回 Hashmap 里面的所有的 key。
// 注意：key 的顺序是随机的。
func (m *HashMap[T, ValType]) Keys() []T {
	res := make([]T, 0)
	// 1.遍历所有桶
	for _, bucketNode := range m.hashmap {
		curNode := bucketNode
		// 2.遍历每个桶的链表
		for curNode != nil {
			// 3.将 key 添加到结果切片
			res = append(res, curNode.key)
			// 4.继续遍历下一个节点
			curNode = curNode.next
		}
	}
	return res
}

// Values 返回 Hashmap 里面的所有的 value。
// 注意：value 的顺序是随机的。
func (m *HashMap[T, ValType]) Values() []ValType {
	res := make([]ValType, 0)
	// 1.遍历所有桶
	for _, bucketNode := range m.hashmap {
		curNode := bucketNode
		// 2.遍历每个桶的链表
		for curNode != nil {
			// 3.将 value 添加到结果切片
			res = append(res, curNode.value)
			// 4.继续遍历下一个节点
			curNode = curNode.next
		}
	}
	return res
}

func NewHashMap[T Hashable, ValType any](size int) *HashMap[T, ValType] {
	return &HashMap[T, ValType]{
		nodePool: syncx.NewPool[*node[T, ValType]](func() *node[T, ValType] {
			return &node[T, ValType]{}
		}),
		hashmap: make(map[uint64]*node[T, ValType], size),
	}
}

var _ mapi[Hashable, any] = (*HashMap[Hashable, any])(nil)

// Delete 第一个返回值为删除key的值，第二个是hashmap是否真的有这个key
func (m *HashMap[T, ValType]) Delete(key T) (ValType, bool) {
	// 1.计算哈希，定位桶
	root, ok := m.hashmap[key.Code()]
	if !ok {
		// 2.如果桶为空，返回 (零值, false)
		var t ValType
		return t, false
	}

	pre := root
	num := 0
	for root != nil {
		// 3.遍历链表查找 key，维护 pre 前驱
		if root.key.Equals(key) {
			// 4.找到 key，删除节点
			if num == 0 && root.next == nil {
				// 4.1 如果是链表头且无后继：从 hashmap 删除该桶
				delete(m.hashmap, key.Code())
			} else if num == 0 && root.next != nil {
				// 4.2 如果是链表头且有后继：更新桶头指针
				m.hashmap[key.Code()] = root.next
			} else {
				// 4.3 不是链表头：更新前驱 next 指针
				pre.next = root.next
			}
			// 5.清理节点：调用 formatting() 重置，放回对象池
			val := root.value
			root.formatting()
			m.nodePool.Put(root)
			return val, true
		}
		// 6.继续遍历下一个节点
		num++
		pre = root
		root = root.next
	}
	// 7.没有找到 key，返回 (零值, false)
	var t ValType
	return t, false
}

func (n *node[T, ValType]) formatting() {
	var val ValType
	var t T
	n.key = t
	n.value = val
	n.next = nil
}

// Len 获取桶的数量
func (m *HashMap[T, ValType]) Len() int64 {
	return int64(len(m.hashmap))
}

// Iterate 随机顺序遍历，并对每个键值对执行cb(k, v)
// 如果cb的返回值为 true 则继续遍历，否则遍历结束
func (m *HashMap[T, ValType]) Iterate(cb func(key T, val ValType) bool) {
	// 1.遍历所有桶
	for _, nodeHead := range m.hashmap {
		cur := nodeHead
		// 2.遍历每个桶的链表
		for ; cur != nil; cur = cur.next {
			// 3.执行回调函数
			if !cb(cur.key, cur.value) {
				// 4.回调函数返回 false，遍历结束
				return
			}
		}
	}
}
