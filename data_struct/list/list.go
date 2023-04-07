package list

// cdList 插入删除列表元素
func cdList() {
	list := make([]int, 0)
	list = nil

	list = append(list, 1)
	list = append(list, 5)
	list = append(list, 2)
	list = append(list, 8)

	// 中间插入元素, 在索引 3 处插入数字 6
	list = append(list[:3], append([]int{6}, list[3:]...)...)

	// 删除元素,  删除索引 3 处的元素
	list = append(list[:3], list[4:]...)

	// 拼接两个列表
	list1 := []int{2, 3, 5, 4}
	list = append(list, list1...)
}

// MyList 列表类简易实现
type MyList struct {
	lCaps       int   // 容量
	lNums       []int // 数据
	lSize       int   // 当前元素数量
	lExpandSize int   //扩容倍数
}

func NewMyList() *MyList {
	return &MyList{
		lCaps:       10,              // 列表容量
		lNums:       make([]int, 10), // 数组（存储列表元素）
		lSize:       0,               // 列表长度（即当前元素数量）
		lExpandSize: 2,               // 每次列表扩容的倍数
	}
}

// size 获取列表长度（即当前元素数量）
func (l *MyList) size() int {
	return l.lSize
}

// cap 获取列表容量
func (l *MyList) cap() int {
	return l.lCaps
}

// expand 列表扩容
func (l *MyList) expand() {
	if l.lSize == l.lCaps {
		l.lNums = append(l.lNums, make([]int, l.lCaps*(l.lExpandSize-1))...)
		l.lCaps = len(l.lNums)
	}
}

// toArray 返回有效长度的列表
func (l *MyList) toArray() []int {
	// 仅转换有效长度范围内的列表元素
	return l.lNums[:l.lSize]
}

// get 访问元素
func (l *MyList) get(idx int) int {
	if idx < 0 || idx >= l.lSize {
		panic("数组越界")
	}

	return l.lNums[idx]
}

// set 更新元素
func (l *MyList) set(num, idx int) {
	if idx < 0 || idx >= l.lSize {
		panic("数组越界")
	}

	l.lNums[idx] = num
}

// add 尾部添加元素
func (l *MyList) add(num int) {
	l.expand()

	l.lNums[l.lSize] = num
	l.lSize++
}

// insert 中间插入元素
func (l *MyList) insert(num, idx int) {
	if idx < 0 || idx >= l.lSize {
		panic("数组越界")
	}

	l.expand()

	for i := l.lSize; i > idx; i-- {
		l.lNums[i] = l.lNums[i-1]
	}
	l.lNums[idx] = num
	l.lSize++
}

// remove 删除元素
func (l *MyList) remove(idx int) int {
	if idx < 0 || idx >= l.lSize {
		panic("数组越界")
	}

	val := l.lNums[idx]

	for i := idx; i < l.lSize-1; i++ {
		l.lNums[i] = l.lNums[i+1]
	}
	l.lSize--

	return val
}
