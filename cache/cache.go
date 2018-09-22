package cache

import (
//"time"
)

type ChartCache struct {
	Value     int64
	TimeStamp int64
}

type ChartCaches struct {
	Caches []ChartCache
}

var ttl = 1000

func InitChartCache() *ChartCaches {
	chartCaches := make([]ChartCache, 0)
	//timer := time.NewTimer(time.Second)
	//go func() {
	//	<-timer.C
	//	println("Timer expired")
	//}()

	return &ChartCaches{
		Caches: chartCaches,
	}

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
