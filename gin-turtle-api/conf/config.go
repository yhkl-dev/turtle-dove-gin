package conf

import (
	"fmt"

	"github.com/Unknwon/goconfig"
	"github.com/yhkl-dev/turtle-dove-gin/gin-turtle-api/utils/logx"
)

var (
	SysConfig *goconfig.ConfigFile
)

// InitConfig init config file
func init() {
	tutleDove := `
  _______         _   _      _____                 
 |__   __|       | | | |    |  __ \                
    | |_   _ _ __| |_| | ___| |  | | _____   _____ 
    | | | | | '__| __| |/ _ \ |  | |/ _ \ \ / / _ \
    | | |_| | |  | |_| |  __/ |__| | (_) \ V /  __/
    |_|\__,_|_|   \__|_|\___|_____/ \___/ \_/ \___|
	`
	fmt.Println(tutleDove)
	var err error
	SysConfig, err = goconfig.LoadConfigFile("./conf/app.conf")
	if err != nil {
		panic(err)
	}

	logLevel, err := SysConfig.GetValue("Log", "log_level")
	if err != nil {
		logLevel = "debug"
	}

	logx.BuildLogger(logLevel)
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		logx.Log().Panic("翻译文件加载失败", err)
	}
}
