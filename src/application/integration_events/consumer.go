package integrationevents

import (
	"context"
	"fmt"
	"log"
	"time"

	cardcreated "github.com/SardorSaidov8084/restaurant-payment-service/src/application/integration_events/events/card_created"
	carddeleted "github.com/SardorSaidov8084/restaurant-payment-service/src/application/integration_events/events/card_deleted"
	cardcreatedhandler "github.com/SardorSaidov8084/restaurant-payment-service/src/application/integration_events/handlers/card_created"
	cardcdeletedhandler "github.com/SardorSaidov8084/restaurant-payment-service/src/application/integration_events/handlers/card_deleted"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/application/services"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/infrastructure/config"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/infrastructure/pubsub"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

const (
	clientID = "restaurant-payment-svc-consumer"
	groupID  = "restaurant-payment-svc-group"
)

var (
	topics = []string{
		"card.added",
		"card.deleted",
	}
)

func StartConsumer(
	ctx context.Context,
	conf config.Config,
	logger *zap.Logger,
	cardApp services.PaymentCardApplicationService,
) {
	if err := createTopics(ctx, conf.KafkaAddress); err != nil {
		logger.Warn("error creating topics", zap.Error(err))
	}

	consumer, err := pubsub.NewConsumer(ctx, conf.KafkaAddress, clientID, groupID, pubsub.WithLogger(logger))
	if err != nil {
		log.Fatalf("error creating pubsub.Consumer: %v", err)
	}

	if err := consumer.Subscribe(
		cardcreated.Event{}.Topic(),
		cardcreatedhandler.NewHandler(cardApp, logger).Handle,
	); err != nil {
		log.Fatalf("error register topic handler: %v", err)
	}

	if err := consumer.Subscribe(
		carddeleted.Event{}.Topic(),
		cardcdeletedhandler.NewHandler(cardApp, logger).Handle,
	); err != nil {
		log.Fatalf("error register topic handler: %v", err)
	}
}

func createTopics(ctx context.Context, serverAddr string) error {
	adminConfig := &kafka.ConfigMap{
		"bootstrap.servers": serverAddr,
	}

	// Create an AdminClient
	adminClient, err := kafka.NewAdminClient(adminConfig)
	if err != nil {
		return fmt.Errorf("error creating AdminClient: %s", err.Error())
	}
	defer adminClient.Close()

	var specs []kafka.TopicSpecification

	for _, topic := range topics {
		specs = append(specs, kafka.TopicSpecification{
			Topic:             topic,
			NumPartitions:     1, // Number of partitions for the topic
			ReplicationFactor: 1, // Replication factor for the topic
		})
	}

	// Create the topic
	results, err := adminClient.CreateTopics(
		ctx,
		specs,
		kafka.SetAdminOperationTimeout(5*time.Second),
	)
	if err != nil {
		return fmt.Errorf("error creating topic: %s", err.Error())
	}

	// Print the results
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			return fmt.Errorf("failed to create topic: %s", result.Error)
		}
	}

	return nil
}
