package logging

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(name string, includeConsole bool) zerolog.Logger {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	rotator := &lumberjack.Logger{
		Filename:   "logs/" + name + ".log",
		MaxSize:    5, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Compress:   true,
	}

	var writers []io.Writer
	if includeConsole {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: zerolog.TimeFieldFormat})
	}
	writers = append(writers, rotator)
	writer := zerolog.MultiLevelWriter(writers...)

	return zerolog.New(writer).With().Timestamp().Logger()
}
