package validators

import (
	"time"

	"github.com/matheusrbarbosa/gofin/domain/models"
)

type CreateTransactionRequest struct {
	Value       float32   `json:"value" binding:"required"`
	Date        time.Time `json:"date" binding:"required" time_format:"2006-01-02"`
	Observation string    `json:"observation" binding:"lte=255"`
}

func (r *CreateTransactionRequest) ParseToTransaction(vaultId int) models.Transaction {
	return models.Transaction{
		VaultId:     vaultId,
		Value:       r.Value,
		Date:        r.Date,
		Observation: r.Observation,
	}
}
