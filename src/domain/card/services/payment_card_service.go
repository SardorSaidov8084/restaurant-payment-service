package services

import (
	"context"
	"time"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/card/models"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/card/repositories"
)

type PaymentCardService interface {
	SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
}

type cardSvcImpl struct {
	cardRepo repositories.PaymentCardRepository
}

func NewCardSvcImpl(cardRepo repositories.PaymentCardRepository) PaymentCardService {
	return &cardSvcImpl{
		cardRepo: cardRepo,
	}
}

func (s *cardSvcImpl) SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error) {

	card := &models.PaymentCard{
		ID:        cardID,
		CardToken: cardToken,
		CreatedAt: time.Now().UTC(),
	}

	err := s.cardRepo.SaveCard(ctx, card)

	if err != nil {
		return nil, err
	}

	return card, nil
}
func (s *cardSvcImpl) DeleteCard(ctx context.Context, cardID string) error {

	return s.cardRepo.DeleteCard(ctx, cardID)
}
func (s *cardSvcImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {

	return s.cardRepo.GetCard(ctx, cardID)
}
