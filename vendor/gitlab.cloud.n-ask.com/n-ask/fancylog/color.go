package fancylog

import "go.uber.org/zap/buffer"

type ColorLogger struct {
	*buffer.Buffer
}

var (
	_pool = buffer.NewPool()
	// Get retrieves a buffer from the pool, creating one if necessary.
	Get = _pool.Get
)
// color pallete map
var (
	colorOff    = []byte("\u001B[0m")
	colorRed    = []byte("\u001B[0;31m")
	colorGreen  = []byte("\u001B[0;32m")
	colorOrange = []byte("\u001B[0;33m")
	colorBlue   = []byte("\u001B[0;34m")
	colorPurple = []byte("\u001B[0;35m")
	colorCyan   = []byte("\u001B[0;36m")
	colorGray   = []byte("\u001B[0;37m")

	colorDarkOrange  = []byte("\u001b[1m\u001b[38;5;202m")
	colorBrightWhite = []byte("\u001b[1m\u001b[38;5;255m")
	colorNicePurple  = []byte("\u001b[1m\u001b[38;5;99m")
)

func NewColorLogger() ColorLogger {
	return ColorLogger{Get()}
}

// Off apply no color to the data
func (cb ColorLogger) Off() {
	_, _ = cb.Write(colorOff)
}

// Red apply red color to the data
func (cb *ColorLogger) Red() {
	_, _ = cb.Write(colorRed)
}

// Green apply green color to the data
func (cb *ColorLogger) Green() {
	_, _ = cb.Write(colorGreen)
}

// Orange apply orange color to the data
func (cb *ColorLogger) Orange() {
	_, _ = cb.Write(colorOrange)
}

// Blue apply blue color to the data
func (cb *ColorLogger) Blue() {
	_, _ = cb.Write(colorBlue)
}

// Purple apply purple color to the data
func (cb *ColorLogger) Purple() {
	_, _ = cb.Write(colorPurple)
}

// Cyan apply cyan color to the data
func (cb *ColorLogger) Cyan() {
	_, _ = cb.Write(colorCyan)
}

// Gray apply gray color to the data
func (cb *ColorLogger) Gray() {
	_, _ = cb.Write(colorGray)
}

// White apply gray color to the data
func (cb *ColorLogger) White() {
	_, _ = cb.Write(colorBrightWhite)
}

// BrightOrange apply gray color to the data
func (cb *ColorLogger) BrightOrange() {
	_, _ = cb.Write(colorDarkOrange)
}

// NicePurple apply gray color to the data
func (cb *ColorLogger) NicePurple() {
	_, _ = cb.Write(colorNicePurple)
}

func (cb *ColorLogger) AppendSpace() {
	_, _ = cb.Write([]byte(" "))
}

// Append byte slice to buffer
func (cb *ColorLogger) Append(data []byte) {
	_, _ = cb.Write(data)
}

// mixer mix the color on and off byte with the actual data
func mixer(data []byte, color []byte) []byte {
	var result []byte
	return append(append(append(result, color...), data...), colorOff...)
}

// Red apply red color to the data
func Red(data []byte) []byte {
	return mixer(data, colorRed)
}

// Green apply green color to the data
func Green(data []byte) []byte {
	return mixer(data, colorGreen)
}

// Orange apply orange color to the data
func Orange(data []byte) []byte {
	return mixer(data, colorOrange)
}

// Blue apply blue color to the data
func Blue(data []byte) []byte {
	return mixer(data, colorBlue)
}

// Purple apply purple color to the data
func Purple(data []byte) []byte {
	return mixer(data, colorPurple)
}

// Cyan apply cyan color to the data
func Cyan(data []byte) []byte {
	return mixer(data, colorCyan)
}

// Gray apply gray color to the data
func Gray(data []byte) []byte {
	return mixer(data, colorGray)
}
