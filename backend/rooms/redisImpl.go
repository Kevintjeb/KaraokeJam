package main

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisImpl struct {
	c *redis.Client
}

func (r RedisImpl) SRem(key string, members string){
	r.c.SRem(key, members)
}
func (r RedisImpl) Set(key string, value string, expiration time.Duration){
	r.c.Set(key, value, expiration)
}
func (r RedisImpl) SAdd(key string, members string){
	r.c.SAdd(key, members)
}
func (r RedisImpl) Expire(key string, expiration time.Duration){
	r.c.Expire(key, expiration)
}
func (r RedisImpl) Del(keys string){
	r.c.Get(keys)
}
func (r RedisImpl) Get(key string)  (string, error){
	return r.c.Get(key).Result()
}
func (r RedisImpl) Exists(keys string) (int64, error) {
	return r.c.Exists(keys).Result()
}

