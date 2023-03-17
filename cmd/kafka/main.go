// Sources for https://watermill.io/docs/getting-started/
package main

import (
	"hulk/go-webservice/infrastructure"
	"hulk/go-webservice/infrastructure/persist"
	"hulk/go-webservice/infrastructure/services"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

func messageHandler(msg *message.Message) {
	log.Printf("[Pubsub] received message: %s, payload: %s", msg.UUID, string(msg.Payload))
}

var topic string = "example.topic"

func main() {
	persist.InitDB()

	infrastructure.RegisterAppServices()
	services.PubSubInstance.SubscribeTopic(topic, messageHandler)
	publishMessages()
}

func publishMessages() {
	for {
		if err := services.PubSubInstance.PublishMessage(topic, []byte("Hello, world!")); err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}
