package main

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	"encoding/json"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key")

type User struct {
	ID       int
	Username string
	Password string
}

var users = []User{
	{ID: 1, Username: "username", Password: "password"},
}

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
    }))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(api chi.Router) {
		api.Route("/health", func(health chi.Router) {
			health.Get("/", healthCheckHandler)
		})
		api.Route("/login", func(health chi.Router) {
			health.Post("/", loginHandler)
		})
	})

	http.ListenAndServe(":80", r)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var data RequestData
	json.NewDecoder(r.Body).Decode(&data)

	for _, user := range users {
		log.Printf("username: %s", user.Username)
		if user.Username == data.Username && user.Password == data.Password {
			token, err := generateJWT(user)
			if err != nil {
				http.Error(w, "Could not generate token", http.StatusInternalServerError)
				log.Printf("JWT ERROR: %v", err)
				return
			}
			log.Printf("JWT: %s", token)
			writeJSON(w, http.StatusOK, JSONResponse{
				Success: true,
				Message: "Login successful",
				Data:    map[string]string{"token": token},
			})
			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"This is a protected endpoint"}`))
}

func jwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func generateJWT(user User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.ID,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func writeJSON(w http.ResponseWriter, status int, response JSONResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

type RequestData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type JSONResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
