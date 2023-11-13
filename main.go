package main

import (
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
	Contacted bool   `json:"contacted`
}

var dataStore = []Customer{}

func getCustomerList(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Instantiate a new router by invoking the "NewRouter" handler
	router := mux.NewRouter()

	router.HandleFunc("/customers", getCustomerList).Methods("GET")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
