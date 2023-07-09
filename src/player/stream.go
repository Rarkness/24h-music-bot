package player

import (
	"encoding/binary"
	"io"
)

func stream(writer io.Writer, reader io.Reader) error {
	buf := [2]byte{}

	for {
		if _, err := io.ReadFull(reader, buf[:]); err != nil && err == io.EOF {
			break
		}

		frame := binary.LittleEndian.Uint16(buf[:])
		if _, err := io.CopyN(writer, reader, int64(frame)); err != nil && err == io.EOF {
			break
		}
	}

	return nil
}
