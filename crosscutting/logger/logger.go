package logger

type Logger interface {
	Error(args ...interface{})
	Errorf(template string, args ...interface{})

	Infof(template string, args ...interface{})
	Infoln(args ...interface{})

	Panic(args ...interface{})
	Panicln(args ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

var logger ZapLogger

func init() {
	logger = ZapLogger{}
	logger.ctxLogger = newLogger()
}

func GetLogger() Logger {
	return logger
}
