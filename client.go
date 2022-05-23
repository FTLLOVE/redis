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

type instance struct {
}

var Instance = new(instance)

var rdb *redis.Client

// redis client for cluster mode
var clusterRdb *redis.ClusterClient

// InitRedisClient 初始化redis
func InitRedisClient(ip string, port int64, password string, db int) {
	var err error
	options := NewDefaultOptions(ip, port, password, db)
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

func (*instance) Ping(ctx context.Context) (string, error) {
	return rdb.Ping(ctx).Result()
}

func (*instance) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) (string, error) {
	return rdb.Migrate(ctx, host, port, key, db, timeout).Result()
}

func (*instance) Move(ctx context.Context, key string, db int) (bool, error) {
	return rdb.Move(ctx, key, db).Result()
}

func (*instance) Exists(ctx context.Context, keys ...string) (int64, error) {
	return rdb.Exists(ctx, keys...).Result()
}

func (*instance) Expire(ctx context.Context, key string, expire time.Duration) (bool, error) {
	return rdb.Expire(ctx, key, expire).Result()
}

func (*instance) Keys(ctx context.Context, pattern string) ([]string, error) {
	return rdb.Keys(ctx, pattern).Result()
}

func (*instance) Del(ctx context.Context, keys ...string) (int64, error) {
	return rdb.Del(ctx, keys...).Result()
}

func (*instance) ObjectRefCount(ctx context.Context, key string) (int64, error) {
	return rdb.ObjectRefCount(ctx, key).Result()
}

func (*instance) ObjectEncoding(ctx context.Context, key string) (string, error) {
	return rdb.ObjectEncoding(ctx, key).Result()
}

func (*instance) ObjectIdleTime(ctx context.Context, key string) (time.Duration, error) {
	return rdb.ObjectIdleTime(ctx, key).Result()
}

func (*instance) Persist(ctx context.Context, key string) (bool, error) {
	return rdb.Persist(ctx, key).Result()
}

func (*instance) PExpire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return rdb.PExpire(ctx, key, expiration).Result()
}

func (*instance) PTTL(ctx context.Context, key string) (time.Duration, error) {
	return rdb.PTTL(ctx, key).Result()
}
func (*instance) RandomKey(ctx context.Context) (string, error) {
	return rdb.RandomKey(ctx).Result()
}

func (*instance) Rename(ctx context.Context, key, newkey string) (string, error) {
	return rdb.Rename(ctx, key, newkey).Result()
}

func (*instance) Restore(ctx context.Context, key string, ttl time.Duration, value string) (string, error) {
	return rdb.Restore(ctx, key, ttl, value).Result()
}

func (*instance) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) (string, error) {
	return rdb.RestoreReplace(ctx, key, ttl, value).Result()
}

func (*instance) Touch(ctx context.Context, keys ...string) (int64, error) {
	return rdb.Touch(ctx, keys...).Result()
}

func (*instance) TTL(ctx context.Context, key string) (time.Duration, error) {
	return rdb.TTL(ctx, key).Result()
}

func (*instance) Append(ctx context.Context, key, value string) (int64, error) {
	return rdb.Append(ctx, key, value).Result()
}

func (*instance) Decr(ctx context.Context, key string) (int64, error) {
	return rdb.Decr(ctx, key).Result()
}

func (*instance) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	return rdb.DecrBy(ctx, key, decrement).Result()
}

func (*instance) Get(ctx context.Context, key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func (*instance) GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	return rdb.GetRange(ctx, key, start, end).Result()
}

func (*instance) GetSet(ctx context.Context, key string, value interface{}) (string, error) {
	return rdb.GetSet(ctx, key, value).Result()
}

func (*instance) GetEx(ctx context.Context, key string, expiration time.Duration) (string, error) {
	return rdb.GetEx(ctx, key, expiration).Result()
}

func (*instance) GetDel(ctx context.Context, key string) (string, error) {
	return rdb.GetDel(ctx, key).Result()
}

func (*instance) Incr(ctx context.Context, key string) (int64, error) {
	return rdb.Incr(ctx, key).Result()
}

func (*instance) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return rdb.IncrBy(ctx, key, value).Result()
}

func (*instance) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	return rdb.IncrByFloat(ctx, key, value).Result()
}

func (*instance) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	return rdb.MGet(ctx, keys...).Result()
}

