// websockets.go
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
	"github.com/pborman/getopt/v2"

	"wedpad/internal/journal"
	"wedpad/internal/logreader"
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

	// set proggie termination signal handler(s)
	done := make(chan bool)
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		for sig := range sigChan {
			slog.Info("Got signal '%s'", sig)
			done <- true
		}
	}()

	// setup logger
	//slog.Init("wedpad", debug, "2006-01-02 15:04:05")  // no need if using systemd
	slog.Init("", debug, "")

	slog.Info("Started")
	defer slog.Info("Exited")

	// read config file
	if err := sconf.Read(configFile); err != nil {
		slog.Fatal("Config failed: %s", err)
	}

	// init messaging subsystem
	if err := msg.Init(); err != nil {
		slog.Fatal("Msg subsystem init failed: %s", err)
	}

	// init ED log file reader
	if err := logreader.Init(); err != nil {
		slog.Fatal("Logreader subsystem init failed: %s", err)
	}

	// init journal processing
	if err := journal.Init(); err != nil {
		slog.Fatal("Journal subsystem init failed: %s", err)
	}

	// run http server
	go server.Run()

	// now wait
	<-done

}
