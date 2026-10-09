package main

import ( "net/http"
	      "log" 
		  _ "net/http/pprof"
		  "time"

	"OLX-API/internal/config"
	"OLX-API/internal/handlers"
	"OLX-API/internal/db"
)



func main() {

	cfg:=config.MustLoad()

	db, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	mux := http.NewServeMux()  //mux is route handler	

	mux.HandleFunc("GET /healthz", handlers.healthz)

	srv := http.Server{  //server config
		Addr:         ":"+ cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	// start the server
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err) // log the error
	}
	
}