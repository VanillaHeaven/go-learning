package state

func (s *luaState) Len(idx int) {
	val := s.stack.get(idx)

	switch x := val.(type) {
	case nil:
		s.stack.push(int64(0))
	case string:
		s.stack.push(int64(len(x)))
	default:
		panic("length error!")
	}
}

func (s *luaState) Concat(n int) {
	if n == 0 {
		s.stack.push("")
	} else if n >= 2 {
		if s.stack.top < n {
			panic("exceeds the top of the stack")
		}

		for i := 1; i < n; i++ {
			if s.IsString(-1) && s.IsString(-2) {
				s2 := s.ToString(-1)
				s1 := s.ToString(-2)
				s.stack.pop()
				s.stack.pop()
				s.stack.push(s1 + s2)
				continue
			}
			panic("concat error!")
		}
	}
}
