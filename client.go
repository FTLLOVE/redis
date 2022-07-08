package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
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

func (*instance) ZAdd(ctx context.Context, key string, members ...*Z) (int64, error) {
	var args []*redis.Z
	if err := gconv.SliceStruct(members, &args); err != nil {
		return 0, err
	}
	return rdb.ZAdd(ctx, key, args...).Result()
}

func (*instance) ZAddNX(ctx context.Context, key string, members ...*Z) (int64, error) {
	var args []*redis.Z
	if err := gconv.SliceStruct(members, &args); err != nil {
		return 0, err
	}
	return rdb.ZAddNX(ctx, key, args...).Result()
}

func (*instance) ZAddXX(ctx context.Context, key string, members ...*Z) (int64, error) {
	var args []*redis.Z
	if err := gconv.SliceStruct(members, &args); err != nil {
		return 0, err
	}
	return rdb.ZAddXX(ctx, key, args...).Result()
}

func (*instance) ZIncr(ctx context.Context, key string, member *Z) (float64, error) {
	var args *redis.Z
	if err := gconv.Struct(member, &args); err != nil {
		return 0, err
	}
	return rdb.ZIncr(ctx, key, args).Result()
}

func (*instance) ZIncrNX(ctx context.Context, key string, member *Z) (float64, error) {
	var args *redis.Z
	if err := gconv.Struct(member, &args); err != nil {
		return 0, err
	}
	return rdb.ZIncrNX(ctx, key, args).Result()
}

func (*instance) ZIncrXX(ctx context.Context, key string, member *Z) (float64, error) {
	var args *redis.Z
	if err := gconv.Struct(member, &args); err != nil {
		return 0, err
	}
	return rdb.ZIncrXX(ctx, key, args).Result()
}

func (*instance) ZCard(ctx context.Context, key string) (int64, error) {
	return rdb.ZCard(ctx, key).Result()
}

func (*instance) ZCount(ctx context.Context, key, min, max string) (int64, error) {
	return rdb.ZCount(ctx, key, min, max).Result()
}

func (*instance) ZLexCount(ctx context.Context, key, min, max string) (int64, error) {
	return rdb.ZLexCount(ctx, key, min, max).Result()
}

func (*instance) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	return rdb.ZIncrBy(ctx, key, increment, member).Result()
}

func (*instance) ZInter(ctx context.Context, store *ZStore) ([]string, error) {
	var args *redis.ZStore
	if err := gconv.Struct(store, &args); err != nil {
		return nil, err
	}
	return rdb.ZInter(ctx, args).Result()
}

