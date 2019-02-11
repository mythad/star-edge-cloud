package lsyslog

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/hashicorp/go-syslog"
	"github.com/juju/loggo"
)

type Formatter func(loggo.Entry) string

// Format returns the parameters separated by spaces except for filename and
// line which are separated by a colon.
func DefaultFormatter(entry loggo.Entry) string {

	// Just get the basename from the filename
	filename := filepath.Base(entry.Filename)
	return fmt.Sprintf("%s %s:%d %s", entry.Module, filename, entry.Line, entry.Message)
}

type syslogWriter struct {
	syslogger gsyslog.Syslogger
	Formatter Formatter
}

// NewSyslogWriter returns a new writer that writes
// log messages to syslog in a simple format tailored for syslog
func NewSyslogWriter(p gsyslog.Priority, facility, tag string) loggo.Writer {
	syslogger, err := gsyslog.NewLogger(p, facility, tag)
	if err != nil {
		panic(err)
	}
	slw := &syslogWriter{syslogger, DefaultFormatter}
	return slw
}

// NewDefaultSyslogWriter returns a new writer that writes
// log messages to syslog in a simple format tailored for syslog.
// Note this defaults to using LOCAL7.
func NewDefaultSyslogWriter(level loggo.Level, tag, facility string) loggo.Writer {
	if facility == "" {
		facility = "LOCAL7"
	}
	var syslogger gsyslog.Syslogger
	var err error
	// retry opening the syslog device as this may fail under load or in cases where the
	// process is forked.
	for i := 0; i < 3; i++ {
		syslogger, err = gsyslog.NewLogger(convertLevel(level), facility, tag)
		if err == nil {
			continue
		}
		// back off for a second
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		panic(err) // of course this will never happen
	}
	return &syslogWriter{syslogger, DefaultFormatter}
}

func (slog *syslogWriter) Write(entry loggo.Entry) {
	logLine := slog.Formatter(entry)
	slog.syslogger.WriteLevel(convertLevel(entry.Level), []byte(logLine))
}

func convertLevel(level loggo.Level) gsyslog.Priority {
	switch level {
	case loggo.DEBUG:
		return gsyslog.LOG_DEBUG
	case loggo.INFO:
		return gsyslog.LOG_INFO
	case loggo.WARNING:
		return gsyslog.LOG_WARNING
	case loggo.CRITICAL:
		return gsyslog.LOG_CRIT
	case loggo.ERROR:
		return gsyslog.LOG_ERR
	case loggo.TRACE:
		return gsyslog.LOG_DEBUG
	default:
		return gsyslog.LOG_DEBUG
	}

}
