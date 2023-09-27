package repositories

import (
	"context"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/models"
)

type TransactionRepository interface {
	SaveTx(ctx context.Context, tx *models.Transaction) error
	UpdateTx(ctx context.Context, tx *models.Transaction) error
	GetTx(ctx context.Context, txID string) (*models.Transaction, error)
}
