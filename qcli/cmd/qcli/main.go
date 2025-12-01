package main

import (
	"os"

	"github.com/qtraffics/qtfra/log"
	"github.com/qtraffics/qtfra/log/loghandler"
	"github.com/qtraffics/qtfra/sys/sysvars"
	"github.com/qtraffics/repo/qcli"

	"github.com/spf13/cobra"
)

var logNoColor bool

var mainCommand = &cobra.Command{
	Use: qcli.Name,
}

func init() {
	mainCommand.PersistentFlags().BoolVar(
		&logNoColor, "no-color", false, "Disable log color.")
}

func main() {
	setupDefaultLogger()
	if err := mainCommand.Execute(); err != nil {
		panic(err)
	}
}

func setupDefaultLogger() {
	var options loghandler.ConsoleHandlerOption

	options.EnableTime = true
	options.SourceLevel = log.LevelDisable
	options.Level = log.LevelError

	if sysvars.DebugEnabled {
		options.SourceLevel = log.LevelWarn
		options.Level = log.LevelDebug
	}
	if logNoColor {
		options.LevelFormatter = log.EqualLengthLevelFormatter
	}

	handler := loghandler.NewConsoleHandler(os.Stdout, options)
	log.SetDefaultLogger(log.New(handler))
}
