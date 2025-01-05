package main

import (
	"os"
	"net/http"
	"log/slog"
	"github.com/joho/godotenv"
	"github.com/golang-jwt/jwt/v5"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/cors"

	infraRepo "prtl-base-api/internal/infrastructure/repository"
	appService "prtl-base-api/internal/application/service"
	handler "prtl-base-api/internal/presentation/handler"
)

var tokenAuth *jwtauth.JWTAuth

func main() {
    err := godotenv.Load()

    privateKey := os.Getenv("JWT_PRIVATE_KEY")
    publicKey := os.Getenv("JWT_PUBLIC_KEY")
    parsedPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
    if err != nil {
		slog.Error(err.Error())
	}
    parsedPublicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
    if err != nil {
		slog.Error(err.Error())
	}
    tokenAuth = jwtauth.New("RS256", parsedPrivateKey, parsedPublicKey)

    if err != nil {
    	slog.Error(err.Error())
    }
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
			health.Get("/", handler.HealthCheck)
		})

		UserRepo := infraRepo.NewUserRepositoryImpl()
		AuthService := appService.NewAuthService(tokenAuth, UserRepo)
		AuthHandler := handler.NewAuthHandler(AuthService)

		api.Route("/login", func(health chi.Router) {
			health.Post("/", AuthHandler.LoginUser)
		})
		api.Route("/protected", func(health chi.Router) {
			health.Use(jwtauth.Verifier(tokenAuth))
			health.Use(jwtauth.Authenticator(tokenAuth))
			health.Get("/", handler.HealthCheck)
		})

	})

	http.ListenAndServe(":80", r)
}
