package binary

// https://github.com/Mrs4s/MiraiGo/blob/master/binary/writer.go

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/wdvxdr1123/ZeroBot/utils/helper"
)

// Writer 写入
type Writer bytes.Buffer

//nolint: revive
func NewWriterF(f func(writer *Writer)) []byte {
	w := SelectWriter()
	f(w)
	b := append([]byte(nil), w.Bytes()...)
	w.put()
	return b
}

// OpenWriterF must call func cl to close
//nolint: revive
func OpenWriterF(f func(*Writer)) (b []byte, cl func()) {
	w := SelectWriter()
	f(w)
	return w.Bytes(), w.put
}

//nolint: revive
func (w *Writer) FillUInt16() (pos int) {
	pos = w.Len()
	(*bytes.Buffer)(w).Write([]byte{0, 0})
	return
}

//nolint: revive
func (w *Writer) WriteUInt16At(pos int, v uint16) {
	newdata := (*bytes.Buffer)(w).Bytes()[pos:]
	binary.BigEndian.PutUint16(newdata, v)
}

//nolint: revive
func (w *Writer) FillUInt32() (pos int) {
	pos = w.Len()
	(*bytes.Buffer)(w).Write([]byte{0, 0, 0, 0})
	return
}

//nolint: revive
func (w *Writer) WriteUInt32At(pos int, v uint32) {
	newdata := (*bytes.Buffer)(w).Bytes()[pos:]
	binary.BigEndian.PutUint32(newdata, v)
}

//nolint: revive
func (w *Writer) Write(b []byte) (n int, err error) {
	return (*bytes.Buffer)(w).Write(b)
}

//nolint: revive
func (w *Writer) WriteHex(h string) {
	b, _ := hex.DecodeString(h)
	_, _ = w.Write(b)
}

//nolint: revive
func (w *Writer) WriteByte(b byte) error {
	return (*bytes.Buffer)(w).WriteByte(b)
}

//nolint: revive
func (w *Writer) WriteUInt16(v uint16) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	_, _ = w.Write(b)
}

//nolint: revive
func (w *Writer) WriteUInt32(v uint32) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	_, _ = w.Write(b)
}

//nolint: revive
func (w *Writer) WriteUInt64(v uint64) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	_, _ = w.Write(b)
}

//nolint: revive
func (w *Writer) WriteUInt16LE(v uint16) {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	_, _ = w.Write(b)
}

//nolint: revive
func (w *Writer) WriteUInt32LE(v uint32) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	_, _ = w.Write(b)
}

//nolint: revive
func (w *Writer) WriteUInt64LE(v uint64) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, v)
	_, _ = w.Write(b)
}

//nolint: revive
func (w *Writer) WriteString(v string) {
	// w.WriteUInt32(uint32(len(v) + 4))
	(*bytes.Buffer)(w).WriteString(v)
}

/*
//nolint: revive
func (w *Writer) WriteStringShort(v string) {
	w.WriteUInt16(uint16(len(v)))
	(*bytes.Buffer)(w).WriteString(v)
}
*/

//nolint: revive
func (w *Writer) WriteBool(b bool) {
	if b {
		_ = w.WriteByte(0x01)
	} else {
		_ = w.WriteByte(0x00)
	}
}

/*
//nolint: revive
func (w *Writer) WriteBytesShort(data []byte) {
	w.WriteUInt16(uint16(len(data)))
	w.Write(data)
}
*/

//nolint: revive
func (w *Writer) Len() int {
	return (*bytes.Buffer)(w).Len()
}

//nolint: revive
func (w *Writer) Bytes() []byte {
	return (*bytes.Buffer)(w).Bytes()
}

//nolint: revive
func (w *Writer) String() string {
	return helper.BytesToString((*bytes.Buffer)(w).Bytes())
}

//nolint: revive
func (w *Writer) Reset() {
	(*bytes.Buffer)(w).Reset()
}

//nolint: revive
func (w *Writer) Grow(n int) {
	(*bytes.Buffer)(w).Grow(n)
}

func (w *Writer) put() {
	PutWriter(w)
}
