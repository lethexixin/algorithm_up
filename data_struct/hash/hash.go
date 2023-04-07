package hash

import "fmt"

// imitateHash 模拟hash相关操作
func imitateHash() {
	m := make(map[int]string)

	m[12836] = "小哈"
	m[15937] = "小啰"
	m[16750] = "小算"
	m[13276] = "小法"
	m[10583] = "小鸭"

	fmt.Println(m[15937])

	delete(m, 10583)

	for k, v := range m {
		fmt.Println(k, v)
	}
}

// myHashFunc 自定义hash函数
func (a *ListHash) myHashFunc(key int) (idx int) {
	return key % 100
}

type entry struct {
	key int
	val string
}

// ListHash 基于数组简易实现的哈希表
type ListHash struct {
	buckets []*entry
}

func NewListHash() *ListHash {
	// 初始化数组，包含 100 个桶
	buckets := make([]*entry, 100)
	return &ListHash{buckets: buckets}
}

/* get 查询操作 */
func (a *ListHash) get(key int) string {
	idx := a.myHashFunc(key)
	ety := a.buckets[idx]
	if ety == nil {
		return "Not Found"
	}
	return ety.val
}

/* put 添加操作 */
func (a *ListHash) put(key int, val string) {
	idx := a.myHashFunc(key)
	a.buckets[idx] = &entry{
		key: key,
		val: val,
	}
}

/* remove 删除操作 */
func (a *ListHash) remove(key int) {
	idx := a.myHashFunc(key)
	a.buckets[idx] = nil
}

/* keySet 获取所有键 */
func (a *ListHash) keySet() []int {
	keys := make([]int, 0)
	for _, bucket := range a.buckets {
		if bucket != nil {
			keys = append(keys, bucket.key)
		}
	}
	return keys
}

/* valSet 获取所有值 */
func (a *ListHash) valSet() []string {
	vs := make([]string, 0)
	for _, bucket := range a.buckets {
		if bucket != nil {
			vs = append(vs, bucket.val)
		}
	}
	return vs
}

/* print 打印哈希表 */
func (a *ListHash) print() {
	for _, bucket := range a.buckets {
		if bucket != nil {
			fmt.Println(bucket.key, "->", bucket.val)
		}
	}
}
