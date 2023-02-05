// @author cold bin
// @date 2023/2/4

package conf

import (
	"eliminate-male-appearance-anxiety/global"
	"fmt"
	"github.com/cold-bin/goutil/general/mlog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConf() error {
	viper.SetConfigFile("./conf/conf.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&global.Set); err != nil {
		return err
	}

	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		mlog.Warnf("配置文件被修改：%s", in.String())
		if err := viper.Unmarshal(&global.Set); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	return nil
}
