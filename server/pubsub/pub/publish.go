package pubsub

import (
	"context"
	"fmt"
	publish "gin-vue-admin/pubsub/protocal"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Puber struct {
	pubClient publish.DBMSClient
}

func (puber *Puber)Pub()  {
	for {
		resp, err := puber.pubClient.Publish(context.Background(), &publish.PublishRequest{
			Topic: &publish.Topic{
				Name: "golang",
			},
			Messages: &publish.PubsubMessage{
				Data: []byte("welcome!"),
			},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.MessageId)
		time.Sleep(time.Second)
	}
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
