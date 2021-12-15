package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//unmarshal into interface without defining the struct
	birdJson := `{"birds":{"pigeon":"likes to perch","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)
	birds := result["birds"].(map[string]interface{})
	for key, value := range birds {
		fmt.Println(key, value.(string))
	}
	unmarsahlEx()
}

type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

// unmarshal with the struct
func unmarsahlEx() {
	birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("%+v\n", bird)

}
