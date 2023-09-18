package middleware

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Dummy User
		userList := []User{
			{
				Username: "bob",
				Password: "@password123",
			},
			{
				Username: "alice",
				Password: "@password123",
			},
			{
				Username: "eve",
				Password: "@password123",
			},
		}

		// Get Basic Auth Value
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			m := Response{
				Code:    http.StatusUnauthorized,
				Message: "wrong password or username",
			}
			d, _ := json.Marshal(m)
			w.Write(d)
			return
		}

		// Check if username and password correct
		for _, v := range userList {
			if v.Username == username {
				if v.Password != password {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusUnauthorized)
					m := Response{
						Code:    http.StatusUnauthorized,
						Message: "wrong password or username",
					}
					d, _ := json.Marshal(m)
					w.Write(d)
					return
				}
			}
		}

		handler.ServeHTTP(w, r)
	})
}
