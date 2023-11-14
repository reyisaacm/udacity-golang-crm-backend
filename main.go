package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var dataStore = []Customer{}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dataStore)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dataStore)
}

func main() {
	//Initial seed
	var custList = []Customer{
		{
			ID:        "1",
			Name:      "Customer 1",
			Role:      "Sales Engineer",
			Email:     "cust1@dummy.com",
			Phone:     "123456789",
			Contacted: false,
		},
		{
			ID:        "2",
			Name:      "Customer 2",
			Role:      "Civil Engineer",
			Email:     "cust2@dummy.com",
			Phone:     "987654321",
			Contacted: true,
		},
		{
			ID:        "3",
			Name:      "Customer 3",
			Role:      "Aeronautics Engineer",
			Email:     "cust3@dummy.com",
			Phone:     "567891234",
			Contacted: true,
		},
	}

	dataStore = append(dataStore, custList...)

	// Instantiate a new router by invoking the "NewRouter" handler
	router := mux.NewRouter()

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customer/{i}",getCustomer).Methods("GET")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
