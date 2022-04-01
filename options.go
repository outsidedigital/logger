package logger

import (
	"fmt"
	"time"
)

// Options represents a logger configuration context.
type Options struct {
	log *Logger
}

// Logger returns a previously configured logger.
func (o Options) Logger() Logger {
	return *o.log
}

// Fields appends given fields to the logger.
func (o Options) Fields(ff ...Field) Options {
	o.log.ff = append(o.log.ff, ff...)
	return o
}

// Hooks appends given hooks to the logger.
func (o Options) Hooks(hh ...Hook) Options {
	o.log.hh = append(o.log.hh, hh...)
	return o
}

// Level changes the logging priority level.
func (o Options) Level(lvl Level) Options {
	o.log.lvl = lvl
	return o
}

// Writer changes the logger output writer.
func (o Options) Writer(w Writer) Options {
	o.log.w = w
	return o
}

// Bool appends a new field with the given key and boolean value.
func (o Options) Bool(key string, b bool) Options {
	o.log.ff = append(o.log.ff, Bool(key, b))
	return o
}

// Bytes appends a new field with the given key and bytes value.
func (o Options) Bytes(key string, p []byte) Options {
	o.log.ff = append(o.log.ff, Bytes(key, p))
	return o
}

// Caller appends a new field with current file and line number.
func (o Options) Caller(skip int) Options {
	o.log.ff = append(o.log.ff, Caller(skip+1))
	return o
}

// Duration appends a new field with the given key and duration value.
func (o Options) Duration(key string, d time.Duration) Options {
	o.log.ff = append(o.log.ff, Duration(key, d))
	return o
}

// Error appends given error to the entry.
func (o Options) Error(err error) Options {
	o.log.ff = append(o.log.ff, Error{err})
	return o
}

// Errorf appends a new formatted error to the entry.
func (o Options) Errorf(format string, a ...any) Options {
	//nolint:goerr113 // Errorf is a wrapper for errorf.
	o.log.ff = append(o.log.ff, Error{fmt.Errorf(format, a...)})
	return o
}

// Float32 appends a new field with the given key and float32 value.
func (o Options) Float32(key string, f float32) Options {
	o.log.ff = append(o.log.ff, Float32(key, f))
	return o
}

// Float64 appends a new field with the given key and float64 value.
func (o Options) Float64(key string, f float64) Options {
	o.log.ff = append(o.log.ff, Float64(key, f))
	return o
}

// Int appends a new field with the given key and integer value.
func (o Options) Int(key string, i int) Options {
	o.log.ff = append(o.log.ff, Int(key, i))
	return o
}

// Int32 appends a new field with the given key and int32 value.
func (o Options) Int32(key string, i int32) Options {
	o.log.ff = append(o.log.ff, Int32(key, i))
	return o
}

// Int64 appends a new field with the given key and int64 value.
func (o Options) Int64(key string, i int64) Options {
	o.log.ff = append(o.log.ff, Int64(key, i))
	return o
}

// Name appends a new field with the given logger name.
func (o Options) Name(name string) Options {
	o.log.ff = append(o.log.ff, Name(name))
	return o
}

// Span appends a new time span field that begins at the current time.
func (o Options) Span() Options {
	o.log.ff = append(o.log.ff, Span(time.Now()))
	return o
}

// String appends a new field with the given key and string value.
func (o Options) String(key, s string) Options {
	o.log.ff = append(o.log.ff, String(key, s))
	return o
}

// Stringer appends a new field with the given key and value that implements
// stringer interface.
func (o Options) Stringer(key string, v fmt.Stringer) Options {
	o.log.ff = append(o.log.ff, String(key, v.String()))
	return o
}

// Stringf appends a new field with the given key and formatted string value.
func (o Options) Stringf(key, format string, a ...any) Options {
	o.log.ff = append(o.log.ff, String(key, fmt.Sprintf(format, a...)))
	return o
}

// Time appends a new field with the given key and time value.
func (o Options) Time(key string, t time.Time) Options {
	o.log.ff = append(o.log.ff, Time(key, t))
	return o
}

// Timestamp appends a new field with the current time.
func (o Options) Timestamp() Options {
	o.log.ff = append(o.log.ff, Timestamp{})
	return o
}

// Uint appends a new field with the given key and unsigned integer value.
func (o Options) Uint(key string, i uint) Options {
	o.log.ff = append(o.log.ff, Uint(key, i))
	return o
}

// Uint32 appends a new field with the given key and uint32 value.
func (o Options) Uint32(key string, i uint32) Options {
	o.log.ff = append(o.log.ff, Uint32(key, i))
	return o
}

// Uint64 appends a new field with the given key and uint64 value.
func (o Options) Uint64(key string, i uint64) Options {
	o.log.ff = append(o.log.ff, Uint64(key, i))
	return o
}
