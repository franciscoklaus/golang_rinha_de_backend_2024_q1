package controllers

import "net/http"

// CriarTransacao cria uma transação para o cliente
func CriarTransacao(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criar transacao"))
}

// BuscarExtrato busca o extrato do cliente
func BuscarExtrato(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar extrato"))
}
