package logger

import "time"

// Encoder is a generic interface of the logging encoder.
type Encoder interface {
	// EncodeBool encodes a field with the given key and boolean value.
	EncodeBool(key string, b bool)
	// EncodeBytes encodes a field with the given key and bytes value.
	EncodeBytes(key string, p []byte)
	// EncodeDuration encodes a field with the given key and duration value.
	EncodeDuration(key string, d time.Duration)
	// EncodeError encodes a field with the given key and error value.
	EncodeError(key string, err error)
	// EncodeFloat32 encodes a field with the given key and float32 value.
	EncodeFloat32(key string, f float32)
	// EncodeFloat64 encodes a field with the given key and float64 value.
	EncodeFloat64(key string, f float64)
	// EncodeInt encodes a field with the given key and integer value.
	EncodeInt(key string, i int)
	// EncodeInt32 encodes a field with the given key and int32 value.
	EncodeInt32(key string, i int32)
	// EncodeInt64 encodes a field with the given key and int64 value.
	EncodeInt64(key string, i int64)
	// EncodeString encodes a field with the given key and string value.
	EncodeString(key, s string)
	// EncodeTime encodes a field with the given key and time value.
	EncodeTime(key string, t time.Time)
	// EncodeUint encodes a field with the given key and unsigned integer value.
	EncodeUint(key string, i uint)
	// EncodeUint32 encodes a field with the given key and uint32 value.
	EncodeUint32(key string, i uint32)
	// EncodeUint64 encodes a field with the given key and uint64 value.
	EncodeUint64(key string, i uint64)
}
