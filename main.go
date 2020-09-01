package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rs/xid"
)

func main() {
	fmt.Println("7")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	r := xid.New().String()
	isok, err := rdb.SetNX(ctx, "myKey", r, 5*time.Second).Result()
	fmt.Println(isok, err)
	defer func() {
		fmt.Println(666)
		isok, err := rdb.Get(ctx, "myKey").Result()
		if isok == r && err == nil {
			fmt.Println("Release...")
			rdb.Del(ctx, "myKey")
		}
	}()

	time.Sleep(time.Second * 3)

}
