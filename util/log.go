package util

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func InitLog() {
	// 必须添加这句，否则还是会成为debug level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if os.Getenv("debug") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// 显示文件和行号
	zerolog.CallerMarshalFunc = func(file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006/01/02 15:04:05"}).
		With().Caller().Logger()
}
