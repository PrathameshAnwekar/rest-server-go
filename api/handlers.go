package api

import (
	"fmt"
	"net/http"
)

// Hello is a placeholder function.
func Hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
