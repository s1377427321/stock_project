package llog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/leisc/llog/llager"
)

const (
	//DEBUG is a constant of string type
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

//Config is a struct which stores details for maintaining logs
type Config struct {
	LoggerLevel    string
	LoggerFile     string
	Writers        []string
	EnableRsyslog  bool
	RsyslogNetwork string
	RsyslogAddr    string

	LogFormatText bool
}

var config = DefaultConfig()
var m sync.RWMutex

//Writers is a map
var Writers = make(map[string]io.Writer)

//RegisterWriter is used to register a io writer
func RegisterWriter(name string, writer io.Writer) {
	m.Lock()
	Writers[name] = writer
	m.Unlock()
}

//DefaultConfig is a function which retuns config object with default configuration
func DefaultConfig() *Config {
	return &Config{
		LoggerLevel:    INFO,
		LoggerFile:     "",
		EnableRsyslog:  false,
		RsyslogNetwork: "udp",
		RsyslogAddr:    "127.0.0.1:5140",
		LogFormatText:  false,
	}
}

//Init is a function which initializes all config struct variables
func LagerInit(c Config) {
	if c.LoggerLevel != "" {
		config.LoggerLevel = c.LoggerLevel
	}

	if c.LoggerFile != "" {
		config.LoggerFile = c.LoggerFile
		config.Writers = append(config.Writers, "file")
	}

	if c.EnableRsyslog {
		config.EnableRsyslog = c.EnableRsyslog
	}

	if c.RsyslogNetwork != "" {
		config.RsyslogNetwork = c.RsyslogNetwork
	}

	if c.RsyslogAddr != "" {
		config.RsyslogAddr = c.RsyslogAddr
	}
	if len(c.Writers) == 0 {
		config.Writers = append(config.Writers, "stdout")

	} else {
		config.Writers = c.Writers
	}
	config.LogFormatText = c.LogFormatText
	RegisterWriter("stdout", os.Stdout)
	var file io.Writer
	var err error
	if config.LoggerFile != "" {
		file, err = os.OpenFile(config.LoggerFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

	}
	for _, sink := range config.Writers {
		if sink == "file" {
			if file == nil {
				log.Panic("Must set file path")
			}
			RegisterWriter("file", file)
		}
	}
}

//NewLogger is a function
func NewLogger(component string) llager.LLogger {
	return NewLoggerExt(component, component)
}

//NewLoggerExt is a function which is used to write new logs
func NewLoggerExt(component string, appGUID string) llager.LLogger {
	var lagerLogLevel llager.LogLevel
	switch strings.ToUpper(config.LoggerLevel) {
	case DEBUG:
		lagerLogLevel = llager.DEBUG
	case INFO:
		lagerLogLevel = llager.INFO
	case WARN:
		lagerLogLevel = llager.WARN
	case ERROR:
		lagerLogLevel = llager.ERROR
	case FATAL:
		lagerLogLevel = llager.FATAL
	default:
		panic(fmt.Errorf("unknown logger level: %s", config.LoggerLevel))
	}
	logger := llager.NewLoggerExt(component, config.LogFormatText)
	for _, sink := range config.Writers {

		writer, ok := Writers[sink]
		if !ok {
			log.Panic("Unknow writer: ", sink)
		}
		sink := llager.NewReconfigurableSink(llager.NewWriterSink(sink, writer, llager.DEBUG), lagerLogLevel)
		logger.RegisterSink(sink)
	}

	return logger
}

func Debug(action string, data ...llager.Data) {
	Logger.Debug(action, data...)
}

func Info(action string, data ...llager.Data) {
	Logger.Info(action, data...)
}

func Warn(action string, data ...llager.Data) {
	Logger.Warn(action, data...)
}

func Error(action string, err error, data ...llager.Data) {
	Logger.Error(action, err, data...)
}

func Fatal(action string, err error, data ...llager.Data) {
	Logger.Fatal(action, err, data...)
}

func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

func Errorf(err error, format string, args ...interface{}) {
	Logger.Errorf(err, format, args...)
}

func Fatalf(err error, format string, args ...interface{}) {
	Logger.Fatalf(err, format, args...)
}
