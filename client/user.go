package main

import (
	"cache-service/z_generated/pb"
	"context"
	"fmt"
	"strconv"
	"time"
)

func setUserCache(c pb.CacheClient, name string, class string, roll string, metadata string) {

	rollInt, err := strconv.ParseInt(roll, 10, 64)
	if err != nil {
		fmt.Println("rollno must be integer value")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)
	defer cancel()
	_, err = c.SetUser(ctx, &pb.User{
		Name:     name,
		Class:    class,
		Rollnum:  rollInt,
		Metadata: []byte(metadata),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ok")
}

func getUserCache(c pb.CacheClient, name string, class string, roll string) {

	rollInt, err := strconv.ParseInt(roll, 10, 64)
	if err != nil {
		fmt.Println("rollno be integer value")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)
	defer cancel()
	user, err := c.GetUser(ctx, &pb.GetUserInput{
		Name:    name,
		Class:   class,
		Rollnum: rollInt,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user)
}
