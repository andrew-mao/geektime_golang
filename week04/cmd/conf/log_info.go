package conf

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type LogConf struct {
	runLogFile    *os.File
	accessLogFile *os.File
	RunOutput     string `toml:"output"`
	AccessOutput  string `toml:"accessoutput"`
}

//日志初始化
func (l *LogConf) LogInit() {
	var err error
	if l.runLogFile == nil {
		l.runLogFile, err = os.OpenFile(l.RunOutput, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	}
	if l.accessLogFile == nil {
		l.accessLogFile, err = os.OpenFile(l.RunOutput, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	}
}

//运行错误日志
func (l *LogConf) WriteRunOutPut(err error) {
	write(err, l.runLogFile)
}

//接入日志
func (l *LogConf) WriteAccessOutPut(err error) {
	write(err, l.accessLogFile)
}

//日志关闭
func (l *LogConf) LogClose() {
	defer func() {
		l.runLogFile = nil
		l.accessLogFile = nil
	}()
	defer l.runLogFile.Close()
	defer l.accessLogFile.Close()
}

//写入日志到相应文件
func write(err error, file *os.File) {
	if file == nil {
		return
	}
	file.WriteString(fmt.Sprintf("origin: %T \n %+v", errors.Cause(err), errors.Cause(err)))
}
