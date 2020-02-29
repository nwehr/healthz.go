package healthz

import (
	"net/http"
)

// Start sets up /healthz/alive and /healthz/ready endpoints
func Start(listen string, ready *bool) {
	go func() {
		mux := http.NewServeMux()

		mux.HandleFunc("/healthz/alive", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		mux.HandleFunc("/healthz/ready", func(w http.ResponseWriter, r *http.Request) {
			if !*ready {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		http.ListenAndServe(listen, mux)
	}()
}
