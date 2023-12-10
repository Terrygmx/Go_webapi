package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func CreateClusterOptions(nodes []string, pwd string) *redis.ClusterOptions {

	return &redis.ClusterOptions{
		Addrs:    nodes,
		Password: pwd,
	}
}
func CreateCluster(options *redis.ClusterOptions) *redis.ClusterClient {
	redisdb := redis.NewClusterClient(options)
	return redisdb
}

// Redis setting key to value
func SetStringToCluster(redisdb *redis.ClusterClient, key string, v string) (err error) {
	// redisdb := newCluster()
	err = redisdb.Set(ctx, key, v, 0).Err()
	if err != nil {
		panic(err)
	}
	return
}

// Redis getting value
func GetStringFromCluster(rdb *redis.ClusterClient, key string) (v string, err error) {
	v, err = rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		v = "key not exist"
		err = nil
		return
	} else if err != nil {
		panic(err)
	}
	return
}
