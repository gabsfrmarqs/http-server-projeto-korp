package main

import (
    "net/http"
	"encoding/json"
	"time"
	"log"
)

type horarioAtual struct {
    Nome   string `json:"nome"`
    Horario string `json:"horario"`
}

func projetoKorp(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json") 
	response := horarioAtual{
		Nome: "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}
	
    json.NewEncoder(w).Encode(response)
	log.Printf("/projeto-korpRequest recebido %s", response.Horario)
}

func main() {
    http.HandleFunc("/projeto-korp", projetoKorp)
    http.ListenAndServe(":8080", nil)
	log.Printf("Servidor iniciado")
	//Log basiquinho
}