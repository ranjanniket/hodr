package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define Prometheus counters for request metrics
var (
	requestMethodCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)
	responseStatusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_responses_total",
			Help: "Total number of HTTP responses",
		},
		[]string{"status_code", "method", "path"},
	)
)

var log *slog.Logger

func init() {
	log = slog.New(slog.NewTextHandler(os.Stdout, nil))
	// Register Prometheus metrics
	prometheus.MustRegister(requestMethodCounter)
	prometheus.MustRegister(responseStatusCounter)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Increment request method and path counter
	requestMethodCounter.WithLabelValues(r.Method, r.URL.Path).Inc()

	// Set the response header
	w.Header().Set("Content-Type", "text/plain")

	log.Info("request", "method", r.Method, "path", r.RequestURI)
	// Check the request method
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// Respond with "Hodor" for both GET and POST requests
		fmt.Fprint(w, "hodor... hodor... hodor")
	default:
		// Respond with a 405 Method Not Allowed for other request methods
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	responseStatusCounter.WithLabelValues(fmt.Sprint(http.StatusOK), r.Method, r.URL.Path).Inc()
}

func main() {
	// Register the handler function for all routes
	http.HandleFunc("/", handler)
	// Register Prometheus endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Start the server on port 8080
	log.Info("Server is running on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
