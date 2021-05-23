package configs

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tuanden0/user_api/models"
	"github.com/tuanden0/user_api/routes"
)

type AppConfig struct {
	ADDR    string
	LogName string
	DB_NAME string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		ADDR:    ":8000",
		LogName: "app.log",
		DB_NAME: "user.db",
	}
}

func LoadGinConfig(r *gin.Engine) *gin.Engine {

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("[%s] - %s\t\"%s\t%s\t%s\t%d\t%s\t%s\"\n",
			param.TimeStamp.Format(time.RFC1123),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))
	// Setup recovery
	r.Use(gin.Recovery())

	return r
}

func NewRoute() *gin.Engine {

	// Generate gin.Engine
	router := gin.New()

	// Load Gin Config
	router = LoadGinConfig(router)

	// Load User Routes
	router = routes.UserRoute(router)

	return router
}

func (a *AppConfig) GetLogFile() *os.File {

	// Get Log File
	var f *os.File
	if _, err := os.Stat(a.LogName); os.IsNotExist(err) {
		f, _ = os.Create(a.LogName)
	} else {
		f, _ = os.OpenFile(a.LogName, os.O_APPEND, 0666)
	}

	return f
}

func RunServer() {

	// Load AppConfig
	appConfig := NewAppConfig()

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(appConfig.GetLogFile(), os.Stdout)

	// Connect DB
	models.InitDatabase(appConfig.DB_NAME)

	// Add route to global route
	router := NewRoute()

	router.Run(appConfig.ADDR)

}
