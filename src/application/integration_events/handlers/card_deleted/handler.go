package cardcreated

import (
	"context"
	"encoding/json"

	carddeleted "github.com/SardorSaidov8084/restaurant-payment-service/src/application/integration_events/events/card_deleted"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/application/services"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

type Handler struct {
	cardAppSvc services.PaymentCardApplicationService
	logger *zap.Logger
}
func NewHandler(cardAppSvc services.PaymentCardApplicationService, logger *zap.Logger) *Handler {
	return &Handler{
		cardAppSvc: cardAppSvc,
		logger: logger,
	}
}
func (h *Handler) Handle(msg *kafka.Message)  error {

	var (
		event carddeleted.Event
		err error
	)
	h.logger.Debug("started handling event", zap.String("msg", msg.String()))
	defer func() {
		h.logger.Debug("finishing handling event", zap.String("msg", msg.String()))
		if err != nil {
			h.logger.Error("error while handling event",
				zap.String("msg", msg.String()),
				zap.Error(err),
			)
		}
	}()
	if err = json.Unmarshal(msg.Value, &event); err != nil {
		return err
	}
	if err = h.cardAppSvc.HandleCardDeleteEvent(context.Background(), &event);err != nil {
		return err
	}
	return nil
}