package cipher

import (
	"bytes"
	"testing"

	r "github.com/stretchr/testify/require"
)

func TestRot255Reader_Read(t *testing.T) {
	arr := []byte{254, 253, 252}
	reader, err := NewRot255Reader(bytes.NewBuffer(arr))
	r.NoError(t, err)
	r.NotNil(t, reader)

	n, err := reader.Read(arr)
	r.NoError(t, err)
	r.Equal(t, 3, n)
	r.Equal(t, byte(0x1), arr[0])
	r.Equal(t, byte(0x2), arr[1])
	r.Equal(t, byte(0x3), arr[2])
}

func TestRot255Writer_Write(t *testing.T) {
	buf := &bytes.Buffer{}
	writer, err := NewRot255Writer(buf)
	r.NoError(t, err)
	r.NotNil(t, writer)

	n, err := writer.Write([]byte{1, 2, 3})
	r.NoError(t, err)
	r.Equal(t, 3, n)

	arr := buf.Bytes()
	r.Equal(t, byte(0xFE), arr[0])
	r.Equal(t, byte(0xFD), arr[1])
	r.Equal(t, byte(0xFC), arr[2])
}
