package subscriber

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type ShippingService struct{}

func (e *ShippingService) Handle(ctx context.Context, msg *pb.Empty) error {
	log.Info("Handler Received message: ")
	return nil
}

func Handler(ctx context.Context, msg *pb.Empty) error {
	log.Info("Function Received message: ")
	return nil
}
