package log

import (
	"backend-engineering-challenge/internals/config"
	"backend-engineering-challenge/internals/domain"
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
)

var NativeLog *log.Logger

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
	NativeLog = log.New(os.Stdout, ``, log.LstdFlags|log.Lmicroseconds)
}

func paint(val string, color string) string {
	return fmt.Sprintf("%v%v%v", color, val, "\033[0m")
}

func logEntry(logType string, uuid string, prefix string, message interface{}, color string, params ...interface{}) {

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

		message = fmt.Sprintf(`[%+v on %s %d]`, message, file, line)
	} else {
		message = fmt.Sprintf(`[%+v]`, message)
	}

	if logType == fatal {
		NativeLog.Fatalln(format(color, uuid, prefix, message, params...))
	}

	if logType == err {
		NativeLog.Println(format(color, uuid, prefix, message, params...))
		return
	}
	NativeLog.Println(format(color, uuid, prefix, message, params...))

}

func format(typ string, uuid string, prefix string, message interface{}, params ...interface{}) string {
	str := fmt.Sprintf("%s", typ)

	if uuid != "" {
		str = fmt.Sprintf("%s [%s]", str, uuid)
	}

	str = fmt.Sprintf("%s [%s] %v", str, prefix, fmt.Sprintf("%+v", message))

	if len(params) != 0 {
		str = fmt.Sprintf("%s %+v", str, params)
	}

	return str
}

func ErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	logEntry(err, fmt.Sprintf("%+v", ctx.Value(domain.CorrelationIdContextKey)), prefix, message, logColors[err], params...)
}

func InfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	logEntry(info, fmt.Sprintf("%+v", ctx.Value(domain.CorrelationIdContextKey)), prefix, message, logColors[info], params...)
}

func FatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	logEntry(fatal, fmt.Sprintf("%+v", ctx.Value(domain.CorrelationIdContextKey)), prefix, message, logColors[fatal], params...)
}

func TraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	logEntry(trace, fmt.Sprintf("%+v", ctx.Value(domain.CorrelationIdContextKey)), prefix, message, logColors[trace], params...)
}
func DebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	logEntry(debug, fmt.Sprintf("%+v", ctx.Value("uuid")), prefix, message, logColors[debug], params...)
}
