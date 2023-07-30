// main.go

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moaaskt/vehicle_management_backend/models" // 
)

var users []models.User
var vehicles []models.Vehicle

func main() {
	router := mux.NewRouter()

	// Rota para criar um novo usuário
	router.HandleFunc("/api/users", createUser).Methods("POST")

	// Rota para obter todos os usuários
	router.HandleFunc("/api/users", getUsers).Methods("GET")

	// Rota para criar um novo veículo
	router.HandleFunc("/api/vehicles", createVehicle).Methods("POST")

	// Rota para obter todos os veículos
	router.HandleFunc("/api/vehicles", getVehicles).Methods("GET")

	log.Println("Servidor iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simulação de criação do usuário no banco de dados
	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func createVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simulação de criação do veículo no banco de dados
	vehicles = append(vehicles, vehicle)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vehicle)
}

func getVehicles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(vehicles)
}
