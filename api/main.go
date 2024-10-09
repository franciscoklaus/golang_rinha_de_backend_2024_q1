package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()
	fmt.Println("Porta da API:", config.PortaApi)
	fmt.Println("String de conexão:", config.StringConexaoBanco)
	fmt.Println("Rodando API...")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PortaApi), r))
}
