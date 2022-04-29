package redis

import "github.com/gogf/gf/frame/g"

type Pool struct {
	// 连接池最大阻塞等待时间（使用负值表示没有限制）
	MaxWait int
	// 连接池中的最大连接
	MaxIdle int
	// 连接池中的最小连接
	MinIdle int
}

type Options struct {
	// redis 连接模式，standalone、sentinel、cluster
	Mode string
	// 服务器地址
	Host string
	// 服务器连接密码
	Password string
	// 数据库索引（默认为0）
	Database int
	// 连接超时时间（毫秒）
	Timeout int
	// 连接池配置
	Pool Pool
}

func NewDefaultOptions() Options {
	pool := Pool{
		MaxWait: -1,
		MaxIdle: 10,
		MinIdle: 0,
	}

	ip := g.Config().GetString("redis.host")
	port := g.Config().GetString("redis.port")
	password := g.Config().GetString("redis.password")
	database := g.Config().GetInt("redis.database")

	return Options{
		Mode:     "standalone",
		Host:     ip + ":" + port,
		Password: password,
		Database: database,
		Timeout:  500,
		Pool:     pool,
	}
}
