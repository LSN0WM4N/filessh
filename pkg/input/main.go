package input

import (
	"context"
	"os"

	"github.com/LSN0WM4N/filessh/pkg/bus"
	"github.com/LSN0WM4N/filessh/pkg/plugins"
	"golang.org/x/term"
)

func ReadInput(ctx context.Context, b *bus.EventBus, registry *plugins.Registry) {
	fd := int(os.Stdin.Fd())

	buf := make([]byte, 16)
	read := make(chan []byte)

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState)

	go func() {
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				close(read)
				return
			}
			chunk := make([]byte, n)
			copy(chunk, buf[:n])
			read <- chunk
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return

		case chunk, ok := <-read:
			if !ok {
				return
			}

			event := parseInput(chunk)

			if event.Payload.(bus.KeyInfo).Seq == "alt+q" {
				b.Publish(event)
				continue
			}

			if event.Payload.(bus.KeyInfo).Seq == "ctrl+c" {
				b.Publish(bus.Event{Type: bus.EventQuit})
				continue
			}

			if event.Payload.(bus.KeyInfo).Seq == "enter" {
				if focused := registry.Focused(); focused != nil && focused.ID() == "terminal" {
					focused.OnKey(event)
				} else {
					b.Publish(event)
				}
				continue
			}

			if focused := registry.Focused(); focused != nil && focused.ID() == "terminal" {
				focused.OnKey(event)
				continue
			}

			b.Publish(event)
		}
	}
}

func parseInput(buf []byte) bus.Event {
	key := bus.KeyInfo{}

	switch {
	case len(buf) == 2 && buf[0] == 0x1b && buf[1] == 'q':
		key.Seq = "alt+q"

	case len(buf) >= 3 && buf[0] == 0x1b && buf[1] == '[':
		switch buf[2] {
		case 'A':
			key.Seq = "up"
		case 'B':
			key.Seq = "down"
		case 'C':
			key.Seq = "right"
		case 'D':
			key.Seq = "left"
		}

	case len(buf) == 1 && buf[0] == 0x1b:
		key.Seq = "escape"

	case len(buf) == 1 && buf[0] == '\r':
		key.Seq = "enter"

	case len(buf) == 1 && buf[0] == 127:
		key.Seq = "backspace"

	case len(buf) == 1 && buf[0] == 0x03:
		key.Seq = "ctrl+c"

	case len(buf) == 1 && buf[0] == '\t':
		key.Seq = "tab"

	default:
		if len(buf) == 1 && buf[0] >= 32 && buf[0] < 127 {
			key.Key = rune(buf[0])
		}
	}

	return bus.Event{
		Type:    bus.EventKey,
		Payload: key,
	}
}
