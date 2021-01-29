package pub

import (
	"context"
	publish "gin-vue-admin/pubsub/protocal"
	"google.golang.org/grpc"
	"log"
)

type Puber struct {
	pubClient publish.DBMSClient
}

func (puber *Puber)Pub(sentinelClusterID uint32, message *publish.PubsubMessage) error {
	_, err := puber.pubClient.Publish(context.Background(), &publish.PublishRequest{
		Topic: &publish.Topic{
			SentinelClusterId: sentinelClusterID,
		},
		Message: message,
	})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func CreatePuber()  *Puber {
	cc, err := grpc.Dial(":10086", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := publish.NewDBMSClient(cc)
	return &Puber{
		pubClient: client,
	}
}
