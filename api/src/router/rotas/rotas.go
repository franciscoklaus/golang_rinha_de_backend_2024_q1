package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da aplicação
type Rota struct {
	URI    string
	Metodo string
	Funcao func(w http.ResponseWriter, r *http.Request)
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasClientes
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
