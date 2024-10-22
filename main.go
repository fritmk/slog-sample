package main

import (
	"log/slog"
	"logging_sample/loggers"
	"os"
)

func main() {

	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: loggers.ReplaceOption,
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, &opts)
	loggers := slog.New(jsonHandler)

	loggers.Error("[Error] fear is the mind killer")
	loggers.Info("[info] fear is the mind killer")
	loggers.Debug("[debug] fear is the mind killer")

	request := slog.Group("request",
		"method", "get",
		"url", "/")

	loggers.Info("[Info] message1", "key1", "string1", "key2", 1, request)

}
