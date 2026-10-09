package main

import ( "net/http"
	      "log" 
		  _ "net/http/pprof"
		  "time"
)



func main() {
	mux := http.NewServeMux()  //mux is route handler	

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // set the content type to json
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`)) // write the response body
	})

	srv := http.Server{  //server config
		Addr:         ":8090" ,
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