package internal

import (
	"DaysTo/internal/handle"
	"DaysTo/internal/middleware"
	"net/http"
)

func Router() {
	http.HandleFunc("/", middleware.Middleware(handle.DaysTo))
}
