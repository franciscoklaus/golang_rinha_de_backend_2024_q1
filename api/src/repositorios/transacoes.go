package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
	"time"
)

// Transacoes representa um repositório de transações
type Transacoes struct {
	db *sql.DB
}

// NovoRepositorioDeTransacoes cria um novo repositório de transações
func NovoRepositorioDeTransacoes(db *sql.DB) *Transacoes {
	return &Transacoes{db}
}

func (repositorio Transacoes) CriarTransacao(transacoes modelos.TransactionRequest, clienteID uint64) (modelos.TransactionsSuccessResponse, error) {
	statement, erro := repositorio.db.Prepare("insert into transacoes (valor, tipo, descricao, cliente_id) values (?, ?, ?, ?)")
	if erro != nil {
		fmt.Println(erro)
		return modelos.TransactionsSuccessResponse{}, erro

	}
	defer statement.Close()

	if _, erro = statement.Exec(transacoes.Valor, transacoes.Tipo, transacoes.Descricao, clienteID); erro != nil {
		fmt.Println(erro)
		return modelos.TransactionsSuccessResponse{}, erro
	}

	if transacoes.Tipo == "d" {
		statement, erro = repositorio.db.Prepare("update clientes set saldo_atual = saldo_atual - ? where id = ?")
		if erro != nil {
			fmt.Println(erro)
			return modelos.TransactionsSuccessResponse{}, erro
		}
		defer statement.Close()

		if _, erro = statement.Exec(transacoes.Valor, clienteID); erro != nil {
			fmt.Println(erro)
			return modelos.TransactionsSuccessResponse{}, erro
		}
	}

	if transacoes.Tipo == "c" {
		statement, erro = repositorio.db.Prepare("update clientes set saldo_atual = saldo_atual + ? where id = ?")
		if erro != nil {
			return modelos.TransactionsSuccessResponse{}, erro
		}
		defer statement.Close()

		if _, erro = statement.Exec(transacoes.Valor, clienteID); erro != nil {
			return modelos.TransactionsSuccessResponse{}, erro
		}
	}

	// Consultando o saldo e o limite atualizados
	var saldoAtual, limite int64
	erro = repositorio.db.QueryRow("select saldo_atual, limite from clientes where id = ?", clienteID).Scan(&saldoAtual, &limite)
	if erro != nil {
		return modelos.TransactionsSuccessResponse{}, erro
	}

	// Retornando a resposta com o saldo e limite
	response := modelos.TransactionsSuccessResponse{
		Limite: limite,
		Saldo:  saldoAtual,
	}

	return response, nil
}

func (repositorio Transacoes) BuscarExtrato(clienteID uint64) (modelos.Statement, error) {
	var statement modelos.Statement
	// Definindo a data de solicitação do extrato e a data do extrato como o valor de time.Now()
	now := time.Now()
	statement.Saldo.DataExtrato = now // Preenchendo DataExtrato com a data atual

	// Query para buscar o saldo
	querySaldo := `SELECT saldo_atual, limite FROM clientes WHERE id = ?`
	err := repositorio.db.QueryRow(querySaldo, clienteID).Scan(
		&statement.Saldo.Total,
		&statement.Saldo.Limite,
	)
	if err != nil {
		fmt.Println(err)
		return modelos.Statement{}, fmt.Errorf("erro ao buscar saldo: %v", err)
	}

	// Query para buscar as últimas transações
	queryTransacoes := `SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE cliente_id = ? ORDER BY realizada_em DESC LIMIT 10`
	rows, err := repositorio.db.Query(queryTransacoes, clienteID)
	if err != nil {
		fmt.Println(err)
		return modelos.Statement{}, fmt.Errorf("erro ao buscar transações: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var transacao modelos.Transacoes
		err := rows.Scan(
			&transacao.Valor,
			&transacao.Tipo,
			&transacao.Descricao,
			&transacao.DataTransacao,
		)
		if err != nil {
			fmt.Println(err)
			return modelos.Statement{}, fmt.Errorf("erro ao escanear transação: %v", err)
		}
		statement.Transacoes = append(statement.Transacoes, transacao)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return modelos.Statement{}, fmt.Errorf("erro durante a iteração das transações: %v", err)
	}

	return statement, nil
}
