package redis

//定义与Redis交互的客户端
import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

// 初始化
func NewRedisClient() *RedisClient {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // 确保地址正确
			Password: "",               // 如果没有密码，留空
			DB:       0,                // 使用默认数据库
		}),
	}
}

// 将验证码存于redis，设置超时时间
func (r *RedisClient) SetCode(email, code string, expiration time.Duration) error {
	ctx := context.Background()
	err := r.client.Set(ctx, email, code, expiration).Err()
	if err != nil {
		log.Printf("Failed to set code in Redis: %v", err)
		return err
	}
	return nil
}

// 用于从Redis中获取存储的验证码
func (r *RedisClient) GetCode(email string) (string, error) {
	ctx := context.Background()
	code, err := r.client.Get(ctx, email).Result()
	if err == redis.Nil {
		return "", nil // 如果键不存在，返回空字符串
	} else if err != nil {
		return "", err
	}
	return code, nil
}
