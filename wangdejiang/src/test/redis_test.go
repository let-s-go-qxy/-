package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestRedisGet(t *testing.T) {
	rdb.Set(ctx, "name", "aa", time.Second*10)
}

func TestRedisSet(t *testing.T) {
	v, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		t.Fail()
	}
	fmt.Println(v)
}
