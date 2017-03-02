package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var counter = 0

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {

			output := map[string]interface{}{}

			// Add data fields
			output["hello"] = "world"
			output["elasticon"] = "2017"
			output["counter"] = counter

			// convert to json
			bytes, err := json.Marshal(output)
			if err != nil {
				fmt.Fprintln(w, err)
			}

			fmt.Fprintln(w, string(bytes))
			counter = counter + 1
		},
	)
	http.ListenAndServe(":8080", nil)
}
