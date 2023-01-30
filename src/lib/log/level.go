package log

// Level ...
type Level int

const (
	// DebugLevel debug
	DebugLevel Level = iota
	// InfoLevel info
	InfoLevel
	// WarningLevel warning
	WarningLevel
	// ErrorLevel error
	ErrorLevel
	// FatalLevel fatal
	FatalLevel
)
