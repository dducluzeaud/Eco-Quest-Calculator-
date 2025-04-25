package main

import (
	"eco-quest-calculator/backend/models"
	"eco-quest-calculator/backend/routes"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize database
	if err := models.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create new router
	mux := http.NewServeMux()

	// Register routes
	routes.RegisterAuthRoutes(mux)

	// Add middleware
	handler := corsMiddleware(mux)

	// Create server
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Start server
	log.Printf("Server starting on http://localhost%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
