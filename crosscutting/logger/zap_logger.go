package logger

import (
	"log"

	cc "github.com/matheusrbarbosa/gofin/crosscutting"
	"go.uber.org/zap"
)

type ZapLogger struct {
	ctxLogger *zap.SugaredLogger
}

func (z ZapLogger) Error(args ...interface{}) {
	z.ctxLogger.Error(args...)
}

func (z ZapLogger) Errorf(template string, args ...interface{}) {
	z.ctxLogger.Errorf(template, args...)
}

func (z ZapLogger) Infoln(args ...interface{}) {
	z.ctxLogger.Infoln(args...)
}

func (z ZapLogger) Infof(template string, args ...interface{}) {
	z.ctxLogger.Infof(template, args...)
}

func (z ZapLogger) Panic(args ...interface{}) {
	z.ctxLogger.Panic(args...)
}

func (z ZapLogger) Panicln(args ...interface{}) {
	z.ctxLogger.Panicln(args...)
}

func (z ZapLogger) Fatal(args ...interface{}) {
	z.ctxLogger.Fatal(args...)
}

func (z ZapLogger) Fatalf(template string, args ...interface{}) {
	z.ctxLogger.Fatalf(template, args)
}

func newLogger() *zap.SugaredLogger {
	var logger *zap.Logger
	var err error

	if cc.IsProduction() {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	logger.Sync()
	return logger.Sugar()
}
