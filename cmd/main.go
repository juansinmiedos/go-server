package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/juansinmiedos/go-server/handler"
)

func main() {
	n := negroni.New()
	r := handler.New()
	n.UseHandler(r)

	fmt.Println("Server running in port 3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("Error while serving")
	}
}
