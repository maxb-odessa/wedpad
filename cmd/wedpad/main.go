// websockets.go
package main

import (
	"os"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
	"github.com/pborman/getopt/v2"

	"wedpad/internal/msg"
	"wedpad/internal/server"
)

func main() {

	// get cmdline args and parse them
	help := false
	debug := 0
	configFile := os.ExpandEnv("$HOME/.local/etc/wedpad.conf")
	getopt.HelpColumn = 0
	getopt.FlagLong(&help, "help", 'h', "Show this help")
	getopt.FlagLong(&debug, "debug", 'd', "Set debug log level")
	getopt.FlagLong(&configFile, "config", 'c', "Path to config file")
	getopt.Parse()

	// help-only requested
	if help {
		getopt.Usage()
		return
	}

	// setup logger
	slog.Init("wedpad", debug, "2006-01-02 15:04:05")

	slog.Info("Started")

	// read config file
	if err := sconf.Read(configFile); err != nil {
		slog.Fatal("Config failed: %s", err)
	}

	// init journal subsystem

	// init messaging subsystem
	if err := msg.Init(); err != nil {
		slog.Warn("Messages system init is incomplete: %s", err)
	}

	// init audio sybsytem

	// run http subsystem
	server.Run()
}
