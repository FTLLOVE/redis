package redis

import (
	"context"
	"testing"
	"time"
)

//
// client_test.go
// @org: gfx
// @auth: fangtongle
// @date: 2022/5/22 13:20
// @version: 1.0.0.develop
//

var (
	ip         = "127.0.0.1"
	port int64 = 6379
	db         = 0
	pwd        = ""
)

func init() {
	InitRedisClient(ip, port, pwd, db)
}

func TestInstance_Ping(t *testing.T) {
	t.Log(Instance.Ping(context.TODO()))
}

func TestInstance_Exists(t *testing.T) {
	t.Log(Instance.Exists(context.TODO(), "test", "test1"))
}

func TestInstance_Expire(t *testing.T) {
	t.Log(Instance.Expire(context.TODO(), "test", time.Second*10))
}

func TestInstance_Keys(t *testing.T) {
	t.Log(Instance.Keys(context.TODO(), "*"))
}

func TestInstance_Del(t *testing.T) {
	t.Log(Instance.Del(context.TODO(), "test"))
}
