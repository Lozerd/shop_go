package logging

import (
	"fmt"
	"log"
	"os"
    "github.com/rs/zerolog"
)

var Logger map[string]*zerolog.Logger

func CreateLogger(filename string) *zerolog.Logger {
	if filename == "" {
		log.Panic("Filename can't be empty")
	}

    if _, err := os.Stat("logs"); os.IsNotExist(err) {
        os.Mkdir("logs", os.ModePerm)
    }

	file, err := os.OpenFile("logs/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
    defer file.Close()
	if err != nil {
		log.Panic(fmt.Sprintf("Couldn't open file at logs/%s", filename))
	}

    logger := zerolog.New(file).With().Timestamp().Logger()
    return &logger
}

func InitLogging() {
    Logger = make(map[string]*zerolog.Logger)
    Logger["main"] = CreateLogger("main.log")
    Logger["request"] = CreateLogger("request.log")
}
