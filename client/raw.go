package main

import (
	"cache-service/z_generated/pb"
	"context"
	"fmt"
	"time"
)

func setCache(c pb.CacheClient, key string, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)
	defer cancel()
	_, err := c.Set(ctx, &pb.CacheSetInput{Key: key, Value: []byte(value)})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ok")
}

func getCache(c pb.CacheClient, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)
	defer cancel()
	resp, err := c.Get(ctx, &pb.CacheGetInput{Key: key})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(resp.GetValue()))
}
