package g

import "golang.org/x/exp/constraints"

func MergeMap[K constraints.Ordered, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	var merged = make(map[K]V)
	for k, v := range m1 {
		merged[k] = v
	}
	for k, v := range m2 {
		merged[k] = v
	}
	return merged
}
