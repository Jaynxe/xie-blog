package core

// 同时实现时间和日志水平分割
import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/sirupsen/logrus"
)

var logLevels = []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel}

type logLevelWriter struct {
	files    map[logrus.Level]*os.File
	logPath  string
	fileDate string
}
type LogFormatter struct {
}

func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timeStamp := time.Now().Local().Format("2006-01-02 15:04:05")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	var msg string
	if global.GVB_CONFIG.Logger.ShowLine {
		msg = fmt.Sprintf("%s [%s] %s [%s:%d] %s \n", global.GVB_CONFIG.Logger.Prefix, strings.ToUpper(entry.Level.String()),
			timeStamp, file, len, entry.Message)
	} else {
		msg = fmt.Sprintf("%s [%s] %s %s \n", global.GVB_CONFIG.Logger.Prefix, strings.ToUpper(entry.Level.String()),
			timeStamp, entry.Message)
	}
	return []byte(msg), nil
}

// Levels 实现 Hook 接口的 Levels 方法
func (hook *logLevelWriter) Levels() []logrus.Level {
	// 表示在哪些日志级别才会会触发hook
	return logrus.AllLevels
}

// Fire 实现 Hook 接口的 Fire 方法
func (p *logLevelWriter) Fire(entry *logrus.Entry) error {
	if p == nil {
		return errors.New("LogFileWriter is nil")
	}
	// 获取日志级别
	level := entry.Level
	// 根据日志级别选择对应的文件写入器
	file, ok := p.files[level]
	if !ok {
		return fmt.Errorf("no file for log level %s", level)
	}
	// 判断是否需要切换日期
	fileDate := time.Now().Format("2006-01-02")
	if p.fileDate != fileDate {
		for _, file := range p.files {
			file.Close()
		}
		err := os.MkdirAll(fmt.Sprintf("%s/%s", p.logPath, fileDate), os.ModePerm)
		if err != nil {
			return err
		}
		for _, level := range logLevels {
			filename := fmt.Sprintf("%s/%s/%s.log", p.logPath, fileDate, strings.ToLower(level.String()))
			file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
			if err != nil {
				return err
			}
			p.files[level] = file
		}

	}
	// 写入日志数据到选定的文件
	data, _ := entry.String()
	file.Write([]byte(data))
	return nil
}

func InitLog() {
	log := logrus.New()

	// 创建不同级别的日志文件，并初始化写入器
	writers := make(map[logrus.Level]*os.File)

	for _, level := range logLevels {
		fileDate := time.Now().Format("2006-01-02")
		err := os.MkdirAll(fmt.Sprintf("%s/%s", global.GVB_CONFIG.Logger.Director, fileDate), os.ModePerm)
		if err != nil {
			log.Error(err)
			return
		}
		fileName := fmt.Sprintf("%s/%s/%s.log", global.GVB_CONFIG.Logger.Director, fileDate, strings.ToLower(level.String()))
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log.Error(err)
			return
		}
		writers[level] = file
	}

	// 创建日志文件写入器
	fileWriter := &logLevelWriter{
		files:    writers,
		logPath:  global.GVB_CONFIG.Logger.Director,
		fileDate: time.Now().Format("2006-01-02"),
	}

	// 添加日志钩子
	log.AddHook(fileWriter)
	// 设置报告调用者
	log.SetReportCaller(true)
	// 设置格式化器
	log.SetFormatter(&LogFormatter{})

	// 设置日志级别
	level, _ := logrus.ParseLevel(global.GVB_CONFIG.Logger.Level)
	log.SetLevel(level)
	// 是否禁用控制台输出
	if !global.GVB_CONFIG.Logger.LogInConsole {
		log.SetOutput(io.Discard)
	}
	global.GVB_LOGGER = log
}
