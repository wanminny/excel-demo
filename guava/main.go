package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/goburrow/cache"
	"log"
)

func main() {
	load := func(k cache.Key) (cache.Value, error) {
		log.Println(11111)
		return fmt.Sprintf("%d", k), nil
	}
	// Create a new cache
	c := cache.NewLoadingCache(load,
		cache.WithMaximumSize(1000),
		cache.WithExpireAfterAccess(5*time.Second),
		cache.WithRefreshAfterWrite(20*time.Second),
	)


	getTicker := time.Tick(10 * time.Millisecond)
	reportTicker := time.Tick(1 * time.Second)
	for {
		select {
		case <-getTicker:
			//_, _ = c.Get(rand.Intn(2000))
			_, _ = c.Get(rand.Intn(22))

		case <-reportTicker:
			st := cache.Stats{}
			c.Stats(&st)
			fmt.Printf("%+v\n", st)
		}
	}
}