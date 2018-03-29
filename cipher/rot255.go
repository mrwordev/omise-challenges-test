package cipher

import (
	"io"
)

type Rot255Reader struct{ reader io.Reader }

func NewRot255Reader(r io.Reader) (*Rot255Reader, error) {
	return &Rot255Reader{reader: r}, nil
}

func (r *Rot255Reader) Read(p []byte) (int, error) {
	if n, err := r.reader.Read(p); err != nil {
		return n, err
	} else {
		rot255(p[:n])
		return n, nil
	}
}

type Rot255Writer struct {
	writer io.Writer
	buffer []byte // not thread-safe
}

func NewRot255Writer(w io.Writer) (*Rot255Writer, error) {
	return &Rot255Writer{
		writer: w,
		buffer: make([]byte, 255, 255),
	}, nil
}

func (w *Rot255Writer) Write(p []byte) (int, error) {
	n := copy(w.buffer, p)
	rot255(w.buffer[:n])
	return w.writer.Write(w.buffer[:n])
}

func rot255(buf []byte) {
	for idx, b := range buf {
		buf[idx] = 255 - b
	}
}
