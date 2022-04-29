package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/guid"
	"strings"
	"time"
)

//
// client.go - Redis client
// @org: gfx
// @auth: fangtongle
// @date: 2022/4/29 09:53
// @version: 1.0.0.develop
//

type Lock struct {
	// 锁的key
	key string
	// 锁的值
	value string
	// 锁过期时间
	expire int
	// 是否锁
	isLock bool
}

var rdb *redis.Client

var clusterRdb *redis.ClusterClient

// InitRedisClient 初始化redis
func InitRedisClient() {
	var err error
	options := NewDefaultOptions()
	switch options.Mode {
	case "standalone":
		rdb = redis.NewClient(&redis.Options{
			Addr:         options.Host,
			Password:     options.Password,
			DB:           options.Database,
			PoolSize:     options.Pool.MaxIdle,
			MinIdleConns: options.Pool.MinIdle,
			PoolTimeout:  time.Duration(options.Pool.MaxWait) * time.Millisecond,
			DialTimeout:  time.Duration(options.Timeout) * time.Millisecond,
		})
		if _, err = rdb.Ping(context.TODO()).Result(); err != nil {
			glog.Errorf("standalone redis connect error: %s", err.Error())
			return
		}
	case "sentinel":
		rdb = redis.NewFailoverClient(&redis.FailoverOptions{
			SentinelAddrs: strings.Split(options.Host, ","),
			Password:      options.Password,
			DB:            options.Database,
			PoolSize:      options.Pool.MaxIdle,
			MinIdleConns:  options.Pool.MinIdle,
			PoolTimeout:   time.Duration(options.Pool.MaxWait) * time.Millisecond,
			DialTimeout:   time.Duration(options.Timeout) * time.Millisecond,
		})
		if _, err = rdb.Ping(context.TODO()).Result(); err != nil {
			glog.Errorf("sentinel redis connect error: %s", err.Error())
			return
		}
	case "cluster":
		clusterRdb = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        strings.Split(options.Host, ","),
			Password:     options.Password,
			PoolSize:     options.Pool.MaxIdle,
			MinIdleConns: options.Pool.MinIdle,
			PoolTimeout:  time.Duration(options.Pool.MaxWait) * time.Millisecond,
			DialTimeout:  time.Duration(options.Timeout) * time.Millisecond,
		})
		if _, err = rdb.Ping(context.TODO()).Result(); err != nil {
			glog.Errorf("cluster redis connect error: %s", err.Error())
			return
		}
	default:
		glog.Errorf("redis client init error: %s", "mode error")
		return
	}
}

// Expire 指定缓存失效时间
// key: 键
// timeout: 时间(秒)
func Expire(key string, timeout int) (bool, error) {
	return rdb.Expire(context.TODO(), key, time.Duration(timeout)*time.Second).Result()
}

// GetExpire 根据key 获取过期时间
// key: 键 不能为null
// return: 时间(秒) 返回0代表为永久有效
func GetExpire(key string) (int, error) {
	val, err := rdb.TTL(context.TODO(), key).Result()
	return int(val.Seconds()), err
}

// HasKey 判断key是否存在
// key 键
// return true 存在 false不存在
func HasKey(key string) (bool, error) {
	val, err := rdb.Exists(context.TODO(), key).Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}

// Del 删除缓存
// key: 可以传一个值 或多个
func Del(keys ...string) {
	rdb.Del(context.TODO(), keys...)
}

// ============================String=============================

// Get 普通缓存获取
// key 键
// return 值
func Get(key string) (string, error) {
	return rdb.Get(context.TODO(), key).Result()
}

// GetStruct 普通缓存获取
// key 键
// target 参数地址
// return 值
func GetStruct(key string, target interface{}) error {
	val, err := Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), target)
}

// SetStruct 普通缓存放入
// key
// value 转为json string
func SetStruct(key string, value interface{}) error {
	return SetStructExpire(key, value, 0)
}

// SetStructExpire 普通缓存放入
// key
// value 转为json string
// expire 过期时间，秒
func SetStructExpire(key string, value interface{}, expire int) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return rdb.Set(context.TODO(), key, string(bytes), time.Duration(expire)*time.Second).Err()
}

// Set 普通缓存放入
// key 键
// value 值
func Set(key string, value interface{}) error {
	return SetExpire(key, value, 0)
}

// SetExpire 普通缓存放入并设置时间
// key 键
// value 值
// expire 时间(秒) time要大于0 如果time小于等于0 将设置无限期
func SetExpire(key string, value interface{}, expire int) error {
	var t = expire
	if expire < 0 {
		t = 0
	}
	return rdb.Set(context.TODO(), key, value, time.Duration(t)*time.Second).Err()
}

// SetNXExpire 放入缓存并设置时间
// key 键
// value 值
// expire 时间(秒) time要大于0 如果time小于等于0 将设置无限期
func SetNXExpire(key string, value interface{}, expire int) error {
	return rdb.SetNX(context.TODO(), key, value, time.Duration(expire)*time.Second).Err()
}

