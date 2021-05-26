package adder

import (
	"context"

	"github.com/lenarbatdalov/grpcadder/pkg/api"
)

type GRPCServer struct{}

// интерфейс AdderServer требует реализовать метод Add
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponce, error) {
	return &api.AddResponce{Result: req.GetX() + req.GetY()}, nil
}
