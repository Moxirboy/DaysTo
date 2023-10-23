package main

import (
	"awesomeProject4/internal"
	"net/http"
)

func main() {
	internal.Router()
	http.ListenAndServe(":8080", nil)
}