// SetNX 放入缓存并设置时间
// key 键
// value 值
func SetNX(key string, value interface{}) error {
	return SetNXExpire(key, value, 0)
}

// Incr 递增
// key 键
func Incr(key string) (int64, error) {
	return rdb.Incr(context.TODO(), key).Result()
}

// IncrBy 递增
// key 键
// delta 要减少几(小于0)
func IncrBy(key string, delta int) (int64, error) {
	return rdb.IncrBy(context.TODO(), key, int64(delta)).Result()
}

// ================================Map=================================

// HGet HashGet
// key 键 不能为null
// item 项 不能为null
// return 值
func HGet(key string, item string) (string, error) {
	return rdb.HGet(context.TODO(), key, item).Result()
}

// HGetAll 获取hashKey对应的所有键值
// key 键
// return 对应的多个键值
func HGetAll(key string) (map[string]string, error) {
	return rdb.HGetAll(context.TODO(), key).Result()
}

// HMSet HashSet
// key 键
// map 对应多个键值
func HMSet(key string, fields map[string]interface{}) error {
	return rdb.HMSet(context.TODO(), key, fields).Err()
}

// HSet 向一张hash表中放入数据,如果不存在将创建
// key 键
// item 项
// value 值
func HSet(key string, item string, value interface{}) error {
	return rdb.HSet(context.TODO(), key, item, value).Err()
}

// HDel 删除hash表中的值
// key 键
// item 项 可以使多个
func HDel(key string, item ...string) {
	rdb.HDel(context.TODO(), key, item...)
}

// HHasKey 判断hash表中是否有该项的值
// key 键
// item 项
// return true 存在 false不存在
func HHasKey(key string, item string) (bool, error) {
	return rdb.HExists(context.TODO(), key, item).Result()
}

// HIncr hash递增 如果不存在,就会创建一个 并把新增后的值返回
// key 键
// item 项
// by 要增加几
func HIncr(key string, item string, by int64) (int64, error) {
	return rdb.HIncrBy(context.TODO(), key, item, by).Result()
}

// ============================set=============================

// SGet 根据key获取Set中的所有值
// key 键
func SGet(key string) ([]string, error) {
	return rdb.SMembers(context.TODO(), key).Result()
}

// SIsMember 根据value从一个set中查询,是否存在
//  key 键
// value 值
func SIsMember(key string, value interface{}) (bool, error) {
	return rdb.SIsMember(context.TODO(), key, value).Result()
}

// SRemove 移除值为value的
// key 键
// values 值 可以是多个
func SRemove(key string, values ...interface{}) (int64, error) {
	return rdb.SRem(context.TODO(), key, values).Result()
}

// ===============================list=================================

// GetListSize 获取list缓存的长度
// key 键
func GetListSize(key string) (int64, error) {
	return rdb.LLen(context.TODO(), key).Result()
}

// GetIndex 通过索引 获取list中的值
// key 键
// index 索引 index>=0时， 0 表头，1 第二个元素，依次类推；index<0时，-1，表尾，-2倒数第二个元素，依次类推
func GetIndex(key string, index int) (string, error) {
	return rdb.LIndex(context.TODO(), key, int64(index)).Result()
}

// RPush 将value放入list右缓存
// key 键
// values 值
func RPush(key string, values ...interface{}) (int64, error) {
	return rdb.RPush(context.TODO(), key, values...).Result()
}

// LPop 从list左pop一个元素
// key 键
func LPop(key string) (string, error) {
	return rdb.LPop(context.TODO(), key).Result()
}

// LRemove 移除N个值为value
// key 键
// count 移除多少个
// value 值
func LRemove(key string, count int, value interface{}) (int64, error) {
	return rdb.LRem(context.TODO(), key, int64(count), value).Result()
}

// ================ 锁 =====================

const lockPrefix = "LOCK:"

// New 获取锁对象
// key
// expire 时间(秒) time要大于0 如果time小于等于0 将设置无限期
func New(key string, expire int) *Lock {
	return &Lock{
		key:    lockPrefix + key,
		value:  guid.S(),
		expire: expire,
	}
}

// TryLock 获取锁
func (s *Lock) TryLock() (bool, error) {
	err := SetNXExpire(s.key, s.value, s.expire)
	if err != nil {
		return false, err
	}
	s.isLock = true
	return true, err
}

// UnLock 解锁
func (s *Lock) UnLock() bool {
	if s.isLock {
		val, err := Get(s.key)
		if err != nil {
			return false
		}
		if s.value == val {
			Del(s.key)
		}
	}
	return true
}

// ================ lua脚本 =====================

func NewScript(script string, params []string) (interface{}, error) {
	newScript := redis.NewScript(script)
	return newScript.Run(context.TODO(), rdb, params).Result()
}
