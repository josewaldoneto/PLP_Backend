package controllers

import (
	"PLP_Backend/classes"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller para consultar crimes por heroi e severidade
func ConsultaCrimesHS(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		NomeHeroi        string `json:"nome_heroi"`
		SeveridadeMinima int    `json:"severidade_minima"`
		SeveridadeMaxima int    `json:"severidade_maxima"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	nomeHeroi := requestData.NomeHeroi
	severidadeMinima := requestData.SeveridadeMinima
	severidadeMaxima := requestData.SeveridadeMaxima

	// Configura o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	crimes, err := classes.ConsultaCrimesPorHeroiESeveridade(nomeHeroi, severidadeMinima, severidadeMaxima)
	if err != nil {
		http.Error(w, "Crimes não encontrado ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(crimes)
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}

// Controller para consultar crimes por heroi
func ConsultaCrimesHeroi(w http.ResponseWriter, r *http.Request) {
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

	crimes, err := classes.ConsultaCrimesPorHeroi(nomeHeroi)
	if err != nil {
		http.Error(w, "Crimes não encontrado ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(crimes)
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}

// Controller para consultar crimes de acordo com a severidade
func ConsultaCrimesSeveridade(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		SeveridadeMinima int `json:"severidade_minima"`
		SeveridadeMaxima int `json:"severidade_maxima"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	severidadeMinima := requestData.SeveridadeMinima
	severidadeMaxima := requestData.SeveridadeMaxima

	// Configura o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	crimes, err := classes.ConsultaCrimesPorSeveridade(severidadeMinima, severidadeMaxima)
	if err != nil {
		http.Error(w, "Crimes não encontrado ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(crimes)
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}

func CtrlDeleteCrime(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		IDCrime   int    `json:"id_crime"`
		NomeHeroi string `json:"nome_heroi"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	idCrime := requestData.IDCrime
	nomeHeroi := requestData.NomeHeroi

	// Configura o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	err = classes.DeleteCrime(idCrime, nomeHeroi)
	if err != nil {
		http.Error(w, "Erro ao deletar crime", http.StatusNotFound)
		return
	}
}

func CtrlAtualizarCrime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idOcorrencia := vars["id"]

	var crimes classes.Crimes
	err := json.NewDecoder(r.Body).Decode(&crimes)
	if err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		return
	}

	err = classes.AtualizarCrime(idOcorrencia, crimes)
	if err != nil {
		http.Error(w, "Erro ao atualizar missão", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
