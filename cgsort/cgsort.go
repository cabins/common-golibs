package cgsort

import (
	"sort"
)

type result struct {
	Key   string
	Value interface{}
}

func SortMapByKey(m map[string]int) []result {
	// Transfer m to slice
	var results []result
	for key, value := range m {
		results = append(results, result{key, value})
	}

	// sort the slice with sort.Slice method
	sort.Slice(results, func(i, j int) bool {
		return results[i].Key < results[j].Key
	})

	return results
}

func SortMapByIntValue(m map[string]int) []result {
	var results []result
	for key, value := range m {
		results = append(results, result{key, value})
	}

	// sort the slice with sort.Slice method
	sort.Slice(results, func(i, j int) bool {
		return results[i].Key < results[j].Key
	})

	return results
}
