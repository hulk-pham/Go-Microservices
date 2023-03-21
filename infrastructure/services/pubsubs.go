package services

import (
	"context"
	"fmt"
	"hulk/go-webservice/infrastructure/config"
	"log"

	"github.com/Shopify/sarama"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

type PubSubService struct {
	subscriber *kafka.Subscriber
	publisher  *kafka.Publisher
}

type messageHandler func(*message.Message)

var PubSubInstance *PubSubService

func InitPubSubService() {
	config := config.AppConfig()
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{fmt.Sprintf("%s:%s", config.KafkaHost, config.KafkaPort)},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         config.KafkaConsumeGroup,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{fmt.Sprintf("%s:%s", config.KafkaHost, config.KafkaPort)},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		log.Printf("[Pubsub] Unable init kafka connection. Err: %s", err.Error())
	}

	PubSubInstance = &PubSubService{
		subscriber,
		publisher,
	}
}

func (service *PubSubService) SubscribeTopic(topicName string, handler messageHandler) (err error) {
	messages, err := PubSubInstance.subscriber.Subscribe(context.Background(), topicName)
	if err != nil {
		return err
	}

	go processInternal(messages, handler)
	return
}

func (service *PubSubService) PublishMessage(topicName string, data []byte) (err error) {
	msg := message.NewMessage(watermill.NewUUID(), data)

	log.Printf("[Pubsub] publish message: %s, payload: %s, topic: %s", msg.UUID, string(msg.Payload), topicName)
	if err := PubSubInstance.publisher.Publish(topicName, msg); err != nil {
		return err
	}
	return
}

func processInternal(messages <-chan *message.Message, handler messageHandler) {
	for msg := range messages {
		handler(msg)
		msg.Ack()
	}
}
