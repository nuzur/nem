package auth

import (
	"crypto/sha256"
	"crypto/subtle"
	"go.uber.org/config"
	"net/http"
)

func BasicAuth(config config.Provider, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		configUsername := config.Get("auth.basic.username").String()
		configPassword := config.Get("auth.basic.password").String()

		if configUsername != "" && configUsername != "<nil>" && configPassword != "" && configPassword != "<nil>" {
			username, password, ok := r.BasicAuth()
			if ok {
				usernameHash := sha256.Sum256([]byte(username))
				passwordHash := sha256.Sum256([]byte(password))
				expectedUsernameHash := sha256.Sum256([]byte(configUsername))
				expectedPasswordHash := sha256.Sum256([]byte(configPassword))

				usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
				passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

				if usernameMatch && passwordMatch {
					next.ServeHTTP(w, r)
					return
				}
			}

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
			return
		}
	})
}
