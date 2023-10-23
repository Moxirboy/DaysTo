package internal

import (
	"awesomeProject4/internal/handle"
	"awesomeProject4/internal/middleware"
	"net/http"
)

func Router() {
	http.HandleFunc("/", middleware.Middleware(handle.DaysTo))
}
