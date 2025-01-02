package main

import (
	"os"
	"net/http"
	"log/slog"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	infraRepo "prtl-base-api/internal/infrastructure/repository"
	infraJWT "prtl-base-api/internal/infrastructure/jwt"
	appService "prtl-base-api/internal/application/service"
	handler "prtl-base-api/internal/presentation/handler"
)

func main() {
    err := godotenv.Load()
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
		JWTProvider := infraJWT.NewJWTProvider(
			os.Getenv("JWT_PRIVATE_KEY"),
			os.Getenv("JWT_PUBLIC_KEY"),
		)
		AuthService := appService.NewAuthService(*JWTProvider, UserRepo)
		AuthHandler := handler.NewAuthHandler(AuthService)

		api.Route("/login", func(health chi.Router) {
			health.Post("/", AuthHandler.LoginUser)
		})
	})

	http.ListenAndServe(":80", r)
}
