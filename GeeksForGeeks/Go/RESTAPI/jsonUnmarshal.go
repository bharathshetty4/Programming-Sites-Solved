package main

import (
	"encoding/json"
	"fmt"
)

type Items struct {
	Vegetables map[string]int `json:"vegetables"`
	Fruits     Seasonal       `json:"fruits"`
}

type Seasonal struct {
	Fruits map[string]int `json:"seasonal"`
}

func main() {
	var items *Items
	data := []byte("{\"vegetables\":{\"carrot\":5,\"beans\":4},\"fruits\":{\"seasonal\":{\"apple\":3,\"banana\":2}}}")
	json.Unmarshal(data, &items)

	for key, value := range items.Vegetables {
		fmt.Printf("%s: %d\n", key, value)
	}

	for key, value := range items.Fruits.Fruits {
		fmt.Printf("%s: %d\n", key, value)
	}
}
