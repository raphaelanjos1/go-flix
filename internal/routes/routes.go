package routes

import "net/http"

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/health-check", healthCheckHandler)
	return router
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
