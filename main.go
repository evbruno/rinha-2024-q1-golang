package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Valor       int    `json:"valor"`
	Tipo        string `json:"tipo"`
	Descricao   string `json:"descricao"`
	RealizadaEm string `json:"realizada_em"`
}
type Account struct {
	ID         int
	Limite     int
	Saldo      int
	Transacoes []Transaction
}

var accounts = map[int]*Account{}

func main() {
	addSampleAccounts()
	http.HandleFunc("/clientes/", clientesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func clientesHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(path[2])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	action := path[3]

	if r.Method == http.MethodPost && action == "transacoes" {
		handlePost(w, r, id)
	} else if r.Method == http.MethodGet && action == "extrato" {
		handleGet(w, r, id)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request, id int) {
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
		http.Error(w, "Invalid transaction type", http.StatusBadRequest)
		return
	}

	if len(transaction.Descricao) == 0 || len(transaction.Descricao) > 10 {
		http.Error(w, "Invalid description", http.StatusBadRequest)
		return
	}

	// REGRAS REGRAS REGRAS REGRAS
	novoSaldo := account.Saldo
	if transaction.Tipo == "c" {
		novoSaldo += transaction.Valor
	} else {
		novoSaldo -= transaction.Valor
	}

	if novoSaldo < -account.Limite {
		http.Error(w, "Invalid balance", http.StatusUnprocessableEntity)
		return
	}

	account.Saldo = novoSaldo
	transaction.RealizadaEm = time.Now().Format(time.RFC3339)

	//account.Transacoes = append(account.Transacoes, transaction)
	// prepend
	account.Transacoes = append([]Transaction{transaction}, account.Transacoes...)

	response := map[string]int{
		"limite": account.Limite,
		"saldo":  account.Saldo,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleGet(w http.ResponseWriter, r *http.Request, id int) {
	account, ok := accounts[id]
	if !ok {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	take := min(len(account.Transacoes), 10)

	response := map[string]interface{}{
		"saldo": map[string]interface{}{
			"total":        account.Saldo,
			"data_extrato": time.Now().Format(time.RFC3339),
			"limite":       account.Limite,
		},
		"transacoes": account.Transacoes[:take],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Function to add a sample account for demonstration
func addAccount(id int, limite int, saldo int) {
	account := &Account{
		ID:         id,
		Limite:     limite,
		Saldo:      saldo,
		Transacoes: []Transaction{},
	}
	accounts[id] = account
}

func addSampleAccounts() {
	addAccount(1, 100000, 0)
	addAccount(2, 80000, 0)
	addAccount(3, 1000000, 0)
	addAccount(4, 10000000, 0)
	addAccount(5, 500000, 0)
}
