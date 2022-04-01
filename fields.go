package logger

import (
	"fmt"
	"runtime"
	"time"
)

// Field is a generic interface of the logging entry field.
type Field interface {
	// Encode encodes the field with the given encoder.
	Encode(Encoder)
}

// FieldFunc is an adapter to allow the use of ordinary functions as fields.
type FieldFunc func(Encoder)

// Encode encodes the field with the given encoder.
func (f FieldFunc) Encode(enc Encoder) {
	f(enc)
}

// Well-known field names.
const (
	FieldCaller  = "caller"
	FieldError   = "error"
	FieldLevel   = "level"
	FieldMessage = "message"
	FieldName    = "log"
	FieldSpan    = "span"
	FieldTime    = "time"
)

// Bool creates a new field with the given key and boolean value.
func Bool(key string, b bool) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeBool(key, b)
	}
}

// Bytes creates a new field with the given key and bytes value.
func Bytes(key string, p []byte) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeBytes(key, p)
	}
}

// Caller creates a new field with current file and line number.
func Caller(skip int) FieldFunc {
	_, file, line, ok := runtime.Caller(skip + 1)
	caller := fmt.Sprintf("%s:%d", file, line)
	return func(enc Encoder) {
		if ok {
			enc.EncodeString(FieldCaller, caller)
		}
	}
}

// Duration creates a new field with the given key and duration value.
func Duration(key string, d time.Duration) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeDuration(key, d)
	}
}

// Error represents an error field.
type Error struct {
	error
}

// Encode encodes the error with the given encoder.
func (err Error) Encode(enc Encoder) {
	enc.EncodeError(FieldError, err.error)
}

// Unwrap returns an underlying error.
func (err Error) Unwrap() error {
	return err.error
}

// Float32 creates a new field with the given key and float32 value.
func Float32(key string, f float32) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeFloat32(key, f)
	}
}

// Float64 creates a new field with the given key and float64 value.
func Float64(key string, f float64) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeFloat64(key, f)
	}
}

// Int creates a new field with the given key and integer value.
func Int(key string, i int) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeInt(key, i)
	}
}

// Int32 creates a new field with the given key and int32 value.
func Int32(key string, i int32) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeInt32(key, i)
	}
}

// Int64 creates a new field with the given key and int64 value.
func Int64(key string, i int64) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeInt64(key, i)
	}
}

// Message represents a message field.
type Message string

// Encode encodes the message with the given encoder.
func (msg Message) Encode(enc Encoder) {
	enc.EncodeString(FieldMessage, string(msg))
}

// Name represents a logger name field.
type Name string

// Encode encodes the logger name with the given encoder.
func (name Name) Encode(enc Encoder) {
	enc.EncodeString(FieldName, string(name))
}

// Span represents a time span field and contains a start time of the span.
type Span time.Time

// Encode encodes the time span with the given encoder.
func (span Span) Encode(enc Encoder) {
	enc.EncodeDuration(FieldSpan, time.Since(time.Time(span)))
}

// String creates a new field with the given key and string value.
func String(key, s string) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeString(key, s)
	}
}

// Time creates a new field with the given key and time value.
func Time(key string, t time.Time) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeTime(key, t)
	}
}

// Timestamp represents the current time field.
type Timestamp struct{}

// Encode encodes the timestamp with the given encoder.
func (_ Timestamp) Encode(enc Encoder) {
	enc.EncodeTime(FieldTime, time.Now())
}

// Uint creates a new field with the given key and unsigned integer value.
func Uint(key string, i uint) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeUint(key, i)
	}
}

// Uint32 creates a new field with the given key and uint32 value.
func Uint32(key string, i uint32) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeUint32(key, i)
	}
}

// Uint64 creates a new field with the given key and uint64 value.
func Uint64(key string, i uint64) FieldFunc {
	return func(enc Encoder) {
		enc.EncodeUint64(key, i)
	}
}
