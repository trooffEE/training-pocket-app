package server

import (
	"encoding/json"
	"net/http"

	"github.com/trooffEE/training-app/cmd/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) MountRoutes() http.Handler {
	r := chi.NewRouter()
	// csrfKey := mustGenerateCSRFKey()
	// csrfMiddleware := csrf.Protect(csrfKey, csrf.TrustedOrigins([]string{"localhost:8080"}), csrf.FieldName("_csrf"))

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)

	r.Get("/health", s.healthHandler)
	r.Post("/api/auth", web.AuthWebHandler)
	r.Post("/web/components/auth/register", web.RegisterWebHandler)
	r.Post("/web/components/auth/login", web.LoginWebHandler)

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
