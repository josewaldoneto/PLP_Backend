package main

import (
	"PLP_Backend/controllers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Define as rotas da aplicação
func Loading() {
	r := mux.NewRouter()

	// suporte a CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	// Rotas para os herois
	r.HandleFunc("/", controllers.MostraTodosOsNomesHerois).Methods("GET")
	r.HandleFunc("/heroi", controllers.MostraPorNome).Methods("POST")
	r.HandleFunc("/heroipop", controllers.MostraPopularidade).Methods("POST")
	r.HandleFunc("/heroicadastra", controllers.CadastraHeroi).Methods("POST")
	r.HandleFunc("/delete", controllers.DeletaAKAralha).Methods("DELETE")
	r.HandleFunc("/heroistatus", controllers.MostraPorStatus).Methods("POST")
	r.HandleFunc("/poderes", controllers.MostraTodosPoderes).Methods("GET")
	r.HandleFunc("/editar", controllers.EditarHeroiHandler).Methods("POST")

	// Rotas para os crimes
	r.HandleFunc("/heroieseveridadecrime", controllers.ConsultaCrimesHS).Methods("POST")
	r.HandleFunc("/heroicrime", controllers.ConsultaCrimesHeroi).Methods("POST")
	r.HandleFunc("/severidadecrime", controllers.ConsultaCrimesSeveridade).Methods("POST")
	r.HandleFunc("/deletecrime/", controllers.CtrlDeleteCrime).Methods("DELETE")
	r.HandleFunc("/editacrime/{id}", controllers.CtrlAtualizarCrime).Methods("POST")

	// Rotas para as missoes
	r.HandleFunc("/missao", controllers.ConsultaMissaoHeroi).Methods("POST")
	r.HandleFunc("/missoes", controllers.ListarTodasMissoesHandler).Methods("GET")
	r.HandleFunc("/missao/{id}", controllers.ConsultaMissaoPorId).Methods("GET")
	r.HandleFunc("/missao/{id}", controllers.AtualizarMissao).Methods("POST")
	r.HandleFunc("/missao/{id}", controllers.DeletarMissaoHandler).Methods("DELETE")

	r.HandleFunc("/simularbatalha", controllers.SimularBatalhaController).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins, credentials)(r)))
}
