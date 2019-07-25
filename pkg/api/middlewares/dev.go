package middlewares

import "net/http"

// DevMode sets headers required for local testing
func DevMode(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		if r.Method == http.MethodOptions {
			w.Write([]byte(""))
			return
		}
		next.ServeHTTP(w, r)
	})
}
