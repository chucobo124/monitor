package main

import (
	"fmt"
	"github.com/monitor/cache"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/platforms/firmata"
	"time"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/tty.usbmodem14611")
	sensor := aio.NewAnalogSensorDriver(firmataAdaptor, "0")
	chartCache := cache.InitChartCache()

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {

			chartCache.SetChartCache(cache.ChartCache{
				Value:     int64(data.(int)),
				TimeStamp: time.Now().Unix(),
			})

		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor},
		work,
	)

	robot.Start()
}
