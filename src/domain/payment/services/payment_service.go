package services

import (
	"context"
	"time"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/models"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/repositories"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/infrastructure/rand"
)

type PaymentService interface {
	CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error)
	GetReceipt(ctx context.Context, receiptID string) (*models.Receipt, error)
	PayReceipt(ctx context.Context, receiptID string) error
	GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error)
	UpdateReceipt(ctx context.Context, receipt *models.Receipt) error
	SaveTransaction(ctx context.Context, tx *models.Transaction) error
	UpdateTransaction(ctx context.Context, tx *models.Transaction) error
	GetTransaction(ctx context.Context, txID string) (*models.Transaction, error)
}

type paymentSvcImpl struct {
	receiptRepo repositories.ReceiptRepository
	txRepo      repositories.TransactionRepository
}

func NewPaymentSvcImpl(receiptRepo repositories.ReceiptRepository, txRepo repositories.TransactionRepository) PaymentService {
	return &paymentSvcImpl{
		receiptRepo: receiptRepo,
		txRepo:      txRepo,
	}
}

func (s *paymentSvcImpl) CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error) {

	receipt, err := s.receiptRepo.GetReceiptByOrder(ctx, orderID)

	if err != nil {
		return s.createReceipt(ctx, orderID, cardID, amount)
	}
	receipt.CardID = cardID
	receipt.Amount = amount

	if err = s.receiptRepo.UpdateReceipt(ctx, receipt); err != nil {
		return nil, err
	}
	return receipt, nil
}
func (s *paymentSvcImpl) createReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error) {

	now := time.Now().UTC()
	receipt := models.Receipt{
		ID:        rand.UUID(),
		OrderID:   orderID,
		CardID:    cardID,
		Amount:    amount,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.receiptRepo.SaveReceipt(ctx, &receipt)
	if err != nil {
		return nil, err
	}
	return &receipt, nil
}
func (s *paymentSvcImpl) PayReceipt(ctx context.Context, receiptID string) error {
	return nil
}
func (s *paymentSvcImpl) GetReceipt(ctx context.Context, receiptID string) (*models.Receipt, error) {

	return s.receiptRepo.GetReceipt(ctx, receiptID)
}
func (s *paymentSvcImpl) GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error) {
	return s.receiptRepo.GetReceiptByOrder(ctx, orderID)
}
func (s *paymentSvcImpl) UpdateReceipt(ctx context.Context, receipt *models.Receipt) error {

	receipt.UpdatedAt = time.Now().UTC()
	return s.UpdateReceipt(ctx, receipt)
}
func (s *paymentSvcImpl) SaveTransaction(ctx context.Context, tx *models.Transaction) error {

	return s.txRepo.SaveTx(ctx, tx)
}
func (s *paymentSvcImpl) UpdateTransaction(ctx context.Context, tx *models.Transaction) error {

	tx.UpdatedAt = time.Now().UTC()
	return s.txRepo.UpdateTx(ctx, tx)
}
func (s *paymentSvcImpl) GetTransaction(ctx context.Context, txID string) (*models.Transaction, error) {

	return s.txRepo.GetTx(ctx, txID)
}
