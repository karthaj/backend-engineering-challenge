package logger

import (
	"backend-engineering-challenge/internals/config"
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
)

var nativeLog *log.Logger

var (
	fatal = `FATAL`
	err   = `ERROR`
	debug = `DEBUG`
	trace = `TRACE`
	info  = `INFO`
)

var logColors = map[string]string{
	err:   paint(`[ ERROR ]`, "\033[0;31m"),
	info:  paint(`[ INFO  ]`, "\033[0;34m"),
	trace: paint(`[ TRACE ]`, "\033[0;33m"),
	debug: paint(`[ DEBUG ]`, "\033[0;35m"),
	fatal: "[0;31m[ FATAL ]\033[0;31m",
}

func Init() {
	nativeLog = log.New(os.Stdout, ``, log.LstdFlags|log.Lmicroseconds)
}

func paint(val string, color string) string {
	return fmt.Sprintf("%v%v%v", color, val, "\033[0m")
}

func logEntry(logType string, uuid string, message interface{}, color string, params ...interface{}) {

	var file string
	var line int
	if config.AppConf.Debug {
		_, f, l, ok := runtime.Caller(2)
		if !ok {
			f = `<Unknown>`
			l = 1
		}

		file = f
		line = l

		message = fmt.Sprintf(`[%s] [%+v on %s %d]`, uuid, message, file, line)
	} else {
		message = fmt.Sprintf(`[%s] [%+v]`, uuid, message)
	}

	if logType == fatal {
		nativeLog.Fatalln(format(color, message, params...))
	}

	if logType == err {
		nativeLog.Println(format(color, message, params...))
		return
	}
	nativeLog.Println(format(color, message, params...))

}

func format(typ string, message interface{}, params ...interface{}) string {

	var messageFmt = "%s %s %v"

	return fmt.Sprintf(messageFmt,
		typ,
		fmt.Sprintf("%+v", message),
		fmt.Sprintf("%+v", params))
}

func ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntry(err, fmt.Sprintf("%+v", ctx.Value("uuid")), message, logColors[err], params...)
}

func InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntry(info, fmt.Sprintf("%+v", ctx.Value("uuid")), message, logColors[info], params...)
}

func FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntry(fatal, fmt.Sprintf("%+v", ctx.Value("uuid")), message, logColors[fatal], params...)
}

func TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntry(trace, fmt.Sprintf("%+v", ctx.Value("uuid")), message, logColors[trace], params...)
}
func DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntry(debug, fmt.Sprintf("%+v", ctx.Value("uuid")), message, logColors[debug], params...)
}
