package logger

import (
	"github.com/nkralles/masters-web/internal/cfg"
	"gitlab.cloud.n-ask.com/n-ask/fancylog"
	"os"
)

var Internal *fancylog.Logger
var REST *fancylog.Logger
var PG *fancylog.Logger

func init() {
	if cfg.Config.Verbose {
		Internal = fancylog.NewWithName("internal", os.Stdout).WithColor().WithTimestamp().WithDebug()
		Internal.Debug("Running in Verbose Debug level")
		REST = fancylog.NewWithName("api", os.Stdout).WithColor().WithTimestamp().WithDebug()
		PG = fancylog.NewWithName("postgres", os.Stdout).WithColor().WithTimestamp().WithDebug()
	} else {
		Internal = fancylog.NewWithName("internal", os.Stdout).WithColor().WithTimestamp()
		REST = fancylog.NewWithName("api", os.Stdout).WithColor().WithTimestamp()
		PG = fancylog.NewWithName("postgres", os.Stdout).WithColor().WithTimestamp()
	}
}
