package number

import "strconv"

func ParseInteger(s string) (int64, bool) {
	r, err := strconv.ParseInt(s, 10, 64)
	return r, err == nil
}

func ParseFloat(s string) (float64, bool) {
	r, err := strconv.ParseFloat(s, 64)
	return r, err == nil
}
