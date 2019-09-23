package main

import (
	"context"
	"github.com/fwidjaya20/demo-distributed-event-store/pkg/protobuf/eventstore"
	"github.com/fwidjaya20/demo-distributed-order-system/cmd/grpcclient"
	"github.com/gofrs/uuid"
	"log"
)

func main() {
	log.Println("Order System")

	grpcClient := grpcclient.InitClients()

	eventId, _ := uuid.NewV4()
	aggregateId, _ := uuid.NewV4()

	res, err := grpcClient.EventStore.CreateEvent(context.Background(), &eventstore.Event{
		EventId:		eventId.String(),
		EventType:     	"OrderSystem.Order",
		AggregateId:   	aggregateId.String(),
		AggregateType: 	"order",
		EventData:     	`{"id": 1, "description": "order has been created"}`,
		Channel:       	"OrderSystem.Order",
	})

	log.Println("Error", err)
	log.Println("Result", res)
}