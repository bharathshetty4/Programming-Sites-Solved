package main

import "fmt"

func main() {
	keys, vals := getKeys(map[any]any{1: "one", "two": 2})
	fmt.Printf("keys: %v, values: %v", keys, vals)
}

func getKeys(m map[any]any) ([]any, []any) {
	var keys []any
	var vals []any
	for k, v := range m {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	return keys, vals
}
