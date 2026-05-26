package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/LSN0WM4N/filessh/pkg/bus"
	"github.com/LSN0WM4N/filessh/pkg/explorer"
	"github.com/LSN0WM4N/filessh/pkg/input"
	"github.com/LSN0WM4N/filessh/pkg/plugins"
	"github.com/LSN0WM4N/filessh/pkg/pty"
	"github.com/LSN0WM4N/filessh/pkg/sshclient"
)

func main() {
	godotenv.Load()

	host := os.Getenv("SSH_HOST")
	port := os.Getenv("SSH_PORT")
	user := os.Getenv("SSH_USER")
	pass := os.Getenv("SSH_PASS")

	// Stablish a connection
	client, err := sshclient.SetupConnection(sshclient.UserConfig{
		Host:     host,
		Port:     port,
		Username: &user,
		Password: &pass,
	})

	if err != nil {
		panic(err)
	}

	defer client.Close()

	// Setup main bus and plugins
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventBus := bus.NewEventBus()
	registry := plugins.NewRegistry()

	pluginCtx := plugins.PluginContext{
		Bus:     eventBus,
		Session: client,
		Fd:      int(os.Stdin.Fd()),
	}

	loadedPlugins := []plugins.Plugin{
		// &DirectoryTreePlugin{},
		&explorer.ExplorerPlugin{},
		&pty.TerminalPlugin{},
	}
	for _, p := range loadedPlugins {
		p.Init(pluginCtx)
		registry.Register(p)
	}

	eventBus.Subscribe(bus.EventKey, func(e bus.Event) {
		fmt.Println(e)

		// focused := registry.Focused().ID()
		if e.Payload.(bus.KeyInfo).Seq == "alt+q" {
			// if focused == "terminal" {
			// 	registry.SetFocus("explorer")
			// } else {
			// 	registry.SetFocus("terminal")
			// }
			fmt.Printf("[Changed focus]\n")
		}
	})

	// eventBus.Subscribe(bus.EventFocus, func(e bus.Event) {
	// 	registry.SetFocus(e.Payload.(string))
	// })

	// eventBus.Subscribe(bus.EventKey, func(e bus.Event) {
	// 	if p := registry.Focused(); p != nil {
	// 		p.OnKey(e)
	// 	}
	// })

	// eventBus.Subscribe(bus.EventResize, func(e bus.Event) {
	// 	for _, p := range registry.All() {
	// 		p.OnEvent(e)
	// 	}
	// })

	// registry.SetFocus("explorer")

	eventBus.Subscribe(bus.EventQuit, func(e bus.Event) {
		cancel()
	})

	go eventBus.Run(ctx)
	go input.ReadInput(ctx, eventBus, registry)

	// PTYSession, _ := sshclient.PTYMode(session, ctx, eventBus)
	// TUISession, _ := sshclient.PipeMode(session, ctx, eventBus)

	<-ctx.Done()
	os.Exit(0)
}
