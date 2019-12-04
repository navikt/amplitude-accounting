package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response struct {
	Event_id         int
	Event_properties struct {
		Application_name string
	}
}

func main() {
	var err error
	var data response
	var counts = make(map[string]int)

	decoder := json.NewDecoder(os.Stdin)
	for err == nil {
		err = decoder.Decode(&data)
		counts[data.Event_properties.Application_name]++
	}

	for key, count := range counts {
		fmt.Println(count, key)
	}
}
