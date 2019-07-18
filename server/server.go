package main

import (
	"fmt"
	"net/http"
)

// PlayerServer does something...
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
