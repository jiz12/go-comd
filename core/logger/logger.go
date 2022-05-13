package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type LevelMode string

const (
	LevelWarn  LevelMode = "warn"
	LevelInfo  LevelMode = "info"
	LevelDebug LevelMode = "debug"
	LevelTrace LevelMode = "trace"
)

type Config struct {
	LoggerConfig *LogConfig
}

type LogConfig struct {
	Level string
	Path  *string
}

var Logger *logrus.Logger

// InitLogger init logger
func InitLogger(configName string, configPaths []string) error {
	config, err := readConfig(configName, configPaths)
	if err != nil {
		return nil
	}

	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2018-01-02 15:04:05",
	})

	// set log level
	levelMode := LevelMode(config.LoggerConfig.Level)
	Logger.SetLevel(levelMode.Level())

	// print log to file
	if config.LoggerConfig.Path != nil {
		logfile, _ := os.OpenFile(*config.LoggerConfig.Path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		Logger.SetOutput(logfile)

	}
	Logger.SetReportCaller(true)
	return nil
}

// readConfig read config from config.yml
func readConfig(configName string, configPaths []string) (*Config, error) {

	var config Config
	vp := viper.New()
	vp.SetConfigName(configName)

	for _, configPath := range configPaths {
		vp.AddConfigPath(configPath)
	}

	if err := vp.ReadInConfig(); err != nil {
		return &config, err
	}

	err := vp.Unmarshal(&config)
	if err != nil {
		return &config, err
	}

	return &config, nil
}

func (l LevelMode) Level() logrus.Level {

	switch l {
	case LevelWarn:
		return logrus.WarnLevel
	case LevelInfo:
		return logrus.InfoLevel
	case LevelDebug:
		return logrus.DebugLevel
	case LevelTrace:
		return logrus.TraceLevel
	}
	return logrus.WarnLevel

}
