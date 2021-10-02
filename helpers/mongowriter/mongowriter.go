package mongowriter

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoWriter struct {
	Session *mgo.Session
}

// type ConfigLogger struct {
// 	Format           interface{} `yaml:"format"`
// 	CustomTimeFormat string      `yaml:"custom_time_format"`
// 	Output           io.Writer
// }

// func (loggerConf *ConfigLogger) Init() middleware.LoggerConfig {
// 	loggerConf.Format = bson.UnmarshalJSON([]byte(`"method=${method}, uri=${uri}, status=${status}"`), &loggerConf.Format)
// 	return middleware.LoggerConfig{
// 		Format: loggerConf.Format,
// 	}
// }

func (mw *MongoWriter) Write(p []byte) (n int, err error) {
	origLen := len(p)
	// if len(p) > 0 && p[len(p)-1] == '\n' {
	// 	p = p[:len(p)-1]
	// }

	// var message interface{}
	// err = bson.UnmarshalJSON([]byte(p), &message)
	// if err != nil {
	// 	return
	// }

	// if p[0] != 'm' {
	// 	return
	// }

	c := mw.Session.DB("logger").C("log")
	res := string(p)
	if res[0] != 'm' {
		return
	}

	err = c.Insert(bson.M{
		"created": time.Now(),
		"msg":     res,
	})

	// err = c.Insert(message)
	if err != nil {
		return
	}
	return origLen, nil
}
