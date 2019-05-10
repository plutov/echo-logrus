package echologrus

import (
	"io"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// Logger : implement logrus Logger
type Logger struct {
	*logrus.Logger
}

// Level returns logger level
func (l Logger) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	default:
		l.Panic("Invalid level")
	}

	return log.OFF
}

// SetHeader is a stub to satisfy interface
// It's controlled by logrus
func (l Logger) SetHeader(_ string) {}

// SetPrefix It's controlled by logrus
func (l Logger) SetPrefix(s string) {}

// Prefix It's controlled by logrus
func (l Logger) Prefix() string {
	return ""
}

// SetLevel set level to logger from given log.Lvl
func (l Logger) SetLevel(lvl log.Lvl) {
	switch lvl {
	case log.DEBUG:
		logrus.SetLevel(logrus.DebugLevel)
	case log.WARN:
		logrus.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		logrus.SetLevel(logrus.ErrorLevel)
	case log.INFO:
		logrus.SetLevel(logrus.InfoLevel)
	default:
		l.Panic("Invalid level")
	}
}

// Output logger output func
func (l Logger) Output() io.Writer {
	return l.Out
}

// SetOutput change output, default os.Stdout
func (l Logger) SetOutput(w io.Writer) {
	logrus.SetOutput(w)
}

// Printj print json log
func (l Logger) Printj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Print()
}

// Debugj debug json log
func (l Logger) Debugj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Debug()
}

// Infoj info json log
func (l Logger) Infoj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Info()
}

// Warnj warning json log
func (l Logger) Warnj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Warn()
}

// Errorj error json log
func (l Logger) Errorj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Error()
}

// Fatalj fatal json log
func (l Logger) Fatalj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Fatal()
}

// Panicj panic json log
func (l Logger) Panicj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Panic()
}

// Print string log
func (l Logger) Print(i ...interface{}) {
	logrus.Print(i[0].(string))
}

// Debug string log
func (l Logger) Debug(i ...interface{}) {
	logrus.Debug(i[0].(string))
}

// Info string log
func (l Logger) Info(i ...interface{}) {
	logrus.Info(i[0].(string))
}

// Warn string log
func (l Logger) Warn(i ...interface{}) {
	logrus.Warn(i[0].(string))
}

// Error string log
func (l Logger) Error(i ...interface{}) {
	logrus.Error(i[0].(string))
}

// Fatal string log
func (l Logger) Fatal(i ...interface{}) {
	logrus.Fatal(i[0].(string))
}

// Panic string log
func (l Logger) Panic(i ...interface{}) {
	logrus.Panic(i[0].(string))
}

func logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()

	p := req.URL.Path

	bytesIn := req.Header.Get(echo.HeaderContentLength)

	logrus.WithFields(map[string]interface{}{
		"time_rfc3339":  time.Now().Format(time.RFC3339),
		"remote_ip":     c.RealIP(),
		"host":          req.Host,
		"uri":           req.RequestURI,
		"method":        req.Method,
		"path":          p,
		"referer":       req.Referer(),
		"user_agent":    req.UserAgent(),
		"status":        res.Status,
		"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		"latency_human": stop.Sub(start).String(),
		"bytes_in":      bytesIn,
		"bytes_out":     strconv.FormatInt(res.Size, 10),
	}).Info("Handled request")

	return nil
}

func logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return logrusMiddlewareHandler(c, next)
	}
}

// Hook is a function to process middleware.
func Hook() echo.MiddlewareFunc {
	return logger
}
