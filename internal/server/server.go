package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"storeback/internal/keys"
	"sync"
	"time"
)

func CreateServer(ctx context.Context) (server *http.Server) {
	mux := http.NewServeMux()

	addRoutes(ctx, mux)
	server = &http.Server{
		Addr:    fmt.Sprintf(":%d", ctx.Value(keys.StringKey("PORT")).(int16)),
		Handler: mux,
	}

	return server
}

func RunServer(ctx context.Context, server *http.Server) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		log.Printf("Server running on port: %d\n", ctx.Value(keys.StringKey("PORT")).(int16))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("error running server: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("could not shut down server: %s\n", err)
		} else {
			log.Printf("Server shut down successfully\n")
		}
	}()

	wg.Wait()
}
