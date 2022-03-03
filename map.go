package g

// TODO: After Go 1.18, uses generic
func MergeMap(m1 map[string]string, m2 map[string]string) map[string]string {
	var merged = make(map[string]string)
	for k, v := range m1 {
		merged[k] = v
	}
	for k, v := range m2 {
		merged[k] = v
	}
	return merged
}
