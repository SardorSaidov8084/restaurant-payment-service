package services

import (
	"context"

	cardcreated "github.com/SardorSaidov8084/restaurant-payment-service/src/application/integration_events/events/card_created"
	carddeleted "github.com/SardorSaidov8084/restaurant-payment-service/src/application/integration_events/events/card_deleted"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/card/services"
)
type PaymentCardApplicationService interface {
	HandleCardCreatedEvent(ctx context.Context, event *cardcreated.Event) error
	HandleCardDeleteEvent(ctx context.Context, event *carddeleted.Event) error
}

type paymentCardAppSvcImpl struct {
	cardSvc services.PaymentCardService
}

func NewPaymentCardApplicationService(cardSvc services.PaymentCardService) PaymentCardApplicationService{
	return &paymentCardAppSvcImpl {
		cardSvc: cardSvc,
	}
}

func (s *paymentCardAppSvcImpl) HandleCardCreatedEvent(ctx context.Context, event *cardcreated.Event) error {

	_, err := s.cardSvc.SaveCard(ctx, event.Card.ID, event.Card.CardToken)
	if err != nil {
		return err
	}
	return nil
}
func (s *paymentCardAppSvcImpl) HandleCardDeleteEvent(ctx context.Context, event *carddeleted.Event) error {

	if err := s.cardSvc.DeleteCard(ctx, event.CardID); err != nil {
		return err
	}
	return nil
}
