package modelos

type TransactionsSuccessResponse struct {
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}
