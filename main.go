package main

import (
	"bilibili-live/config"
	_ "bilibili-live/danmu"
	_ "bilibili-live/decode"
	"bilibili-live/live"
	_ "bilibili-live/monitor"
	"bilibili-live/routers"
	_ "bilibili-live/tools"

	"flag"
	"fmt"

	"github.com/kataras/golog"
)

var (
	port       string
	configName string
)

func init() {
	flag.StringVar(&port, "p", "9855", "端口号")
	flag.StringVar(&configName, "c", "config.yml", "配置文件")
	flag.Parse()
	config.ConfigFile = fmt.Sprintf("./%s", configName)
	c := config.New()
	err := c.LoadConfig()
	if err != nil {
		golog.Fatal(fmt.Sprintf("Load config error: %s", err))
	}
	for _, v := range c.Conf.Live {
		live.AddRoom(v.RoomID)
	}
	// go flushLiveStatus()
	live.StartTimingTask("Upload2BaiduPCS", c.Conf.RcConfig.NeedBdPan, c.Conf.RcConfig.UploadTime, live.Upload2BaiduPCS)
	live.StartTimingTask("CleanRecordingDir", c.Conf.RcConfig.NeedRegularClean, c.Conf.RcConfig.RegularCleanTime, live.CleanRecordingDir)
}

func main() {

	routers.GIN.Run(fmt.Sprintf(":%s", port))
}
