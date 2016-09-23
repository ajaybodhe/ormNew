package conf

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/scalingdata/gcfg"

)

const (
	gconfFile = "conf/ormNew.conf"
)

/* CFillerConfig stores the global configuration structure for cache filler */
var OrmNewConfig struct {
	DB struct {
		Host     string
		Port     int
		Username string
		Password string
		Protocol string
		DB       string
		ConnID   string
		MaxIdle  int
		MaxConn  int
	}
	NSQ struct {
		Host     string
		Port     string
	}
	Consumer struct {
		ChannelName string
	}
	Queue struct {
		Topic string
	}
}

const allowAllFilesCommand = "allowAllFiles=true"
const charsetUTF = "charset=utf8"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	ReadConfig(gconfFile, &OrmNewConfig)
	OrmNewConfig.DB.ConnID = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s&%s", OrmNewConfig.DB.Username, OrmNewConfig.DB.Password, OrmNewConfig.DB.Protocol,
		OrmNewConfig.DB.Host, OrmNewConfig.DB.Port, OrmNewConfig.DB.DB, allowAllFilesCommand, charsetUTF)
}

/*ReadConfig - reads the flags for --conf and if its found reads file and sets configuration into out. If --conf is not provided, then defaultPath is used. */
func ReadConfig(defaultPath string, out interface{}) {
	confFile := flag.String("conf", defaultPath, "Configuration file path")
	flag.Parse()
	glog.Info("conffile:", *confFile)
	err := gcfg.ReadFileInto(out, *confFile)
	if err != nil {
		glog.Fatal("error: util.conf.init:", err.Error())
	}
	glog.Info(os.Stdout, "boot.util.conf.init.success:\n***************Configuration:***************\n%+v\n*****************END****************\n", out)
}
