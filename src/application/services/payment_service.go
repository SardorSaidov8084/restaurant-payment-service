package services

import (
	"context"
	"log"

	cardsvc "github.com/SardorSaidov8084/restaurant-payment-service/src/domain/card/services"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/models"
	paymentsvc "github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/services"
)

type PaymentApplicationService interface {
	CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error)
	PayOrderReceipt(ctx context.Context, orderID string) error
	UpdateTransaction(ctx context.Context, tx *models.Transaction) error
}

type paymentAppSvcImpl struct {
	paymentSvc paymentsvc.PaymentService
	cardSvc    cardsvc.PaymentCardService
}

func NewPaymentApplicationService(paymentSvc paymentsvc.PaymentService, cardSvc cardsvc.PaymentCardService) PaymentApplicationService {

	return &paymentAppSvcImpl{
		paymentSvc: paymentSvc,
		cardSvc:    cardSvc,
	}
}
func (s *paymentAppSvcImpl) CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error) {

	receipt, err := s.paymentSvc.CreateReceipt(ctx, orderID, cardID, amount)

	if err != nil {
		return nil, err
	}
	log.Printf("Payme Creating receipt %s\n", receipt.OrderID)

	err = s.paymentSvc.UpdateReceipt(ctx, receipt)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}
func (s *paymentAppSvcImpl) PayOrderReceipt(ctx context.Context, orderID string) error {

	receipt, err := s.paymentSvc.GetReceiptByOrder(ctx, orderID)

	if err != nil {
		return err
	}
	card, err := s.cardSvc.GetCard(ctx, receipt.CardID)

	if err != nil {
		return err
	}
	log.Printf("Paying receipt: orderID %s - cardToken %s\n", receipt.OrderID, card.CardToken)

	return nil
}
func (s *paymentAppSvcImpl) UpdateTransaction(ctx context.Context, tx *models.Transaction) error {

	err := s.paymentSvc.UpdateTransaction(ctx, tx)

	if err != nil {
		return err
	}

	return nil
}
