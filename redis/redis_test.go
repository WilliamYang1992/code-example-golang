package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

const (
	Host     = "127.0.0.1"
	Port     = 6379
	Password = ""
	DB       = 0
)

func Example() {
	// 创建 client
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", Host, Port),
		Password: Password,
		DB:       DB,
	})
	// 连接测试
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	// 设置键值
	key := "key"
	val := "value"
	err := client.Set(key, val, 0).Err()
	if err != nil {
		panic(err)
	}
	// 获取键值
	val, err = client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %s\n", key, val)
	// Output:
	// key: value
}
