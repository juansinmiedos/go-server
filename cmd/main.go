package main

import (
	"net/http"
)

func main() {

	err := http.ListenAndServe(":3000")
	if err != nil {
		//
	}
}