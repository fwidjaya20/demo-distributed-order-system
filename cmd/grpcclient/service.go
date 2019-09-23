package grpcclient

import (
	"github.com/fwidjaya20/demo-distributed-event-store/pkg/protobuf/eventstore"
	"github.com/fwidjaya20/demo-distributed-order-system/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
)

type Clients struct {
	EventStore eventstore.EventStoreClient
	Connections []*grpc.ClientConn
}

func InitClients() *Clients {
	eventStoreGrpc := config.GetEnv(config.EVENT_STORE_GRPC_ADDR)
	eventStoreConn, err := grpc.Dial(eventStoreGrpc, grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		PermitWithoutStream: true,
	}))

	if nil != err {
		log.Fatal("could not connect to", eventStoreGrpc, err)
	}

	return &Clients{
		EventStore: eventstore.NewEventStoreClient(eventStoreConn),
		Connections: []*grpc.ClientConn{
			eventStoreConn,
		},
	}
}

func (c *Clients) Close() {
	for _, con := range c.Connections {
		_ = con.Close()
	}
}