package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type BloomCache struct {
	client *redis.Client
}

func NewBuildBloom(client *redis.Client) *BloomCache {
	return &BloomCache{
		client: client,
	}
}

// CreatBloom 创建布隆过滤器
func (b *BloomCache) CreatBloom(key string, errRate int, size int) error {
	cmd := redis.NewCmd(context.Background(), "BF.RESERVE", key, errRate, size)
	err := b.client.Process(context.Background(), cmd)
	return err
}

// SetBloomValue 添加数据
func (b *BloomCache) SetBloomValue(key string, element string) error {
	cmd := redis.NewCmd(context.Background(), "BF.ADD", key, element)
	err := b.client.Process(context.Background(), cmd)
	return err
}

// CheckBloomValue 检查数据
func (b *BloomCache) CheckBloomValue(key string, element string) (bool, error) {
	cmd := redis.NewCmd(context.Background(), "BF.EXISTS", key, element)
	err := b.client.Process(context.Background(), cmd)
	if err != nil {
		return false, err
	}
	return cmd.Val() == 1, err
}

// 删除
func (b *BloomCache) removeBloomValue(key string, element string) error {
	cmd := redis.NewCmd(context.Background(), "BF.DEL", key, element)
	err := b.client.Process(context.Background(), cmd)
	return err
}
