package middlewares

import (
	"Testgorillamux/util"
	"encoding/json"
	"net/http"
)

func JwtVerify(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("jwt")
		if cookie == nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}
		// header := r.Header.Get("JWT")
		roleId, _ := util.ParseJwt(cookie.Value)

		if roleId != "1" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LoginVerify(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idCookie, _ := r.Cookie("userId")
		if idCookie == nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}
		next.ServeHTTP(w, r)
	})
}
