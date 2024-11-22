package controllers

import (
	"PLP_Backend/classes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ConsultaMissaoHeroi(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		NomeHeroi string `json:"nome_heroi"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	nomeHeroi := requestData.NomeHeroi

	// Configura o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	missao, err := classes.ConsultaMissoesPorHeroi(nomeHeroi)
	if err != nil {
		http.Error(w, "Missão não encontrada ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(missao)
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}

func ListarTodasMissoesHandler(w http.ResponseWriter, r *http.Request) {
	// Adicionar headers CORS específicos
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Se for uma requisição OPTIONS, retornar imediatamente
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	log.Println("Recebida requisição para listar missões") // Log para debug

	missoes, err := classes.ListarTodasMissoes()
	if err != nil {
		log.Printf("Erro ao listar missões: %v", err) // Log do erro
		http.Error(w, "Erro ao listar missões", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	log.Printf("Retornando %d missões", len(missoes)) // Log da quantidade de missões

	if err := json.NewEncoder(w).Encode(missoes); err != nil {
		log.Printf("Erro ao codificar resposta: %v", err) // Log de erro de codificação
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}

func ConsultaMissaoPorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idMissao := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	missao, err := classes.ConsultaMissaoPorId(idMissao)
	if err != nil {
		http.Error(w, "Missão não encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(missao)
}

func AtualizarMissao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idMissao := vars["id"]

	var missao classes.Missoes
	err := json.NewDecoder(r.Body).Decode(&missao)
	if err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		return
	}

	err = classes.AtualizarMissao(idMissao, missao)
	if err != nil {
		http.Error(w, "Erro ao atualizar missão", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeletarMissaoHandler(w http.ResponseWriter, r *http.Request) {
	// Configurar headers CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Se for uma requisição OPTIONS, retornar imediatamente
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Pegar o ID da missão dos parâmetros da URL
	vars := mux.Vars(r)
	idMissao := vars["id"]

	err := classes.DeletarMissao(idMissao)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao deletar missão: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Missão deletada com sucesso"})
}
