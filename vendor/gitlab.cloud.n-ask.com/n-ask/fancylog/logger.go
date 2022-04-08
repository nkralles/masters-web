package fancylog

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

// FdWriter interface extends existing io.Writer with file descriptor function
// support
type FdWriter interface {
	io.Writer
	Fd() uintptr
}

// Logger struct define the underlying storage for single logger
type Logger struct {
	name      string
	color     bool
	out       FdWriter
	err       FdWriter
	debug     bool
	timestamp bool
	quiet     bool
	mu        sync.Mutex
	//buf       ColorLogger
}

// Prefix struct define plain and color byte
type Prefix struct {
	Plain []byte
	Color []byte
	File  bool
}

var (
	// Plain prefix template
	plainFatal = []byte("[FATAL] ")
	plainError = []byte("[ERROR] ")
	plainWarn  = []byte("[WARN]  ")
	plainInfo  = []byte("[INFO]  ")
	plainDebug = []byte("[DEBUG] ")
	plainTrace = []byte("[TRACE] ")

	// FatalPrefix show fatal prefix
	FatalPrefix = Prefix{
		Plain: plainFatal,
		Color: Red(plainFatal),
		File:  true,
	}

	// ErrorPrefix show error prefix
	ErrorPrefix = Prefix{
		Plain: plainError,
		Color: Red(plainError),
		File:  true,
	}

	// WarnPrefix show warn prefix
	WarnPrefix = Prefix{
		Plain: plainWarn,
		Color: Orange(plainWarn),
	}

	// InfoPrefix show info prefix
	InfoPrefix = Prefix{
		Plain: plainInfo,
		Color: Green(plainInfo),
	}

	// DebugPrefix show info prefix
	DebugPrefix = Prefix{
		Plain: plainDebug,
		Color: Purple(plainDebug),
		File:  true,
	}

	// TracePrefix show info prefix
	TracePrefix = Prefix{
		Plain: plainTrace,
		Color: Cyan(plainTrace),
	}
)

// New returns new Logger instance with predefined writer output and
// automatically detect terminal coloring support
func New(out FdWriter) *Logger {
	return &Logger{
		color:     terminal.IsTerminal(int(out.Fd())),
		out:       out,
		err:       out,
		timestamp: true,
	}
}

// NewWithError returns new Logger instance with predefined writer output and
// automatically detect terminal coloring support
func NewWithError(out FdWriter, err FdWriter) *Logger {
	return &Logger{
		color:     terminal.IsTerminal(int(out.Fd())),
		out:       out,
		err:       err,
		timestamp: true,
	}
}

// NewWithName {(name string out FdWriter) *Logger { returns new Logger instance with predefined writer output and
// automatically detect terminal coloring support
func NewWithName(name string, out FdWriter) *Logger {
	return &Logger{
		name:      name,
		color:     terminal.IsTerminal(int(out.Fd())),
		out:       out,
		err:       out,
		timestamp: true,
	}
}

// NewWithNameAndError {(name string out FdWriter) *Logger { returns new Logger instance with predefined writer output and
// automatically detect terminal coloring support
func NewWithNameAndError(name string, out FdWriter, err FdWriter) *Logger {
	return &Logger{
		name:      name,
		color:     terminal.IsTerminal(int(out.Fd())),
		out:       out,
		err:       err,
		timestamp: true,
	}
}

// WithColor explicitly turn on colorful features on the log
func (l *Logger) WithColor() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.color = true
	return l
}

// WithoutColor explicitly turn off colorful features on the log
func (l *Logger) WithoutColor() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.color = false
	return l
}

// WithDebug turn on debugging output on the log to reveal debug and trace level
func (l *Logger) WithDebug() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.debug = true
	return l
}

// WithoutDebug turn off debugging output on the log
func (l *Logger) WithoutDebug() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.debug = false
	return l
}

// IsDebug check the state of debugging output
func (l *Logger) IsDebug() bool {
	return l.debug
}

// WithTimestamp turn on timestamp output on the log
func (l *Logger) WithTimestamp() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timestamp = true
	return l
}

// WithoutTimestamp turn off timestamp output on the log
func (l *Logger) WithoutTimestamp() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timestamp = false
	return l
}

// Quiet turn off all log output
func (l *Logger) Quiet() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.quiet = true
	return l
}

// NoQuiet turn on all log output
func (l *Logger) NoQuiet() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.quiet = false
	return l
}

// IsQuiet check for quiet state
func (l *Logger) IsQuiet() bool {
	return l.quiet
}

func (l *Logger) OutputMap(prefix Prefix, data map[string]interface{}, isErr bool) error {
	// Check if quiet is requested, and try to return no error and be quiet
	if l.IsQuiet() {
		return nil
	}
	b := NewColorLogger()

	// Reset buffer so it start from the begining
	b.Reset()
	if len(l.name) > 0 {
		if l.color {
			b.NicePurple()
		}
		b.Append([]byte("<"))
		b.Append([]byte(l.name))
		b.Append([]byte("> "))
		if l.color {
			b.Off()
		}
	}

	// Write prefix to the buffer
	if l.color {
		b.Append(prefix.Color)
	} else {
		b.Append(prefix.Plain)
	}
	// Check if the log require timestamping
	if l.timestamp {
		// Print timestamp color if color enabled
		if l.color {
			b.Blue()
		}
		b.AppendTime(time.Now().UTC(), time.RFC3339)
		b.AppendSpace()
		// Print reset color if color enabled
		if l.color {
			b.Off()
		}
	}

	for key, val := range data {
		if l.color {
			b.Purple()
		}
		b.Append([]byte(key))
		if l.color {
			b.Orange()
		}
		b.Append([]byte("="))
		if l.color {
			b.Cyan()
		}
		b.Append([]byte(fmt.Sprintf("%+v", val)))
		b.AppendSpace()
		if l.color {
			b.Off()
		}
	}
	b.AppendByte('\n')
	// Flush buffer to output
	var err error
	if isErr {
		_, err = l.err.Write(b.Bytes())
	} else {
		_, err = l.out.Write(b.Bytes())
	}
	b.Free()
	return err
}

