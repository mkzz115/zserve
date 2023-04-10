package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Error(format string, a ...interface{})
}

var l *logger

type logger struct {
	*logrus.Logger
}

func Init(path string) error {
	l = &logger{logrus.New()}
	l.Formatter = new(logrus.TextFormatter)
	l.Formatter.(*logrus.TextFormatter).DisableColors = true
	l.Level = logrus.InfoLevel
	//isExist, err := PathExists(path)
	//if err != nil {
	//	return err
	//}
	//if !isExist {
	//	err := os.Mkdir(filepath.Dir(path), 0755)
	//	if err != nil {
	//		return err
	//	}
	//}

	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(5 * 24)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	//logPath := path + ".log"
	//file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	return err
	//}
	//l.Out = file
	l.SetOutput(writer)
	return nil
}

func SetLevel(level int) {
	l.SetLevel(logrus.Level(level))
}

func Debug(format string, a ...interface{}) {
	l.Debugf(format, a...)
}

func Info(format string, a ...interface{}) {
	//fmt.Printf(format, a...)
	l.Infof(format, a...)
}

func Warn(format string, a ...interface{}) {
	//fmt.Printf(format, a...)
	l.Warnf(format, a...)
}

func Error(format string, a ...interface{}) {
	//fmt.Printf(format, a...)
	l.Errorf(format, a...)
}
