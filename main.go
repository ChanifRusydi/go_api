package main

import (
    "encoding/json"
    "log"
    "net/http"
	"time"
)

func main() {
    router := NewRouter()
	log.Println("Server started on: http://localhost:3000")
    log.Fatal(http.ListenAndServe(":3000", router))
}
// NewRouter creates and returns a new router
func NewRouter() *http.ServeMux {
    router := http.NewServeMux()
    router.HandleFunc("/time", getCurrentTime)
    return router
}

// getCurrentTime handles requests to the /time endpoint
func getCurrentTime(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now()
    response := map[string]string{
        "current_time": currentTime.Format(time.RFC3339),
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}