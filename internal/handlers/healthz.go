package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")  // set the content type to json
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))  // write the response body
}
