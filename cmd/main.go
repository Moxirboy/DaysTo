package main

import (
	"DaysTo/internal"
	"net/http"
)

func main() {
	internal.Router()
	http.ListenAndServe(":8080", nil)
}
