package repositories

import (
	"context"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/models"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/repositories"
	"gorm.io/gorm"
)

type txRepoImpl struct {
	db *gorm.DB
}

func NewTxRepoImpl(db *gorm.DB) repositories.TransactionRepository {
	return &txRepoImpl{
		db: db,
	}
}

const tableTx = "transactions"

func (r *txRepoImpl) SaveTx(ctx context.Context, tx *models.Transaction) error {

	result := r.db.WithContext(ctx).Table(tableTx).Create(tx)

	return result.Error
}
func (r *txRepoImpl) UpdateTx(ctx context.Context, tx *models.Transaction) error {

	result := r.db.WithContext(ctx).Table(tableTx).Save(tx)

	return result.Error
}
func (r *txRepoImpl) GetTx(ctx context.Context, txID string) (*models.Transaction, error) {

	var tx models.Transaction

	result := r.db.WithContext(ctx).Table(tableTx).First(&tx, "id = ?", txID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &tx, nil
}
