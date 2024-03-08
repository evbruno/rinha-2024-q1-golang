package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type Transaction struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type Account struct {
	ID         string
	Limite     int
	Saldo      int
	Transacoes []Transaction
}

var accounts = map[string]*Account{}

func main() {
	addSampleAccount()
	http.HandleFunc("/clientes/", clientesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func clientesHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.NotFound(w, r)
		return
	}

	id := path[2]
	if id == "" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodPost:
		handlePost(w, r, id)
	case http.MethodGet:
		handleGet(w, r, id)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request, id string) {
	account, ok := accounts[id]
	if !ok {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	var transaction Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if transaction.Tipo != "c" && transaction.Tipo != "d" {
		http.Error(w, "Invalid transaction type", 5400)
		return
	}

	account.Transacoes = append(account.Transacoes, transaction)

	if transaction.Tipo == "c" {
		account.Saldo += transaction.Valor
	} else {
		account.Saldo -= transaction.Valor
	}

	response := map[string]int{
		"limite": account.Limite,
		"saldo":  account.Saldo,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleGet(w http.ResponseWriter, r *http.Request, id string) {
	account, ok := accounts[id]
	if !ok {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	total := 0
	for _, transaction := range account.Transacoes {
		if transaction.Tipo == "c" {
			total += transaction.Valor
		} else {
			total -= transaction.Valor
		}
	}

	response := map[string]interface{}{
		"saldo": map[string]interface{}{
			"total":        total,
			"data_extrato": time.Now().Format(time.RFC3339),
			"limite":       account.Limite,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Function to add a sample account for demonstration
func addSampleAccount() {
	account := &Account{
		ID:     "123",
		Limite: 1000,
		Saldo:  500,
		Transacoes: []Transaction{
			{Valor: 200, Tipo: "c", Descricao: "Initial Deposit"},
		},
	}

	accounts["123"] = account
}
