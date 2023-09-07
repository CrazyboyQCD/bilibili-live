package controllers

import (
	"bilibili-live/live"
	"bilibili-live/tools"

	"github.com/gin-gonic/gin"
)

func GetBaseStatus(c *gin.Context) {
	c.JSON(200, struct {
		TotalDownload int64 `json:"totalDownload"`
		FileNum       int64 `json:"fileNum"`
		// DeviceInfo		tools.DeviceInfo	`json:"deviceInfo"`
	}{
		tools.DirSize("./recording", 0),
		tools.CacRecordingFileNum(),
		// tools.GetDeviceInfo(),
	})
}

func GetLiveStatus(c *gin.Context) {
	c.JSON(200, live.Lives)
}
