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

func main() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger = zerolog.New(consoleWriter)

	text, err := clipboard.ReadAll()
	if err != nil {
		logger.Error().Err(err).Send()
		os.Exit(1)
	}

	// 無限ループしながらclipboardが更新されるのを待つ
	for {
		now, err := clipboard.ReadAll()
		if err != nil {
			logger.Error().Err(err).Send()
			os.Exit(1)
		}

		if text == now {
			continue
		} else {
			len := utf8.RuneCountInString(now)
			msg := fmt.Sprintf("count: %d", len)
			logger.Info().Str("text", now).Int("count", len).Send()
			err := beeep.Notify("c4", msg, "")
			if err != nil {
				panic(err)
			}
			text = now
		}
		time.Sleep(1 * time.Second)
	}
}
