package state

/*
 * 定义了 lua 底层的栈结构,
 * 提供了基础的栈操作功能 push, pop
 * lua 的栈允许对非栈顶元素进行读写操作 get, set
 */
type luaStack struct {
	slots []luaValue
	top   int
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

func (ls *luaStack) check(n int) {
	free := len(ls.slots) - ls.top

	for i := free; i < n; i++ {
		ls.slots = append(ls.slots, nil)
	}
}

func (ls *luaStack) push(val luaValue) {
	if ls.top == len(ls.slots) {
		panic("stack overflow!")
	}
	ls.slots[ls.top] = val
	ls.top++
}

func (ls *luaStack) pop() luaValue {
	if ls.top == 0 {
		panic("stack underflow!")
	}
	ls.top--
	val := ls.slots[ls.top]
	ls.slots[ls.top] = nil
	return val
}

// idx is -1, it means get the first one from stack top
func (ls *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}

	return idx + ls.top + 1
}

func (ls *luaStack) isAbsIdxValid(absIdx int) bool {
	return absIdx > 0 && absIdx <= ls.top
}

func (ls *luaStack) isValid(idx int) bool {
	absIdx := ls.absIndex(idx)
	return ls.isAbsIdxValid(absIdx)
}

func (ls *luaStack) get(idx int) luaValue {
	absIdx := ls.absIndex(idx)
	if ls.isAbsIdxValid(absIdx) {
		return ls.slots[absIdx-1]
	}
	return nil
}

func (ls *luaStack) set(idx int, val luaValue) {
	absIdx := ls.absIndex(idx)
	if ls.isAbsIdxValid(absIdx) {
		ls.slots[absIdx-1] = val
		return
	}

	panic("invalid index!")
}

func (ls *luaStack) reverse(from, to int) {
	slots := ls.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
