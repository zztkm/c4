package main

import (
	"fmt"
	"os"
	"time"
	"unicode/utf8"

	"github.com/atotto/clipboard"
	"github.com/gen2brain/beeep"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func run() int {
	text, err := clipboard.ReadAll()
	if err != nil {
		logger.Error().Err(err).Send()
		return 1
	}

	// 無限ループしながらclipboardが更新されるのを待つ
	for {
		time.Sleep(1 * time.Second)
		now, err := clipboard.ReadAll()
		if err != nil {
			logger.Error().Err(err).Send()
			return 1
		}

		if len(now) == 0 {
			continue
		}

		if text == now {
			continue
		} else {
			len := utf8.RuneCountInString(now)
			msg := fmt.Sprintf("count: %d", len)
			logger.Info().Str("text", now).Int("count", len).Send()
			err := beeep.Notify("c4", msg, "")
			if err != nil {
				logger.Error().Err(err).Send()
				return 1
			}
			text = now
		}
	}
}

func main() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger = zerolog.New(consoleWriter)
	os.Exit(run())
}
