package rotas

import "net/http"

var rotasClientes = []Rota{
	{
		URI:    "/clientes/{clienteId}/transacoes",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
	},
	{
		URI:    "/clientes/{clienteId}/extrato",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
	},
}
