package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"bakersfieldtechnology.com/components"
)

func main() {

	component := components.Hello("Tyler")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3005")
	http.ListenAndServe(":3005", nil)
}
