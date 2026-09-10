package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type horarioAtual struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Número total de requisições HTTP recebidas por rota e método",
	},
	[]string{"path", "method"},
)

var httpRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duração das requisições HTTP em segundos",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"path", "method"},
)

func projetoKorp(w http.ResponseWriter, req *http.Request) {
	//Timer para o httpRequestDuration
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		httpRequestDuration.WithLabelValues("/projeto-korp", req.Method).Observe(duration)
	}()

	//Contador de request
	httpRequestsTotal.WithLabelValues("/projeto-korp", req.Method).Inc()

	w.Header().Set("Content-Type", "application/json")
	response := horarioAtual{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(response)
	log.Printf("/projeto-korpRequest recebido %s", response.Horario)
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	httpRequestsTotal.WithLabelValues("/health", req.Method).Inc()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		httpRequestsTotal,
		httpRequestDuration,
	)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/projeto-korp", projetoKorp)

	log.Println("Registered routes:")
	log.Println("- /metrics")
	log.Println("- /health")
	log.Println("- /projeto-korp")
	log.Println("Servidor iniciado na porta :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
