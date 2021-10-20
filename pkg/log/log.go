package log

import (
	"go-fast/pkg/config"
	"io"
	"os"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func InitLog(conf config.LogConfig) {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	// 是否输出到日志日志文件中
	if conf.LogToStdout {
		// 按日日期格式分割 .log_202110200000
		storePath := conf.LogBase + "_%Y%m%d%H%M"
		rl, err := rotatelogs.New(storePath,
			rotatelogs.WithRotationSize(conf.RollSize))
		if err != nil {
			panic(err)
		}
		mw := io.MultiWriter(os.Stdout, rl)
		logrus.SetOutput(mw)
		return
	}
	logrus.SetOutput(os.Stderr)
}
