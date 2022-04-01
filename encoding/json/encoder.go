package json

import (
	"time"

	"github.com/outsidedigital/logger/buffer"
)

// Encoder represents a logging json encoder.
type Encoder struct {
	buf *buffer.Buffer
	n   int
}

// NewEncoder creates a new json encoder that writes to the given buffer.
func NewEncoder(buf *buffer.Buffer) *Encoder {
	return &Encoder{buf: buf, n: 0}
}

// EncodeBool encodes a field with the given key and boolean value.
func (enc *Encoder) EncodeBool(key string, b bool) {
	enc.appendKey(key)
	enc.buf.AppendBool(b)
}

// EncodeBytes encodes a field with the given key and bytes value.
func (enc *Encoder) EncodeBytes(key string, p []byte) {
	enc.appendKey(key)
	l := len(p)
	if l == 0 {
		enc.appendNull()
		return
	}
	enc.buf.AppendByte('"')
	for i := 0; i < l; i++ {
		enc.buf.AppendUint(uint64(p[i]), 64)
	}
	enc.buf.AppendByte('"')
}

// EncodeDuration encodes a field with the given key and duration value.
func (enc *Encoder) EncodeDuration(key string, d time.Duration) {
	enc.appendKey(key)
	enc.buf.AppendDuration(d)
}

// EncodeError encodes a field with the given key and error value.
func (enc *Encoder) EncodeError(key string, err error) {
	enc.appendKey(key)
	if err == nil {
		enc.appendNull()
		return
	}
	enc.buf.AppendQuote(err.Error())
}

// EncodeFloat32 encodes a field with the given key and float32 value.
func (enc *Encoder) EncodeFloat32(key string, f float32) {
	enc.appendKey(key)
	enc.buf.AppendFloat(float64(f), 32)
}

// EncodeFloat64 encodes a field with the given key and float64 value.
func (enc *Encoder) EncodeFloat64(key string, f float64) {
	enc.appendKey(key)
	enc.buf.AppendFloat(f, 64)
}

// EncodeInt encodes a field with the given key and integer value.
func (enc *Encoder) EncodeInt(key string, i int) {
	enc.appendKey(key)
	enc.buf.AppendInt(int64(i), 10)
}

// EncodeInt32 encodes a field with the given key and int32 value.
func (enc *Encoder) EncodeInt32(key string, i int32) {
	enc.appendKey(key)
	enc.buf.AppendInt(int64(i), 10)
}

// EncodeInt64 encodes a field with the given key and int64 value.
func (enc *Encoder) EncodeInt64(key string, i int64) {
	enc.appendKey(key)
	enc.buf.AppendInt(i, 10)
}

// EncodeString encodes a field with the given key and string value.
func (enc *Encoder) EncodeString(key, s string) {
	enc.appendKey(key)
	enc.buf.AppendQuote(s)
}

// EncodeTime encodes a field with the given key and time value.
func (enc *Encoder) EncodeTime(key string, t time.Time) {
	enc.appendKey(key)
	enc.buf.AppendByte('"')
	enc.buf.AppendTime(t, time.RFC3339)
	enc.buf.AppendByte('"')
}

// EncodeUint encodes a field with the given key and unsigned integer value.
func (enc *Encoder) EncodeUint(key string, i uint) {
	enc.appendKey(key)
	enc.buf.AppendUint(uint64(i), 10)
}

// EncodeUint32 encodes a field with the given key and uint32 value.
func (enc *Encoder) EncodeUint32(key string, i uint32) {
	enc.appendKey(key)
	enc.buf.AppendUint(uint64(i), 10)
}

// EncodeUint64 encodes a field with the given key and uint64 value.
func (enc *Encoder) EncodeUint64(key string, i uint64) {
	enc.appendKey(key)
	enc.buf.AppendUint(i, 10)
}

func (enc *Encoder) appendKey(key string) {
	if enc.n > 0 {
		enc.buf.AppendByte(',')
	}
	enc.buf.AppendQuote(key)
	enc.buf.AppendByte(':')
	enc.n++
}

func (enc *Encoder) appendNull() {
	enc.buf.AppendString("null")
}
