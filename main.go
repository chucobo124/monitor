package main

import (
	"github.com/gin-gonic/gin"
	"github.com/monitor/cache"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/platforms/firmata"
	"io"
	"time"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/tty.usbmodem14611")
	sensor := aio.NewAnalogSensorDriver(firmataAdaptor, "0")
	chartCache := cache.InitChartCache()

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			time.Sleep(1000 * time.Millisecond)
			chartCache.SetChartCache(cache.ChartCache{
				Value:     int(data.(int)),
				TimeStamp: int(time.Now().Unix()),
			})
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor},
		work,
	)

	go robot.Start()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/charts", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		dataQueue := make(chan interface{})
		go func() {
			dataQueue <- chartCache.Charts()
		}()
		c.Stream(func(w io.Writer) bool {
			c.SSEvent("message", <-dataQueue)
			return true
		})
	})
	r.Run()
}
