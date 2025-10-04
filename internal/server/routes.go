package server

import (
	"context"
	"net/http"
	"storeback/internal/handlers"
)

func addRoutes(ctx context.Context, server *http.ServeMux) {
	server.HandleFunc("GET /products/all", handlers.ListAllProducts(ctx))
}
