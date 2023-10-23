package middleware

import (
	_const "awesomeProject4/internal/const"
	"fmt"
	"net/http"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")

		if token == "" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, _const.NotAuth)
			return
		}

		next(w, r)
	}
}
