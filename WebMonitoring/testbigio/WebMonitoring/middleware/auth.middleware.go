package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func JwtVerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			ResponseWithError(w, http.StatusBadRequest, "Authorization required")
			return
		}
		baererToken := strings.Split(authHeader, " ")
		if len(baererToken) != 2 {
			ResponseWithError(w, http.StatusBadRequest, "Authorization required")
			return
		}
		token, err := jwt.Parse(baererToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Something Wrong")
			}
			return []byte("I love bee"), nil
		})
		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Println("Claims")
		if !ok && !token.Valid {
			fmt.Println(err)
			ResponseWithError(w, http.StatusBadRequest, "Unauthorization")
		}
		userId := strconv.FormatFloat(claims["id"].(float64), 'g', 1, 64)
		r.Header.Set("user_id", userId)

	})
}

func ResponseWithError(w http.ResponseWriter, code int, message interface{}) {
	payload := map[string]interface{}{"error": message}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
