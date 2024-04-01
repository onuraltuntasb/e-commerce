package middleware

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/onuraltuntasb/e-commerce/internal/auth"
	"github.com/onuraltuntasb/e-commerce/internal/utils"
)

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", " http://127.0.0.1:5173")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
			return

		} else {
			next.ServeHTTP(w, r)
		}
	})
}		

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


		_, _, err := auth.GetTokenFromHeaderAndVerify(w, r)
		if err != nil {
			fmt.Printf("\n middleware error : %v \n", err)
			w.WriteHeader(http.StatusUnauthorized)
			utils.WriteJSON(w, http.StatusUnauthorized, "status unauthorized")
			return
		}
 
		next.ServeHTTP(w, r)
	})
}
