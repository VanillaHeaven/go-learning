package state

import (
	. "luago/api"
	"luago/number"
)

/*
 * 定义了对栈值的操作函数
 * 识别栈值类型，
 * 定义栈值类型的强制转换函数
 */
type luaValue interface{}

func typeOf(val luaValue) LuaType {
	switch val.(type) {
	case nil:
		return LUA_TNIL
	case bool:
		return LUA_TBOOLEAN
	case int64:
		return LUA_TNUMBER
	case float64:
		return LUA_TNUMBER
	case string:
		return LUA_TSTRING
	default:
		panic("todo!")
	}
}

func convertToBoolean(val luaValue) bool {
	switch x := val.(type) {
	case nil:
		return false
	case bool:
		return x
	default:
		return true
	}
}

func convertToFloat(val luaValue) (float64, bool) {
	switch x := val.(type) {
	case int64:
		return float64(x), true
	case float64:
		return x, true
	case string:
		return number.ParseFloat(x)
	default:
		return 0, false
	}
}

func convertInteger(val luaValue) (int64, bool) {
	switch x := val.(type) {
	case int64:
		return x, true
	case float64:
		return int64(x), true
	case string:
		return number.ParseInteger(x)
	default:
		return 0, false
	}
}

func _stringToInteger(s string) (int64, bool) {
	if r, ok := number.ParseInteger(s); ok {
		return r, ok
	} else if r, ok := number.ParseFloat(s); ok {
		return number.FloatToInteger(r)
	}
	return 0, false
}
