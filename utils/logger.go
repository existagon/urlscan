package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger = &log.Logger{
	Out:   os.Stdout,
	Level: log.DebugLevel,
	Formatter: &log.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	},
}
