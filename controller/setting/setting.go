package setting

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/Unknwon/goconfig"
)

var (
	Cfg *goconfig.ConfigFile
)

var (
	ConfPath = "../conf/config.ini"
)

var (
	DriverName string
	DataSource string
	MaxIdle    int
	MaxOpen    int
	DebugLog   bool
)

// LoadConfig loads configuration file.
func LoadConfig() *goconfig.ConfigFile {
	var err error
	// Load configuration, set app version and log level.
	Cfg, err = goconfig.LoadConfigFile(ConfPath)
	if Cfg == nil {
		fmt.Println("Fail to load configuration file: " + err.Error())
		log.Infof("error:", err.Error())
		os.Exit(2)
	}
	Cfg.BlockMode = false
	DriverName = Cfg.MustValue("orm", "driver_name", "mysql")
	DataSource = Cfg.MustValue("orm", "data_source", "root:123456@tcp(127.0.0.1:3306)/helloweb?charset=utf8")
	MaxIdle = Cfg.MustInt("orm", "max_idle_conn", 30)
	MaxOpen = Cfg.MustInt("orm", "max_open_conn", 50)
	DebugLog = Cfg.MustBool("orm", "debug_log", false)
	return Cfg
}
