package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarTransacao cria uma transação para o cliente
func CriarTransacao(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parametros := mux.Vars(r)
	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	if clienteID < 1 || clienteID > 5 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cliente não encontrado"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var transacao modelos.TransactionRequest

	if erro = json.Unmarshal(corpoRequisicao, &transacao); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter os dados da requisição"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeTransacoes(db)
	response, erro := repositorio.CriarTransacao(transacao, clienteID)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar transação"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

// BuscarExtrato busca o extrato do cliente
func BuscarExtrato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parametros := mux.Vars(r)
	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	if clienteID < 1 || clienteID > 5 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cliente não encontrado"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	repositorio := repositorios.NovoRepositorioDeTransacoes(db)
	extrato, erro := repositorio.BuscarExtrato(clienteID)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar extrato"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(extrato)
}
