package redis

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"testing"
)

func TestBasic(t *testing.T) {
	var err error
	InitRedisClient()
	t.Log(Set("test", "test"))
	t.Log(SetExpire("test", "test", 100))
	t.Log(GetExpire("test"))
	HSet("person1", "name", "fangtongle")
	HMSet("person2", map[string]interface{}{
		"name": "fangtongle",
		"age":  "18",
	})
	type Person struct {
		Name string
		Age  int
	}
	hGetAll, _ := HGetAll("person2")
	var p *Person
	if err = gconv.Struct(hGetAll, &p); err != nil {
		glog.Errorf("%v", err)
		return
	}
	t.Log(p)
	Del("person2")
	t.Log(HMSet("product_01", map[string]interface{}{
		"product_name": "满100减10",
		"product_num":  100,
	}))
}

func TestLock(t *testing.T) {
	InitRedisClient()
	lock := New("user_01", 0)
	tryLock, err := lock.TryLock()
	if err != nil {
		t.Error(err)
		return
	}
	glog.Info("tryLock:", tryLock)
	if tryLock {
		glog.Info("lock success")
		lock.UnLock()
		glog.Info("unlock success")
	} else {
		glog.Info("lock failed")
	}
}

func TestLua(t *testing.T) {
	InitRedisClient()
	scriptResult, err := NewScript(`
		if call("hexists", KEYS[1], KEYS[2]) == 1 then
			local stock = tonumber(call("hget", KEYS[1], KEYS[2]))
			if stock > 0 then
				call("hincrby", KEYS[1], KEYS[2], -1)
				return stock
			end
				return 0
		end
	`, []string{"product_01", "product_num"})
	glog.Info(scriptResult, err)
}
