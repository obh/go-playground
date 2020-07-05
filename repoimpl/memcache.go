package repoimpl

import (
    "log"
    "github.com/bradfitz/gomemcache/memcache"

    "github.com/obh/go-playground/config"
)

type Cache struct {
    *memcache.Client
}

func InitCache(cfg config.CacheConfig) (*Cache, error) {
    log.Println("Connecting with Memcache..")
    mc := memcache.New(cfg.Host)
    return &Cache{Client: mc}, nil
}
