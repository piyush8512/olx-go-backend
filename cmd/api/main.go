package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"OLX-API/internal/config"
	"OLX-API/internal/db"
	"OLX-API/internal/handlers"
)

func main() {
	cfg := config.MustLoad()

	database, err := db.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}
	defer database.Close()

	lh := handlers.NewListingHandler(database)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listings/{id}", lh.Delete)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
