package middleware

import "net/http"

func APIKeyAuth(validKeys map[string]struct{}) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			key := r.Header.Get("x-api-Key")

			if key == "" {
				http.Error(w, "unauthenticated", http.StatusUnauthorized)
				return
			}

			if _, ok := validKeys[key]; !ok {
				http.Error(w, "unauthenticated", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}