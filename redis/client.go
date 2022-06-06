package gredis

import (
	"github.com/go-redis/redis"
	"github.com/troyeye/glib/utils/cast"
	"strings"
	"sync"
)

type Client struct {
	c *redis.Client
}


var (
	clients = map[string]*Client{}
	mutex = sync.RWMutex{}
)

func getClient(url string) (*redis.Client, error) {
	mutex.RLock()
	cli, ok := clients[url]
	mutex.RUnlock()
	if ok {
		return cli.c, nil
	}

	mutex.Lock()
	defer mutex.Unlock()
	cli, ok = clients[url]
	if ok {
		return cli.c, nil
	}

	cli, err := newClient(url)
	if err != nil {
		return nil, err
	}
	clients[url] = cli
	return cli.c, nil
}

func newClient(url string) (*Client, error) {
	password, host, db := parseURL(url)
	redisOption := &redis.Options{
		Addr:     host,
		DB:       db,
		Password: password,
	}
	client := redis.NewClient(redisOption)
	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return &Client{c: client}, nil
}

func parseURL(url string) (string, string, int) {
	noHead := parseHeader(url)
	pass, noPass := parsePassword(noHead)
	host, db := parseHost(noPass)
	return pass, host, db
}

func parseHeader(url string) string {
	return strings.TrimPrefix(url, "redis://")
}

func parsePassword(url string) (password, noPassword string) {
	idx := strings.LastIndex(url, "@")
	if idx == -1 {
		return "", url
	}
	return url[:idx], url[idx+1:]
}

func parseHost(url string) (host string, db int) {
	idx := strings.LastIndex(url, "/")
	if idx == -1 {
		return url, 0
	}
	host = url[:idx]
	dbStr := url[idx+1:]
	db = cast.ToInt(dbStr)
	return
}
