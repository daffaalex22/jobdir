package middlewares

// import (
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// 	"main.go/helpers/mongowriter"
// )

// type ConfigLogger struct {
// 	Format string
// }

// func (loggerConf *ConfigLogger) Init() middleware.LoggerConfig {
// 	return middleware.LoggerConfig{
// 		Format: loggerConf.Format,
// 	}
// }

// func (loggerConf *ConfigLogger) MongoOut(e *echo.Echo, mw *mongowriter.MongoWriter) {
// 	e.Logger.SetOutput(mw)
// }
