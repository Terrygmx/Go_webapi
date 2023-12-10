package main

import (
	"testing"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var clusterOptions = &redis.ClusterOptions{
	Addrs:    []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	Password: "password",
}

func newCluster() *redis.ClusterClient {
	redisdb := redis.NewClusterClient(clusterOptions)
	return redisdb
}
func ExampleCluster() {
	redisdb := newCluster()

	err := redisdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := redisdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}


func TestRedisClusterExample(t *testing.T) {
	ExampleCluster()
}
