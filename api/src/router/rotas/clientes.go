package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotasClientes é uma coleção das rotas de clientes
var rotasClientes = []Rota{
	{
		URI:    "/clientes/{clienteId}/transacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarTransacao,
	},
	{
		URI:    "/clientes/{clienteId}/extrato",
		Metodo: http.MethodPost,
		Funcao: controllers.BuscarExtrato,
	},
}
