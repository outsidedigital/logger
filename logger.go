package logger

import "os"

// Logger implements a structured, leveled logger.
type Logger struct {
	w   Writer
	lvl Level
	hh  []Hook
	ff  []Field
}

// NewLogger creates a new logger that outputs to the stderr.
func NewLogger() Logger {
	return Logger{
		w:   TextWriter(os.Stderr),
		lvl: LevelInfo,
	}
}

// Entry creates a new logging entry at the given level.
func (log Logger) Entry(lvl Level) Entry {
	w := HookWriter(log.w, log.lvl)
	for _, h := range log.hh {
		w = HookWriter(w, h)
	}
	return NewEntry(w).With(lvl).With(log.ff...)
}

// Error creates a new logging entry at the error level and appends the given
// errors to it.
func (log Logger) Error(errs ...error) Entry {
	entry := log.Entry(LevelError)
	for _, err := range errs {
		entry.ff = append(entry.ff, Error{err})
	}
	return entry
}

// Warn creates a new logging entry at the warn level and appends the given
// errors to it.
func (log Logger) Warn(errs ...error) Entry {
	entry := log.Entry(LevelWarn)
	for _, err := range errs {
		entry.ff = append(entry.ff, Error{err})
	}
	return entry
}

// Info creates a new logging entry at the info level.
func (log Logger) Info() Entry {
	return log.Entry(LevelInfo)
}

// Debug creates a new logging entry at the debug level.
func (log Logger) Debug() Entry {
	return log.Entry(LevelDebug)
}

// With returns the logger configuration context.
func (log Logger) With() Options {
	return Options{log: &log}
}
