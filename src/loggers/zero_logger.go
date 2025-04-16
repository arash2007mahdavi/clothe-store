package loggers

import (
	"os"
	"store/src/configs"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type zeroLogger struct {
	cfg    *configs.Config
	logger *zerolog.Logger
}

func NewZeroLogger(cfg *configs.Config) *zeroLogger {
	logger := &zeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

var ZeroLogLevels = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	level, exist := ZeroLogLevels[l.cfg.Logging.LogLevel]
	if exist {
		return level
	}
	return zerolog.DebugLevel
}

func (l *zeroLogger) Init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(l.getLogLevel())

	file, err := os.OpenFile(l.cfg.Logging.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(file).
		With().
		Timestamp().
		Str("AppName", "Store").
		Str("LoggerName", "Zerolog").
		Logger()

	l.logger = &logger
}

func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := getParams(extra)

	l.logger.Debug().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}
func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args...)
}

func (l *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := getParams(extra)

	l.logger.Info().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}
func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}

func (l *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := getParams(extra)

	l.logger.Warn().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}
func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}

func (l *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := getParams(extra)

	l.logger.Error().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}
func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args...)
}

func (l *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := getParams(extra)

	l.logger.Fatal().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}
func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}

func getParams(extra map[ExtraKey]interface{}) map[string]interface{} {
	params := make(map[string]interface{})

	for k, v := range extra {
		params[string(k)] = v
	}
	return params
}