func (*instance) MSet(ctx context.Context, values ...interface{}) (string, error) {
	return rdb.MSet(ctx, values...).Result()
}

func (*instance) MSetNX(ctx context.Context, values ...interface{}) (bool, error) {
	return rdb.MSetNX(ctx, values...).Result()
}

func (*instance) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	return rdb.Set(ctx, key, value, expiration).Result()
}

func (*instance) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	return rdb.SetEX(ctx, key, value, expiration).Result()
}

func (*instance) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return rdb.SetNX(ctx, key, value, expiration).Result()
}

func (*instance) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return rdb.SetXX(ctx, key, value, expiration).Result()
}

func (*instance) SetRange(ctx context.Context, key string, offset int64, value string) (int64, error) {
	return rdb.SetRange(ctx, key, offset, value).Result()
}

func (*instance) StrLen(ctx context.Context, key string) (int64, error) {
	return rdb.StrLen(ctx, key).Result()
}

func (*instance) Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) (int64, error) {
	return rdb.Copy(ctx, sourceKey, destKey, db, replace).Result()
}

//------------------------------------Bit------------------------------------------

func (*instance) GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	return rdb.GetBit(ctx, key, offset).Result()
}

func (*instance) SetBit(ctx context.Context, key string, offset int64, value int) (int64, error) {
	return rdb.SetBit(ctx, key, offset, value).Result()
}

func (*instance) BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return rdb.BitOpAnd(ctx, destKey, keys...).Result()
}

func (*instance) BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return rdb.BitOpOr(ctx, destKey, keys...).Result()
}

func (*instance) BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return rdb.BitOpXor(ctx, destKey, keys...).Result()
}

func (*instance) BitOpNot(ctx context.Context, destKey string, key string) (int64, error) {
	return rdb.BitOpNot(ctx, destKey, key).Result()
}

func (*instance) BitPos(ctx context.Context, key string, bit int64, pos ...int64) (int64, error) {
	return rdb.BitPos(ctx, key, bit, pos...).Result()
}

func (*instance) BitField(ctx context.Context, key string, args ...interface{}) ([]int64, error) {
	return rdb.BitField(ctx, key, args...).Result()
}

func (*instance) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return rdb.Scan(ctx, cursor, match, count).Result()
}

func (*instance) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) ([]string, uint64, error) {
	return rdb.ScanType(ctx, cursor, match, count, keyType).Result()
}

func (*instance) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return rdb.SScan(ctx, key, cursor, match, count).Result()
}

func (*instance) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return rdb.HScan(ctx, key, cursor, match, count).Result()
}

func (*instance) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return rdb.ZScan(ctx, key, cursor, match, count).Result()
}

func (*instance) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return rdb.HDel(ctx, key, fields...).Result()
}

func (*instance) HExists(ctx context.Context, key, field string) (bool, error) {
	return rdb.HExists(ctx, key, field).Result()
}

func (*instance) HGet(ctx context.Context, key, field string) (string, error) {
	return rdb.HGet(ctx, key, field).Result()
}

func (*instance) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return rdb.HGetAll(ctx, key).Result()
}

func (*instance) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	return rdb.HIncrBy(ctx, key, field, incr).Result()
}

func (*instance) HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error) {
	return rdb.HIncrByFloat(ctx, key, field, incr).Result()
}

func (*instance) HKeys(ctx context.Context, key string) ([]string, error) {
	return rdb.HKeys(ctx, key).Result()
}

func (*instance) HLen(ctx context.Context, key string) (int64, error) {
	return rdb.HLen(ctx, key).Result()
}

func (*instance) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return rdb.HMGet(ctx, key, fields...).Result()
}

func (*instance) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return rdb.HSet(ctx, key, values...).Result()
}

func (*instance) HMSet(ctx context.Context, key string, values ...interface{}) (bool, error) {
	return rdb.HMSet(ctx, key, values...).Result()
}

func (*instance) HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	return rdb.HSetNX(ctx, key, field, value).Result()
}

func (*instance) HVals(ctx context.Context, key string) ([]string, error) {
	return rdb.HVals(ctx, key).Result()
}

func (*instance) HRandField(ctx context.Context, key string, count int, withValues bool) ([]string, error) {
	return rdb.HRandField(ctx, key, count, withValues).Result()
}