func (*instance) ZInterWithScores(ctx context.Context, store *ZStore) ([]Z, error) {
	var args *redis.ZStore
	if err := gconv.Struct(store, &args); err != nil {
		return nil, err
	}
	zs, err := rdb.ZInterWithScores(ctx, args).Result()
	if err != nil {
		return nil, err
	}
	var result []Z
	if err := gconv.Structs(zs, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (*instance) ZInterStore(ctx context.Context, destination string, store *ZStore) (int64, error) {
	var args *redis.ZStore
	if err := gconv.Struct(store, &args); err != nil {
		return 0, err
	}
	return rdb.ZInterStore(ctx, destination, args).Result()
}

func (*instance) ZMScore(ctx context.Context, key string, members ...string) ([]float64, error) {
	return rdb.ZMScore(ctx, key, members...).Result()
}

func (*instance) ZPopMax(ctx context.Context, key string, count ...int64) ([]Z, error) {
	var result []Z
	zs, err := rdb.ZPopMax(ctx, key, count...).Result()
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(zs, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (*instance) ZPopMin(ctx context.Context, key string, count ...int64) ([]Z, error) {
	var result []Z
	zs, err := rdb.ZPopMin(ctx, key, count...).Result()
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(zs, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (*instance) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return rdb.ZRange(ctx, key, start, stop).Result()
}

func (*instance) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	var result []Z
	zs, err := rdb.ZRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(zs, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (*instance) ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) ([]string, error) {
	var args *redis.ZRangeBy
	if err := gconv.Struct(opt, &args); err != nil {
		return nil, err
	}
	return rdb.ZRangeByScore(ctx, key, args).Result()
}

func (*instance) ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) ([]string, error) {
	var args *redis.ZRangeBy
	if err := gconv.Struct(opt, &args); err != nil {
		return nil, err
	}
	return rdb.ZRangeByLex(ctx, key, args).Result()
}

func (*instance) ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) ([]Z, error) {
	var args *redis.ZRangeBy
	if err := gconv.Struct(opt, &args); err != nil {
		return nil, err
	}
	var data []Z
	zs, err := rdb.ZRangeByScoreWithScores(ctx, key, args).Result()
	if err = gconv.Structs(zs, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (*instance) ZRank(ctx context.Context, key, member string) (int64, error) {
	return rdb.ZRank(ctx, key, member).Result()
}

func (*instance) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return rdb.ZRem(ctx, key, members...).Result()
}

func (*instance) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	return rdb.ZRemRangeByRank(ctx, key, start, stop).Result()
}

func (*instance) ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error) {
	return rdb.ZRemRangeByScore(ctx, key, min, max).Result()
}

func (*instance) ZRemRangeByLex(ctx context.Context, key, min, max string) (int64, error) {
	return rdb.ZRemRangeByLex(ctx, key, min, max).Result()
}

func (*instance) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return rdb.ZRevRange(ctx, key, start, stop).Result()
}

func (*instance) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	zs, err := rdb.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	var result []Z
	if err := gconv.Structs(zs, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (*instance) ZRevRank(ctx context.Context, key, member string) (int64, error) {
	return rdb.ZRevRank(ctx, key, member).Result()
}

func (*instance) ZScore(ctx context.Context, key, member string) (float64, error) {
	return rdb.ZScore(ctx, key, member).Result()
}

func (*instance) ZUnionStore(ctx context.Context, dest string, store *ZStore) (int64, error) {
	var args *redis.ZStore
	if err := gconv.Struct(store, &args); err != nil {
		return 0, err
	}
	return rdb.ZUnionStore(ctx, dest, args).Result()
}

func (*instance) ZUnion(ctx context.Context, store ZStore) ([]string, error) {
	var args redis.ZStore
	if err := gconv.Struct(store, &args); err != nil {
		return nil, nil
	}
	return rdb.ZUnion(ctx, args).Result()
}

func (*instance) ZUnionWithScores(ctx context.Context, store ZStore) ([]Z, error) {
	var args redis.ZStore
	if err := gconv.Struct(store, &args); err != nil {
		return nil, nil
	}
	result, err := rdb.ZUnionWithScores(ctx, args).Result()
	if err != nil {
		return nil, err
	}
	var data []Z
	if err = gconv.Structs(result, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (*instance) ZRandMember(ctx context.Context, key string, count int, withScores bool) ([]string, error) {
	return rdb.ZRandMember(ctx, key, count, withScores).Result()
}

func (*instance) ZDiff(ctx context.Context, keys ...string) ([]string, error) {
	return rdb.ZDiff(ctx, keys...).Result()
}

func (*instance) ZDiffWithScores(ctx context.Context, keys ...string) ([]Z, error) {
	zs, err := rdb.ZDiffWithScores(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}
	var data []Z
	if err := gconv.Structs(zs, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (*instance) ZDiffStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return rdb.ZDiffStore(ctx, destination, keys...).Result()
}

func (*instance) PFAdd(ctx context.Context, key string, els ...interface{}) (int64, error) {
	return rdb.PFAdd(ctx, key, els...).Result()
}

func (*instance) PFCount(ctx context.Context, keys ...string) (int64, error) {
	return rdb.PFCount(ctx, keys...).Result()
}

func (*instance) PFMerge(ctx context.Context, dest string, keys ...string) (string, error) {
	return rdb.PFMerge(ctx, dest, keys...).Result()
}

func (*instance) BgRewriteAOF(ctx context.Context) (string, error) {
	return rdb.BgRewriteAOF(ctx).Result()
}

func (*instance) BgSave(ctx context.Context) (string, error) {
	return rdb.BgSave(ctx).Result()
}

func (*instance) ClientKill(ctx context.Context, ipPort string) (string, error) {
	return rdb.ClientKill(ctx, ipPort).Result()
}

func (*instance) ClientKillByFilter(ctx context.Context, keys ...string) (int64, error) {
	return rdb.ClientKillByFilter(ctx, keys...).Result()
}

func (*instance) ClientList(ctx context.Context) (string, error) {
	return rdb.ClientList(ctx).Result()
}

func (*instance) ClientPause(ctx context.Context, dur time.Duration) (bool, error) {
	return rdb.ClientPause(ctx, dur).Result()
}

func (*instance) ClientID(ctx context.Context) (int64, error) {
	return rdb.ClientID(ctx).Result()
}

func (*instance) ConfigGet(ctx context.Context, parameter string) ([]interface{}, error) {
	return rdb.ConfigGet(ctx, parameter).Result()
}

func (*instance) ConfigResetStat(ctx context.Context) (string, error) {
	return rdb.ConfigResetStat(ctx).Result()
}

func (*instance) ConfigSet(ctx context.Context, parameter, value string) (string, error) {
	return rdb.ConfigSet(ctx, parameter, value).Result()
}

func (*instance) ConfigRewrite(ctx context.Context) (string, error) {
	return rdb.ConfigRewrite(ctx).Result()
}

func (*instance) DBSize(ctx context.Context) (int64, error) {
	return rdb.DBSize(ctx).Result()
}

func (*instance) FlushAll(ctx context.Context) (string, error) {
	return rdb.FlushAll(ctx).Result()
}

func (*instance) FlushAllAsync(ctx context.Context) (string, error) {
	return rdb.FlushAllAsync(ctx).Result()
}

func (*instance) FlushDB(ctx context.Context) (string, error) {
	return rdb.FlushDB(ctx).Result()
}

func (*instance) FlushDBAsync(ctx context.Context) (string, error) {
	return rdb.FlushDBAsync(ctx).Result()
}

func (*instance) Info(ctx context.Context, section ...string) (string, error) {
	return rdb.Info(ctx, section...).Result()
}

func (*instance) LastSave(ctx context.Context) (int64, error) {
	return rdb.LastSave(ctx).Result()
}

func (*instance) Save(ctx context.Context) (string, error) {
	return rdb.Save(ctx).Result()
}

func (*instance) Shutdown(ctx context.Context) (string, error) {
	return rdb.Shutdown(ctx).Result()
}

func (*instance) ShutdownSave(ctx context.Context) (string, error) {
	return rdb.ShutdownSave(ctx).Result()
}

func (*instance) ShutdownNoSave(ctx context.Context) (string, error) {
	return rdb.ShutdownNoSave(ctx).Result()
}

func (*instance) SlaveOf(ctx context.Context, host, port string) (string, error) {
	return rdb.SlaveOf(ctx, host, port).Result()
}

func (*instance) Time(ctx context.Context) (time.Time, error) {
	return rdb.Time(ctx).Result()
}

func (*instance) DebugObject(ctx context.Context, key string) (string, error) {
	return rdb.DebugObject(ctx, key).Result()
}

func (*instance) ReadOnly(ctx context.Context) (string, error) {
	return rdb.ReadOnly(ctx).Result()
}

func (*instance) ReadWrite(ctx context.Context) (string, error) {
	return rdb.ReadWrite(ctx).Result()
}

func (*instance) MemoryUsage(ctx context.Context, key string, samples ...int) (int64, error) {
	return rdb.MemoryUsage(ctx, key, samples...).Result()
}

func (*instance) Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	return rdb.Eval(ctx, script, keys, args...).Result()
}

func (*instance) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return rdb.EvalSha(ctx, sha1, keys, args...).Result()
}

func (*instance) ScriptExists(ctx context.Context, hashes ...string) ([]bool, error) {
	return rdb.ScriptExists(ctx, hashes...).Result()
}

func (*instance) ScriptFlush(ctx context.Context) (string, error) {
	return rdb.ScriptFlush(ctx).Result()
}

func (*instance) ScriptKill(ctx context.Context) (string, error) {
	return rdb.ScriptKill(ctx).Result()
}

func (*instance) ScriptLoad(ctx context.Context, script string) (string, error) {
	return rdb.ScriptLoad(ctx, script).Result()
}

func (*instance) Publish(ctx context.Context, channel string, message interface{}) (int64, error) {
	return rdb.Publish(ctx, channel, message).Result()
}

func (*instance) PubSubChannels(ctx context.Context, pattern string) ([]string, error) {
	return rdb.PubSubChannels(ctx, pattern).Result()
}

func (*instance) PubSubNumSub(ctx context.Context, channels ...string) (map[string]int64, error) {
	return rdb.PubSubNumSub(ctx, channels...).Result()
}

func (*instance) PubSubNumPat(ctx context.Context) (int64, error) {
	return rdb.PubSubNumPat(ctx).Result()
}
