package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"slices"

	"github.com/google/uuid"
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

func generateUuid() string {
	isUnique := false
	var generatedId string
	for !isUnique {
		generatedId = uuid.NewString()
		idx := slices.IndexFunc(dataStore, func(c Customer) bool { return c.ID == generatedId })
		if idx == -1 {
			isUnique = true
		}
	}
	return generatedId
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dataStore)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	idx := slices.IndexFunc(dataStore, func(c Customer) bool { return c.ID == id })

	w.Header().Set("Content-Type", "application/json")
	if idx == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{})
	} else {
		w.WriteHeader(http.StatusOK)
		returnData := dataStore[idx]
		json.NewEncoder(w).Encode(returnData)
	}

}

func addCustomer(w http.ResponseWriter, r *http.Request) {

	var data = Customer{}

	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &data)

	data.ID = generateUuid()
	dataStore = append(dataStore, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{})
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var data = Customer{}

	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &data)

	idx := slices.IndexFunc(dataStore, func(c Customer) bool { return c.ID == id })
	w.Header().Set("Content-Type", "application/json")

	if idx == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{})

	} else {
		dataStore[idx].Name = data.Name
		dataStore[idx].Role = data.Role
		dataStore[idx].Email = data.Email
		dataStore[idx].Phone = data.Phone
		dataStore[idx].Contacted = data.Contacted

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{})
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	w.Header().Set("Content-Type", "application/json")

	idx := slices.IndexFunc(dataStore, func(c Customer) bool { return c.ID == id })
	if idx == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{})
	} else {
		dataStore = append(dataStore[:idx], dataStore[idx+1:]...)
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(map[string]string{})
	}
}

func routeNoMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{})

}

func main() {
	//Initial seed
	var custList = []Customer{
		{
			ID:        "57bddb9a-d4a5-4c26-81be-33c5392b83ad",
			Name:      "Customer 1",
			Role:      "Sales Engineer",
			Email:     "cust1@dummy.com",
			Phone:     "123456789",
			Contacted: false,
		},
		{
			ID:        "b7e0ee32-9280-46bb-a3a3-5a042d6eaf5f",
			Name:      "Customer 2",
			Role:      "Civil Engineer",
			Email:     "cust2@dummy.com",
			Phone:     "987654321",
			Contacted: true,
		},
		{
			ID:        "14bf757c-5162-4b4a-9469-fc713b296e68",
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

	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customer", addCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customer/{id}", deleteCustomer).Methods("DELETE")

	router.HandleFunc("*", routeNoMatch)

	fmt.Println("Server is starting on port 8085...")
	http.ListenAndServe(":8085", router)
}
