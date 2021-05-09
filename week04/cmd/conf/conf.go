package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"mblog/dao"
)

const PATH = "cmd/conf.toml"

var (
	ConfPath string
	Conf     = &Config{}
)

type Config struct {
	Gin   *Gin
	MySQL *dao.Dao
	Log   *LogConf
}

//读取toml配置文件，完成配置参数初始化
func Init(path string) (err error) {
	if path == "" {
		path = PATH
	}
	_, err = toml.DecodeFile(path, &Conf)
	if err != nil {
		return errors.Wrap(err, "get config info failed from "+ConfPath)
	}
	return nil
}
