package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/juju/loggo"
	"github.com/ninjasphere/go-ninja/config"
	"github.com/wolfeidau/bugsnag-go"
	"github.com/wolfeidau/loggo-syslog"
)

// Logger wrapper for the internal logger with some extra helpers
type Logger struct {
	loggo.Logger
}

func init() {
	var level loggo.Level
	useSyslog := true

	if _, logErr := os.Stat("/dev/log"); logErr != nil {
		useSyslog = false
	}

	// snappy doesn't support syslog, so when we detect a snappy app we will disable syslog
	// and leave the default stderr logger (which works because snappy uses systemd)
	// eventually, the error case of syslog not existing should be exposed by loggo-syslog instead

	if os.Getenv("SNAPP_APP_PATH") != "" {
		useSyslog = false
	}

	debug := os.Getenv("DEBUG")
	if debug != "" {
		switch debug {
		case "INFO":
			level = loggo.INFO
		case "WARNING":
			level = loggo.WARNING
		case "ERROR":
			level = loggo.ERROR
		default:
			level = loggo.DEBUG
		}
	} else {
		// set the default level
		level = loggo.INFO

		// kill stderr
		log.SetOutput(ioutil.Discard)

		if useSyslog {
			// remove the default writer
			loggo.RemoveWriter("default")
		}
	}
	loggo.GetLogger("").SetLogLevel(level)
	if level != loggo.INFO {
		loggo.GetLogger("").Infof("Root logger initialized at level %v", level)
	}
	// setup the syslog writer
	if useSyslog {
		loggo.RegisterWriter("syslog", lsyslog.NewDefaultSyslogWriter(loggo.TRACE, path.Base(os.Args[0]), "LOCAL7"), loggo.TRACE)
	}

}

// BugsnagLogger used in bugsnag to ensure panics are written to the logger as well as bugsnag
type BugsnagLogger struct {
	loggo.Logger
}

// Printf used in bugsnag to ensure panics are written to the logger as well as bugsnag
func (lw *BugsnagLogger) Printf(format string, v ...interface{}) {
	lw.Logger.Warningf(format, v...)
}

// GetBugsnagLogger builds a wrapper for loggo which can be used by bugsnag.
func GetBugsnagLogger(name string) *BugsnagLogger {
	l := loggo.GetLogger(name)
	return &BugsnagLogger{l}
}

// GetLogger builds a ninja logger with the given name
func GetLogger(name string) *Logger {
	l := loggo.GetLogger(name)
	return &Logger{l}
}

// HandleError This notifies bugsnag and logs the error.
func (l *Logger) HandleError(err error, msg string) {
	l.Errorf("%s : %v", msg, err)
	// config.GetAll(true)
	bugsnag.Notify(err, bugsnag.MetaData{
		"SphereConfig": config.GetAll(true),
	})
}

// FatalError This notifies bugsnag and logs the error then quits.
func (l *Logger) FatalError(err error, msg string) {
	l.Errorf("%s : %v", msg, err)
	bugsnag.Notify(err, bugsnag.MetaData{
		"SphereConfig": config.GetAll(true),
	})

	os.Exit(1)
}

// HandleErrorf This notifies bugsnag and logs the error based on the args.
func (l *Logger) HandleErrorf(err error, msg string, args ...interface{}) {
	l.Errorf(msg, args)
	bugsnag.Notify(err, bugsnag.MetaData{
		"SphereConfig": config.GetAll(true),
	})
}

// FatalErrorf This notifies bugsnag and logs the error based on the args then quits
func (l *Logger) FatalErrorf(err error, msg string, args ...interface{}) {
	l.Errorf(msg, args)
	bugsnag.Notify(err, bugsnag.MetaData{
		"SphereConfig": config.GetAll(true),
	})
	os.Exit(1)
}

// Fatalf This notifies bugsnag and logs the error based on the args then quits
func (l *Logger) Fatalf(msg string, args ...interface{}) {
	l.Errorf(msg, args)
	bugsnag.Notify(fmt.Errorf(msg, args), bugsnag.MetaData{
		"SphereConfig": config.GetAll(true),
	})
	os.Exit(1)
}
