package config

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

var DB_Redis *redis.Client

func ConnectRedis() {
	opt, err := redis.ParseURL("redis://:313@localhost:6379")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	rdb := redis.NewClient(opt)
	//rdb.Set("test", "ete", -1)
	//fmt.Printf("vALUE %v \n", val)

	DB_Redis = rdb
}
