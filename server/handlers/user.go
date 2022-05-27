package main

import (
	"cache-service/z_generated/pb"
	"context"
	"fmt"
)

func (s *server) SetUser(ctx context.Context, in *pb.User) (*pb.SetUserResp, error) {
	key := userKey(in.GetName(), in.GetClass(), in.GetRollnum())
	err := s.cache.Set(ctx, key, in.GetMetadata())
	return &pb.SetUserResp{}, err
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserInput) (*pb.User, error) {
	key := userKey(in.GetName(), in.GetClass(), in.GetRollnum())
	resp, err := s.cache.Get(ctx, key)
	user := &pb.User{
		Name:     in.GetName(),
		Class:    in.GetClass(),
		Rollnum:  in.GetRollnum(),
		Metadata: resp,
	}
	return user, err
}

func userKey(name string, class string, rollno int64) string {
	return fmt.Sprintf("%s:%s:%d", name, class, rollno)
}
