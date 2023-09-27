package dtos

import (
	"time"

	pb "github.com/SardorSaidov8084/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/models"
)

func ToReceiptPB(receipt *models.Receipt) *pb.Receipt {
    return &pb.Receipt{
		Id: receipt.ID,
		CardId: receipt.CardID,
		Amount: int32(receipt.Amount),
		Data: receipt.Data,
		CreatedAt: receipt.CreatedAt.Format(time.RFC3339),
		UpdatedAt: receipt.UpdatedAt.Format(time.RFC3339),
	}
}