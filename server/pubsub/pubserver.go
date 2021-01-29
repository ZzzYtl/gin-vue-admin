package pubsub

import (
	"context"
	publish "gin-vue-admin/pubsub/protocal"
	"github.com/docker/docker/pkg/pubsub"
	"gin-vue-admin/pubsub/dbms"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	sentinel.RpcServer
	Pub *pubsub.Publisher
}

func (s *Server) Publish(c context.Context, pub *publish.PublishRequest) (*publish.PublishResponse, error) {
	s.Pub.Publish(pub)
	return &publish.PublishResponse{}, nil
}

func (s *Server) Subscribe(req *publish.Topic, stream publish.DBMS_SubscribeServer) error {
	ch := s.Pub.SubscribeTopic(func(v interface{}) bool {
		if it, ok := v.(*publish.PublishRequest); ok {
			if it.Topic.GetSentinelClusterId() == req.GetSentinelClusterId() {
				return true
			}
		}
		return false
	})
	for v := range ch {
		if pub, ok := v.(*publish.PublishRequest); ok {
			if err := stream.Send(pub.GetMessage()); err != nil {
				return err
			}
		}
	}
	return nil
}

func RunPublishServer()  {
	go RunPubslishServerInternal()
}

func RunPubslishServerInternal()  {
	log.Println("启动订阅服务...")
	lis, err := net.Listen("tcp", ":10086")
	if err != nil {
		log.Fatal("订阅服务模块监听失败!", zap.Any("err", err))
	}
	s := grpc.NewServer()
	publish.RegisterDBMSServer(s, &Server{Pub: pubsub.NewPublisher(100*time.Millisecond, 10)})
	if err := s.Serve(lis); err != nil {
		log.Fatal("订阅服务模块启动失败!", zap.Any("err", err))
	}
}