// Output print the actual value
func (l *Logger) Output(depth int, prefix Prefix, data string, isErr bool) error {
	// Check if quiet is requested, and try to return no error and be quiet
	if l.IsQuiet() {
		return nil
	}

	// Temporary storage for file and line tracing
	var file string
	var line int
	var fn string
	// Check if the specified prefix needs to be included with file logging
	if prefix.File {
		var ok bool
		var pc uintptr
		// Get the caller filename and line
		if pc, file, line, ok = runtime.Caller(depth + 1); !ok {
			file = "<unknown file>"
			fn = "<unknown function>"
			line = 0
		} else {
			file = filepath.Base(file)
			fn = runtime.FuncForPC(pc).Name()
		}
	}
	b := NewColorLogger()
	// Reset buffer so it start from the begining
	b.Reset()
	// Write prefix to the buffer
	if len(l.name) > 0 {
		if l.color {
			b.NicePurple()
		}
		b.Append([]byte("<"))
		b.Append([]byte(l.name))
		b.Append([]byte("> "))
		if l.color {
			b.Off()
		}
	}

	if l.color {
		b.Append(prefix.Color)
	} else {
		b.Append(prefix.Plain)
	}
	// Check if the log require timestamping
	if l.timestamp {
		// Print timestamp color if color enabled
		if l.color {
			b.Blue()
		}
		b.AppendTime(time.Now().UTC(), time.RFC3339)
		b.AppendSpace()
		// Print reset color if color enabled
		if l.color {
			b.Off()
		}
	}
	// Add caller filename and line if enabled
	if prefix.File {
		// Print color start if enabled
		if l.color {
			b.Orange()
		}
		// Print filename and line
		b.Append([]byte(fn))
		b.AppendByte(':')
		b.Append([]byte(file))
		b.AppendByte(':')
		b.AppendInt(int64(line))
		b.AppendByte(' ')
		// Print color stop
		if l.color {
			b.Off()
		}
	}

	// Print the actual string data from caller
	b.Append([]byte(data))
	if len(data) == 0 || data[len(data)-1] != '\n' {
		b.AppendByte('\n')
	}
	// Flush buffer to output
	var err error
	if isErr {
		_, err = l.err.Write(b.Bytes())
	} else {
		_, err = l.out.Write(b.Bytes())
	}

	b.Free()
	return err
}

// Fatal print fatal message to output and quit the application with status 1
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(1, FatalPrefix, fmt.Sprintln(v...), true)
	os.Exit(1)
}

// Fatalf print formatted fatal message to output and quit the application
// with status 1
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(1, FatalPrefix, fmt.Sprintf(format, v...), true)
	os.Exit(1)
}
func (l *Logger) FatalMap(v map[string]interface{}) {
	l.OutputMap(FatalPrefix, v, true)
}

// Error print error message to output
func (l *Logger) Error(v ...interface{}) {
	l.Output(1, ErrorPrefix, fmt.Sprintln(v...), true)
}

// Errorf print formatted error message to output
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(1, ErrorPrefix, fmt.Sprintf(format, v...), true)
}
func (l *Logger) ErrorMap(v map[string]interface{}) {
	l.OutputMap(ErrorPrefix, v, true)
}

// Warn print warning message to output
func (l *Logger) Warn(v ...interface{}) {
	l.Output(1, WarnPrefix, fmt.Sprintln(v...), false)
}

// Warnf print formatted warning message to output
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(1, WarnPrefix, fmt.Sprintf(format, v...), false)
}
func (l *Logger) WarnMap(v map[string]interface{}) {
	l.OutputMap(WarnPrefix, v, false)
}

// Info print informational message to output
func (l *Logger) Info(v ...interface{}) {
	l.Output(1, InfoPrefix, fmt.Sprintln(v...), false)
}

// Infof print formatted informational message to output
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(1, InfoPrefix, fmt.Sprintf(format, v...), false)
}

func (l *Logger) InfoMap(v map[string]interface{}) {
	l.OutputMap(InfoPrefix, v, false)
}

// Debug print debug message to output if debug output enabled
func (l *Logger) Debug(v ...interface{}) {
	if l.IsDebug() {
		l.Output(1, DebugPrefix, fmt.Sprintln(v...), false)
	}
}

// Debugf print formatted debug message to output if debug output enabled
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.IsDebug() {
		l.Output(1, DebugPrefix, fmt.Sprintf(format, v...), false)
	}
}

func (l *Logger) DebugMap(v map[string]interface{}) {
	l.OutputMap(DebugPrefix, v, false)
}

// Trace print trace message to output if debug output enabled
func (l *Logger) Trace(v ...interface{}) {
	if l.IsDebug() {
		l.Output(1, TracePrefix, fmt.Sprintln(v...), false)
	}
}

// Tracef print formatted trace message to output if debug output enabled
func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.IsDebug() {
		l.Output(1, TracePrefix, fmt.Sprintf(format, v...), false)
	}
}
