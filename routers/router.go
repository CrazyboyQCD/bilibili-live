package routers

import (
	"io"
	"net/http"
	"os"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"bilibili-live/controllers"
)

var GIN *gin.Engine

var sugarLogger *zap.SugaredLogger

func InitLogger() {

	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// zap.AddCaller()  添加将调用函数信息记录到日志中的功能。
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器

	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

func init() {
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.Discard
	GIN = gin.Default()
	GIN.LoadHTMLGlob("dist/*.html")        // 添加入口index.html
	GIN.LoadHTMLFiles("static/*/*")        // 添加资源路径
	GIN.Static("/static", "./dist/static") // 添加资源路径
	GIN.StaticFile("/", "dist/index.html") // 前端接口
	GIN.StaticFile("/favicon.ico", "dist/favicon.ico")
	GIN.StaticFile("/js", "dist/js")
	GIN.StaticFile("/fonts", "dist/fonts")
	GIN.StaticFile("/css", "dist/css")
	GIN.Use(Cors())
	GIN.Use(ginzap.Ginzap(zap.L(), "2006/01/02 15:04:05", true))
	GIN.Use(ginzap.RecoveryWithZap(zap.L(), true))
	GIN.GET("/basestatus", controllers.GetBaseStatus)
	GIN.GET("/livestatus", controllers.GetLiveStatus)
	GIN.GET("/areainfos", controllers.GetAreaInfos)
	GIN.GET("/anchorlivebacklist/:name", controllers.GetAnchorLivebackList)
	GIN.POST("/blockroom", controllers.ProcessBlockRoom)
	GIN.POST("/decode", controllers.ProcessDecode)
	GIN.POST("/roomhandle", controllers.RoomHandle)
	GIN.POST("/areahandle", controllers.AreaHandle)
	GIN.POST("livebackstatistics", controllers.GetLivebackStatistics)
	GIN.POST("/getliveurl", controllers.GetRoomLiveURL)
	GIN.POST("/getwordcloud", controllers.GetWordCloud)
	GIN.GET("/refresh/:roomID", controllers.RefreshRoomInfo)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session, Access-Control-Allow-Methods, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Credentials, Access-Control-Allow-Headers")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		c.Next()
	}
}
