package fancylog

import (
	"context"
	"github.com/jackc/pgx/v4"
	"regexp"
	"strings"
)

type PgLogger interface {
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
	Crit(msg string, ctx ...interface{})
}

type FancyPGLogger struct {
	l *Logger
}

func NewLogger(l *Logger) *FancyPGLogger {
	return &FancyPGLogger{l: l}
}

func (l *FancyPGLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	if val, ok := data["sql"]; ok {
		sql := val.(string)
		re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
		re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
		final := re_leadclose_whtsp.ReplaceAllString(sql, "")
		final = re_inside_whtsp.ReplaceAllString(final, " ")
		if strings.HasPrefix(strings.ToUpper(sql), "REFRESH MATERIALIZED VIEW") || strings.HasPrefix(strings.ToUpper(sql), "NOTIFY") {
			//EAT
			return
		}
		data["sql"] = final
	}
	logArgs := make([]interface{}, 0, len(data))
	for k, v := range data {
		logArgs = append(logArgs, k, v)
	}
	switch level {
	case pgx.LogLevelTrace:
		l.Debug(msg, append(logArgs, "PGX_LOG_LEVEL", level)...)
	case pgx.LogLevelDebug:
		l.Debug(msg, logArgs...)
	case pgx.LogLevelInfo:
		l.Info(msg, logArgs...)
	case pgx.LogLevelWarn:
		l.Warn(msg, logArgs...)
	case pgx.LogLevelError:
		l.Error(msg, logArgs...)
	default:
		l.Error(msg, append(logArgs, "INVALID_PGX_LOG_LEVEL", level)...)
	}
}

func (l *FancyPGLogger) Debug(msg string, ctx ...interface{}) {
	l.l.DebugMap(toMap(msg, ctx))
}

func (l *FancyPGLogger) Info(msg string, ctx ...interface{}) {
	l.l.InfoMap(toMap(msg, ctx))
}

func (l *FancyPGLogger) Warn(msg string, ctx ...interface{}) {
	l.l.WarnMap(toMap(msg, ctx))
}

func (l *FancyPGLogger) Error(msg string, ctx ...interface{}) {
	l.l.ErrorMap(toMap(msg, ctx))
}

func (l *FancyPGLogger) Crit(msg string, ctx ...interface{}) {
	l.l.FatalMap(toMap(msg, ctx))
}

func toMap(msg string, ctx []interface{}) map[string]interface{} {
	v := make(map[string]interface{})
	v["msg"] = msg
	if len(ctx)%2 == 0 {
		for i := 0; i < len(ctx); i = i + 2 {
			v[ctx[i].(string)] = ctx[i+1]
		}
	} else {
		v["data"] = ctx
	}
	return v
}
