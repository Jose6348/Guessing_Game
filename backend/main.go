package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

// Middleware para permitir CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite qualquer origem
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

var numeroSorteado = rand.IntN(101)
var tentativas []int

type Chute struct {
	Numero int `json:"numero"`
}

type Resposta struct {
	Mensagem   string `json:"mensagem"`
	Acertou    bool   `json:"acertou"`
	Tentativas int    `json:"tentativas"`
	Historico  []int  `json:"historico"`
}

func verificarChute(w http.ResponseWriter, r *http.Request) {
	var chute Chute

	err := json.NewDecoder(r.Body).Decode(&chute)
	if err != nil {
		http.Error(w, "Erro ao processar JSON", http.StatusBadRequest)
		return
	}

	tentativas = append(tentativas, chute.Numero)

	var resposta Resposta
	switch {
	case chute.Numero < numeroSorteado:
		resposta = Resposta{"O número é maior!", false, len(tentativas), tentativas}
	case chute.Numero > numeroSorteado:
		resposta = Resposta{"O número é menor!", false, len(tentativas), tentativas}
	default:
		resposta = Resposta{fmt.Sprintf("Parabéns! Você acertou em %d tentativas!", len(tentativas)), true, len(tentativas), tentativas}
		numeroSorteado = rand.IntN(101) // Novo número após acerto
		tentativas = nil
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/chute", verificarChute)

	// Aplicando o middleware de CORS
	handler := enableCORS(mux)

	fmt.Println("🚀 Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
