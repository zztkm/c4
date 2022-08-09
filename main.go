package main

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/atotto/clipboard"
	"github.com/gen2brain/beeep"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	// 無限ループしながらclipboardが更新されるのを待つ
	for {
		now, err := clipboard.ReadAll()
		if err != nil {
			panic(err)
		}

		if text == now {
			continue
		} else {
			len := utf8.RuneCountInString(now)
			msg := fmt.Sprintf("count: %d", len)
			err := beeep.Notify("c4", msg, "")
			if err != nil {
				panic(err)
			}
			text = now
		}
		time.Sleep(1 * time.Second)
	}
}