func (*instance) BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return rdb.BLPop(ctx, timeout, keys...).Result()
}

func (*instance) BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return rdb.BRPop(ctx, timeout, keys...).Result()
}

func (*instance) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (string, error) {
	return rdb.BRPopLPush(ctx, source, destination, timeout).Result()
}

func (*instance) LIndex(ctx context.Context, key string, index int64) (string, error) {
	return rdb.LIndex(ctx, key, index).Result()
}

func (*instance) LInsert(ctx context.Context, key, op string, pivot, value interface{}) (int64, error) {
	return rdb.LInsert(ctx, key, op, pivot, value).Result()
}

func (*instance) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	return rdb.LInsertBefore(ctx, key, pivot, value).Result()
}

func (*instance) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	return rdb.LInsertAfter(ctx, key, pivot, value).Result()
}

func (*instance) LLen(ctx context.Context, key string) (int64, error) {
	return rdb.LLen(ctx, key).Result()
}

func (*instance) LPop(ctx context.Context, key string) (string, error) {
	return rdb.LPop(ctx, key).Result()
}

func (*instance) LPopCount(ctx context.Context, key string, count int) ([]string, error) {
	return rdb.LPopCount(ctx, key, count).Result()
}

func (*instance) LPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return rdb.LPush(ctx, key, values...).Result()
}

func (*instance) LPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return rdb.LPushX(ctx, key, values...).Result()
}

func (*instance) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return rdb.LRange(ctx, key, start, stop).Result()
}

func (*instance) LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error) {
	return rdb.LRem(ctx, key, count, value).Result()
}
func (*instance) LSet(ctx context.Context, key string, index int64, value interface{}) (string, error) {
	return rdb.LSet(ctx, key, index, value).Result()
}

func (*instance) LTrim(ctx context.Context, key string, start, stop int64) (string, error) {
	return rdb.LTrim(ctx, key, start, stop).Result()
}

func (*instance) RPop(ctx context.Context, key string) (string, error) {
	return rdb.RPop(ctx, key).Result()
}

func (*instance) RPopCount(ctx context.Context, key string, count int) ([]string, error) {
	return rdb.RPopCount(ctx, key, count).Result()
}

func (*instance) RPopLPush(ctx context.Context, source, destination string) (string, error) {
	return rdb.RPopLPush(ctx, source, destination).Result()
}

func (*instance) RPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return rdb.RPush(ctx, key, values...).Result()
}

func (*instance) RPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return rdb.RPushX(ctx, key, values...).Result()
}

func (*instance) LMove(ctx context.Context, source, destination, srcpos, destpos string) (string, error) {
	return rdb.LMove(ctx, source, destination, srcpos, destpos).Result()
}

func (*instance) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) (string, error) {
	return rdb.BLMove(ctx, source, destination, srcpos, destpos, timeout).Result()
}

func (*instance) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return rdb.SAdd(ctx, key, members...).Result()
}

func (*instance) SCard(ctx context.Context, key string) (int64, error) {
	return rdb.SCard(ctx, key).Result()
}

func (*instance) SDiff(ctx context.Context, keys ...string) ([]string, error) {
	return rdb.SDiff(ctx, keys...).Result()
}

func (*instance) SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return rdb.SDiffStore(ctx, destination, keys...).Result()
}

func (*instance) SInter(ctx context.Context, keys ...string) ([]string, error) {
	return rdb.SInter(ctx, keys...).Result()
}

func (*instance) SInterStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return rdb.SInterStore(ctx, destination, keys...).Result()
}

func (*instance) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return rdb.SIsMember(ctx, key, member).Result()
}

func (*instance) SMIsMember(ctx context.Context, key string, members ...interface{}) ([]bool, error) {
	return rdb.SMIsMember(ctx, key, members...).Result()
}

func (*instance) SMembers(ctx context.Context, key string) ([]string, error) {
	return rdb.SMembers(ctx, key).Result()
}

func (*instance) SMembersMap(ctx context.Context, key string) (map[string]struct{}, error) {
	return rdb.SMembersMap(ctx, key).Result()
}

func (*instance) SMove(ctx context.Context, source, destination string, member interface{}) (bool, error) {
	return rdb.SMove(ctx, source, destination, member).Result()
}

func (*instance) SPop(ctx context.Context, key string) (string, error) {
	return rdb.SPop(ctx, key).Result()
}

