package state

import . "luago/api"

func (s *luaState) Compare(idx1, idx2 int, op CompareOP) bool {
	a := s.stack.get(idx1)
	b := s.stack.get(idx2)

	switch op {
	case LUA_OPEQ:
		return _eq(a, b)
	case LUA_OPLT:
		return _lt(a, b)
	case LUA_OPLE:
		return _le(a, b)
	default:
		panic("invalid compare op!")
	}
}

func _eq(a, b luaValue) bool {
	switch x := a.(type) {
	case nil:
		return b == nil
	case bool:
		y, ok := b.(bool)
		return ok && x == y
	case string:
		y, ok := b.(string)
		return ok && x == y
	case int64:
		switch y := b.(type) {
		case int64:
			return x == y
		case float64:
			return float64(x) == y
		default:
			return false
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x == float64(y)
		case float64:
			return x == y
		default:
			panic("compare type unmatch!")
		}
	default:
		return a == b
	}
}

func _lt(a, b luaValue) bool {
	switch x := a.(type) {
	case string:
		y, ok := b.(string)
		return ok && x > y
	case int64:
		switch y := b.(type) {
		case int64:
			return x > y
		case float64:
			return float64(x) > y
		default:
			panic("compare type unmatch!")
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x > float64(y)
		case float64:
			return x > y
		default:
			panic("compare type unmatch!")
		}
	default:
		panic("_lt compare type unmatch!")
	}
}

func _le(a, b luaValue) bool {

	switch x := a.(type) {
	case string:
		y, ok := b.(string)
		return ok && x < y
	case int64:
		switch y := b.(type) {
		case int64:
			return x < y
		case float64:
			return float64(x) < y
		default:
			return false
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x < float64(y)
		case float64:
			return x < y
		default:
			panic("compare type unmatch!")
		}
	default:
		panic("_lt compare type unmatch!")
	}

}
