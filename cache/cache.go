package cache

import (
	"time"
)

type ChartCache struct {
	Value     int
	TimeStamp int
}

type ChartCaches struct {
	Caches []ChartCache
}

func InitChartCache() *ChartCaches {
	chartCaches := make([]ChartCache, 0)
	ticker := time.NewTicker(time.Second)
	c := &ChartCaches{
		Caches: chartCaches,
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				c.ttl()
			}
		}
	}()

	return c
}

func (c *ChartCaches) ttl() bool {
	cacheCount := len(c.Caches)
	if cacheCount > 100 {
		c.Caches = c.Caches[10:cacheCount]
	}
	return true
}

func (c *ChartCaches) Charts() *ChartCaches {
	return c
}

func (c *ChartCaches) SetChartCache(chartCache ChartCache) bool {
	caches := c.Caches
	count := len(caches)
	caches = append(caches, chartCache)
	c.Caches = caches
	if count < len(caches) {
		return true
	} else {
		return false
	}
}
