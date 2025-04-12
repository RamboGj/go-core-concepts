package readAndWriters

import (
	"errors"
	"io"
	"os"
	"strings"
)

/*
	Readers and Writers
	io.Reader: read data in streams and save it inside buffers(param p []byte)
	io.Writer: get slice of bytes (p []byte) and save it inside our stream
*/

func main() {
	str := "Hello world"

	reader := strings.NewReader(str)
	writer := MyWriter{os.Stdout}

	buffer := make([]byte, 2)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			panic(err)
		}

		_, _ = writer.Write(buffer[:n])
	}
}

type MyWriter struct {
	r io.Writer
}

func (mw MyWriter) Write(b []byte) (int, error) {
	for i, bb := range b {
		b[i] = bb + 10
	}

	return mw.r.Write(b) 
}