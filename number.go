package g

import "strconv"

func ForceInt(s string, defaultInt int) int {
	parsed, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultInt
	}
	return int(parsed)
}

func ForceInt32(s string, defaultInt32 int32) int32 {
	parsed, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return defaultInt32
	}
	return int32(parsed)
}

func ForceInt64(s string, defaultInt64 int64) int64 {
	parsed, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultInt64
	}
	return int64(parsed)
}
