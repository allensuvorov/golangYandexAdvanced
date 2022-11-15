package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Config struct {
	once sync.Once
	vals map[string]string
}

func (c *Config) Get(k string) (string, bool) {
	c.once.Do(func() {
		c.vals = map[string]string{
			"host": "127.0.0.1",
			"port": fmt.Sprintf("%d", rand.Intn(65535)),
		}
	})
	//func() {
	//	c.vals = map[string]string{
	//		"host": "127.0.0.1",
	//		"port": fmt.Sprintf("%d", rand.Intn(65535)),
	//	}
	//}()

	v, ok := c.vals[k]
	return v, ok
}

func main() {
	var cfg Config

	keys := []string{"host", "port", "port", "host", "port"}
	for _, k := range keys {
		go func(k string) {
			v, ok := cfg.Get(k)
			if !ok {
				return
			}
			fmt.Printf("%s = %s\n", k, v)
		}(k)
	}
	time.Sleep(1 * time.Second)
}
