package logger

import (
	"github.com/nkralles/masters-web/internal/cfg"
	"gitlab.cloud.n-ask.com/n-ask/fancylog"
	"os"
	"syscall"
)

var Internal *fancylog.Logger
var REST *fancylog.Logger
var PG *fancylog.Logger

func init() {
	devnull := os.NewFile(uintptr(syscall.Stdin), os.DevNull)
	if cfg.Config.Verbose {
		Internal = fancylog.NewWithName("internal", os.Stdout).WithColor().WithTimestamp().WithDebug()
		Internal.Debug("Running in Verbose Debug level")
		REST = fancylog.NewWithName("api", os.Stdout).WithColor().WithTimestamp().WithDebug()
		PG = fancylog.NewWithNameAndError("postgres", os.Stdout, os.Stdout).WithColor().WithTimestamp().WithDebug()
	} else {
		Internal = fancylog.NewWithName("internal", os.Stdout).WithColor().WithTimestamp()
		REST = fancylog.NewWithName("api", os.Stdout).WithColor().WithTimestamp()
		PG = fancylog.NewWithNameAndError("postgres", devnull, os.Stdout).WithColor().WithTimestamp().WithDebug()
	}
}
