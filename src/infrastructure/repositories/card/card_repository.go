package repositories

import (
	"context"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/card/models"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/card/repositories"
	"gorm.io/gorm"
)

type cardRepoImpl struct {
	db *gorm.DB
}

func NewCardRepoImpl(db *gorm.DB) repositories.PaymentCardRepository {

	return &cardRepoImpl{
		db: db,
	}
}

const tablePaymentCard = "payment_cards"

func (r *cardRepoImpl) SaveCard(ctx context.Context, card *models.PaymentCard) error {

	result := r.db.WithContext(ctx).Table(tablePaymentCard).Create(card)

	return result.Error
}
func (r *cardRepoImpl) DeleteCard(ctx context.Context, cardID string) error {

	var card models.PaymentCard

	result := r.db.WithContext(ctx).Table(tablePaymentCard).Delete(&card, "id = ?", cardID)

	return result.Error
}
func (r *cardRepoImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {

	var card models.PaymentCard

	result := r.db.WithContext(ctx).Table(tablePaymentCard).First(&card, "id = ?", cardID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &card, nil
}
