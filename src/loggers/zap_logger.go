package loggers

import (
	"store/src/configs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var ZapLogLevels = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type zapLogger struct {
	cfg *configs.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *configs.Config) *zapLogger {
	l := &zapLogger{cfg: cfg}
	l.Init()
	return l
}

func (l *zapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: l.cfg.Logging.Path,
		MaxSize: 1,
		MaxAge: 10,
		LocalTime: true,
		MaxBackups: 10,
		Compress: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.getLogLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
	logger = logger.With("AppName", "Store", "LoggerName", "ZapLogger")

	l.logger = logger
}

func (l *zapLogger) getLogLevel() zapcore.Level {
	cfg := configs.GetConfig()
	level, exist := ZapLogLevels[cfg.Logging.LogLevel]
	if !exist {
		return zapcore.DebugLevel
	} 
	return level
}

func AddToExtra(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	param := MapToZapParams(extra)
	return param
}

func MapToZapParams(extra map[ExtraKey]interface{}) []interface{} {
	var param []interface{}
	for k, v := range extra {
		param = append(param, string(k))
		param = append(param, v)
	}
	return param
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	param := AddToExtra(cat, sub, extra)
	l.logger.Debugw(msg, param...)
}
func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	param := AddToExtra(cat, sub, extra)
	l.logger.Infow(msg, param...)
}
func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	param := AddToExtra(cat, sub, extra)
	l.logger.Warnw(msg, param...)
}
func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	param := AddToExtra(cat, sub, extra)
	l.logger.Errorw(msg, param...)
}
func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	param := AddToExtra(cat, sub, extra)
	l.logger.Fatalw(msg, param...)
}
func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}