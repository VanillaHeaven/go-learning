package number

import "math"

func IFloorDiv(a, b int64) int64 {
	if (a > 0 && b > 0) || (a < 0 && b < 0) || a%b == 0 {
		return a / b
	}
	return (a / b) - 1
}

func FFloorDiv(a, b float64) float64 {
	return math.Floor(a / b)
}

func IMod(a, b int64) int64 {
	return a - b*IFloorDiv(a, b)
}

func FMod(a, b float64) float64 {
	return a - b*FFloorDiv(a, b)
}

func ShiftRight(a, n int64) int64 {
	if n >= 0 {
		return int64(uint64(a) >> uint64(n))
	}
	return ShiftLeft(a, -n)
}

func ShiftLeft(a, n int64) int64 {
	if n >= 0 {
		return a << uint64(n)
	}
	return ShiftRight(a, -n)
}

func FloatToInteger(f float64) (int64, bool) {
	r := int64(f)
	return r, float64(r) == f
}
