package pubsub

import (
	"context"
	"gin-vue-admin/global"
	publish "gin-vue-admin/pubsub/protocal"
	"github.com/docker/docker/pkg/pubsub"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
	Pub *pubsub.Publisher
}

func (s *Server) Publish(c context.Context, pub *publish.PublishRequest) (*publish.PublishResponse, error) {
	s.Pub.Publish(pub)
	return &publish.PublishResponse{MessageId: 1}, nil
}

func (s *Server) Subscribe(req *publish.Topic, stream publish.Publisher_SubscribeServer) error {
	ch := s.Pub.SubscribeTopic(func(v interface{}) bool {
		if req.GetName() == "" {
			return true
		}
		if it, ok := v.(*publish.PublishRequest); ok {
			if it.Topic.GetName() == req.GetName() {
				return true
			}
		}
		return false
	})
	for v := range ch {
		if pub, ok := v.(*publish.PublishRequest); ok {
			if err := stream.Send(pub.GetMessages()); err != nil {
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
	global.GVA_LOG.Info("启动订阅服务...")
	lis, err := net.Listen("tcp", ":10086")
	if err != nil {
		global.GVA_LOG.Fatal("订阅服务模块监听失败!", zap.Any("err", err))
	}
	s := grpc.NewServer()
	publish.RegisterPublisherServer(s, &Server{Pub: pubsub.NewPublisher(100*time.Millisecond, 10)})
	if err := s.Serve(lis); err != nil {
		global.GVA_LOG.Fatal("订阅服务模块启动失败!", zap.Any("err", err))
	}
}
