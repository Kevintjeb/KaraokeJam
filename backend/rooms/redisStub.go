package main

import (
	"errors"
	"time"

	"github.com/alicebob/miniredis"
)

type StubRedis struct {
	c *miniredis.Miniredis
}

func (r StubRedis) SRem(key string, members string) {
	r.c.SRem(key, members)
}
func (r StubRedis) Set(key string, value string, expiration time.Duration) {
	r.c.Set(key, value)
}
func (r StubRedis) SAdd(key string, members string) {
	r.c.SetAdd(key, members)
}
func (r StubRedis) Expire(key string, expiration time.Duration) {
}
func (r StubRedis) Del(keys string) {
	r.c.Get(keys)
}
func (r StubRedis) Get(key string) (string, error) {
	return r.c.Get(key)
}
func (r StubRedis) Exists(keys string) (int64, error) {
	exists := r.c.Exists(keys)

	if exists {

		return 0, nil
	}
	return 0, errors.New("error")
}
