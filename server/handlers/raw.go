package main

import (
	"cache-service/z_generated/pb"
	"context"
)

func (s *server) Set(ctx context.Context, in *pb.CacheSetInput) (*pb.CacheSetResp, error) {
	err := s.cache.Set(ctx, in.GetKey(), in.GetValue())
	return &pb.CacheSetResp{}, err
}

func (s *server) Get(ctx context.Context, in *pb.CacheGetInput) (*pb.CacheGetResp, error) {

	resp, err := s.cache.Get(ctx, in.GetKey())

	if err != nil {
		return nil, err
	}
	return &pb.CacheGetResp{
		Value: resp,
	}, nil
}