func (*instance) SPopN(ctx context.Context, key string, count int64) ([]string, error) {
	return rdb.SPopN(ctx, key, count).Result()
}

func (*instance) SRandMember(ctx context.Context, key string) (string, error) {
	return rdb.SRandMember(ctx, key).Result()
}

func (*instance) SRandMemberN(ctx context.Context, key string, count int64) ([]string, error) {
	return rdb.SRandMemberN(ctx, key, count).Result()
}

func (*instance) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return rdb.SRem(ctx, key, members...).Result()
}

func (*instance) SUnion(ctx context.Context, keys ...string) ([]string, error) {
	return rdb.SUnion(ctx, keys...).Result()
}

func (*instance) SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return rdb.SUnionStore(ctx, destination, keys...).Result()
}

func (*instance) XDel(ctx context.Context, stream string, ids ...string) (int64, error) {
	return rdb.XDel(ctx, stream, ids...).Result()
}

func (*instance) XLen(ctx context.Context, stream string) (int64, error) {
	return rdb.XLen(ctx, stream).Result()
}

func (*instance) XRange(ctx context.Context, stream, start, stop string) ([]redis.XMessage, error) {
	return rdb.XRange(ctx, stream, start, stop).Result()
}

func (*instance) XRangeN(ctx context.Context, stream, start, stop string, count int64) ([]redis.XMessage, error) {
	return rdb.XRangeN(ctx, stream, start, stop, count).Result()
}

func (*instance) XRevRange(ctx context.Context, stream, start, stop string) ([]redis.XMessage, error) {
	return rdb.XRevRange(ctx, stream, start, stop).Result()
}

func (*instance) XRevRangeN(ctx context.Context, stream, start, stop string, count int64) ([]redis.XMessage, error) {
	return rdb.XRevRangeN(ctx, stream, start, stop, count).Result()
}

func (*instance) XReadStreams(ctx context.Context, streams ...string) ([]redis.XStream, error) {
	return rdb.XReadStreams(ctx, streams...).Result()
}

func (*instance) XGroupCreate(ctx context.Context, stream, group, start string) (string, error) {
	return rdb.XGroupCreate(ctx, stream, group, start).Result()
}

func (*instance) XGroupCreateMkStream(ctx context.Context, stream, group, start string) (string, error) {
	return rdb.XGroupCreateMkStream(ctx, stream, group, start).Result()
}

func (*instance) XGroupSetID(ctx context.Context, stream, group, start string) (string, error) {
	return rdb.XGroupSetID(ctx, stream, group, start).Result()
}

func (*instance) XGroupDestroy(ctx context.Context, stream, group string) (int64, error) {
	return rdb.XGroupDestroy(ctx, stream, group).Result()
}

func (*instance) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) (int64, error) {
	return rdb.XGroupCreateConsumer(ctx, stream, group, consumer).Result()
}

func (*instance) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) (int64, error) {
	return rdb.XGroupDelConsumer(ctx, stream, group, consumer).Result()
}

func (*instance) XAck(ctx context.Context, stream, group string, ids ...string) (int64, error) {
	return rdb.XAck(ctx, stream, group, ids...).Result()
}

func (*instance) XPending(ctx context.Context, stream, group string) (*redis.XPending, error) {
	return rdb.XPending(ctx, stream, group).Result()
}

func (*instance) XTrim(ctx context.Context, key string, maxLen int64) (int64, error) {
	return rdb.XTrim(ctx, key, maxLen).Result()
}

func (*instance) XTrimApprox(ctx context.Context, key string, maxLen int64) (int64, error) {
	return rdb.XTrimApprox(ctx, key, maxLen).Result()
}

func (*instance) XTrimMaxLen(ctx context.Context, key string, maxLen int64) (int64, error) {
	return rdb.XTrimMaxLen(ctx, key, maxLen).Result()
}

func (*instance) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) (int64, error) {
	return rdb.XTrimMaxLenApprox(ctx, key, maxLen, limit).Result()
}

func (*instance) XTrimMinID(ctx context.Context, key string, minID string) (int64, error) {
	return rdb.XTrimMinID(ctx, key, minID).Result()
}

func (*instance) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) (int64, error) {
	return rdb.XTrimMinIDApprox(ctx, key, minID, limit).Result()
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
