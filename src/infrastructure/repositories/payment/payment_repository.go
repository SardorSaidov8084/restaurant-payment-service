package repositories

import (
	"context"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/models"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/repositories"
	"gorm.io/gorm"
)

type paymentRepoImpl struct {
	db *gorm.DB
}

func NewPaymentRepoImpl(db *gorm.DB) repositories.ReceiptRepository {
	return &paymentRepoImpl{
		db: db,
	}
}

const tableReceipt = "receipts"

func (r *paymentRepoImpl) SaveReceipt(ctx context.Context, receipt *models.Receipt) error {

	result := r.db.WithContext(ctx).Table(tableReceipt).Create(receipt)

	return result.Error
}
func (r *paymentRepoImpl) UpdateReceipt(ctx context.Context, receipt *models.Receipt) error {

	result := r.db.WithContext(ctx).Table(tableReceipt).Save(receipt)

	return result.Error
}
func (r *paymentRepoImpl) GetReceipt(ctx context.Context, receiptID string) (*models.Receipt, error) {

	var receipt models.Receipt

	result := r.db.WithContext(ctx).Table(tableReceipt).First(&receipt, "id = ?", receiptID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipt, nil
}
func (r *paymentRepoImpl) GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error) {

	var receipt models.Receipt

	result := r.db.WithContext(ctx).Table(tableReceipt).First(&receipt, "order_id = ?", orderID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipt, nil
}
