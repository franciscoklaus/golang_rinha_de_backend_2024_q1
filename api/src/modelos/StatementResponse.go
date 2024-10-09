package modelos

import "time"

type Statement struct {
	Saldo      Saldo        `json:"saldo"`
	Transacoes []Transacoes `json:"ultimas_transacoes"`
}

type Saldo struct {
	Total       int64     `json:"total"`
	DataExtrato time.Time `json:"data_extrato" `
	Limite      int64     `json:"limite"`
}

type Transacoes struct {
	Valor         int64     `json:"valor" db:"valor"`
	Tipo          string    `json:"tipo" db:"tipo"`
	Descricao     string    `json:"descricao" db:"descricao"`
	DataTransacao time.Time `json:"realizada_em" db:"realizada_em"`
}
