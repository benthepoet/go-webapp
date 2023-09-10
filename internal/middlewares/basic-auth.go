package middlewares

import "net/http"

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			basicAuthFailed(w)
			return
		}

		if u != "ben" || p != "hello" {
			basicAuthFailed(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func basicAuthFailed(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", "Basic realm=protected")
	w.WriteHeader(http.StatusUnauthorized)
}
