package modelos

type TransactionRequest struct {
	Valor     int64  `json:"valor" db:"valor" binding:"required,min=1"`
	Tipo      string `json:"tipo" db:"tipo" binding:"required,max=1,oneof=d c"`
	Descricao string `json:"descricao" db:"descricao" binding:"required,max=10"`
}
