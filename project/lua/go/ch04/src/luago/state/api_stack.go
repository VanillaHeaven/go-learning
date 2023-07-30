package state

func (s *luaState) GetTop() int {
	return s.stack.top
}

func (s *luaState) AbsIndex(idx int) int {
	return s.stack.absIndex(idx)
}

func (s *luaState) CheckStack(n int) bool {
	s.stack.check(n)
	return true
}

func (s *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		s.stack.pop()
	}
}

func (s *luaState) Copy(from, to int) {
	if !s.stack.isValid(from) || !s.stack.isValid(to) {
		panic("Copy invalid index!")
	}
	val := s.stack.get(from)
	s.stack.set(to, val)
}

func (s *luaState) PushValue(idx int) {
	if !s.stack.isValid(idx) {
		panic("PushValue invalid index!")
	}
	val := s.stack.get(idx)
	s.stack.push(val)
}

func (s *luaState) Replace(idx int) {
	if !s.stack.isValid(idx) {
		panic("Replace invalid index!")
	}
	val := s.stack.pop()
	s.stack.set(idx, val)
}

func (s *luaState) Insert(idx int) {
	if !s.stack.isValid(idx) {
		panic("Insert invalid index!")
	}
	s.Rotate(idx, 1)
}

func (s *luaState) Remove(idx int) {
	if !s.stack.isValid(idx) {
		panic("Remove invalid index!")
	}
	s.Rotate(idx, -1)
	s.Pop(1)
}

func (s *luaState) Rotate(idx int, n int) {
	if !s.stack.isValid(idx) {
		panic("Rotate invalid index!")
	}
	t := s.stack.top - 1
	p := s.stack.absIndex(idx) - 1
	var m int
	if n > 0 {
		m = t - n
	} else {
		m = p - n - 1
	}

	s.stack.reverse(p, m)
	s.stack.reverse(m+1, t)
	s.stack.reverse(p, t)
}

func (s *luaState) SetTop(idx int) {
	absIdx := s.stack.absIndex(idx)
	if absIdx < 0 {
		panic("SetTop invalid index!")
	}
	for s.stack.top > absIdx {
		s.stack.pop()
	}
	for s.stack.top < absIdx {
		s.stack.push(nil)
	}
}
