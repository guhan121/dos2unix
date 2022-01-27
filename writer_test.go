package dos2unix

import (
	"bytes"
	"io"
	"testing"
)

type jr struct {
	io.Reader
}

func TestDOS2UnixWriter(t *testing.T) {
	buf := make([]byte, 10)
	for n, test := range []struct {
		Input, Output []byte
	}{
		{
			[]byte("Hello"),
			[]byte("Hello"),
		},
		{
			[]byte("Hello\r\n"),
			[]byte("Hello\n"),
		},
		{
			[]byte("Hello\r\nWorld"),
			[]byte("Hello\nWorld"),
		},
		{
			[]byte("qwertyuiop\r\nasdfghjkl\r\nzxcvbnm\r\n"),
			[]byte("qwertyuiop\nasdfghjkl\nzxcvbnm\n"),
		},
		{
			[]byte("qwertyuiop\rasdfgkl\rzxcbnm\r"),
			[]byte("qwertyuiop\rasdfgkl\rzxcbnm\r"),
		},
	} {
		for i := 1; i < 10; i++ {
			var output bytes.Buffer
			input := test.Input
			d := DOS2UnixWriter(&output)
			_, err := io.CopyBuffer(d, jr{bytes.NewReader(input)}, buf[:i])
			if err != nil {
				t.Errorf("test %d.%d: unexpected write error: %s", n+1, i, err)
			} else if err := d.Flush(); err != nil {
				t.Errorf("test %d.%d: unexpected flush error: %s", n+1, i, err)
			} else if !bytes.Equal(output.Bytes(), test.Output) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}

func TestUnix2DOSWriter(t *testing.T) {
	buf := make([]byte, 10)
	for n, test := range []struct {
		Input, Output []byte
	}{
		{
			[]byte("Hello"),
			[]byte("Hello"),
		},
		{
			[]byte("Hello\n"),
			[]byte("Hello\r\n"),
		},
		{
			[]byte("Hello\nWorld"),
			[]byte("Hello\r\nWorld"),
		},
		{
			[]byte("qwertyuiop\nasdfghjkl\nzxcvbnm\n"),
			[]byte("qwertyuiop\r\nasdfghjkl\r\nzxcvbnm\r\n"),
		},
	} {
		for i := 1; i < 10; i++ {
			var output bytes.Buffer
			input := test.Input
			_, err := io.CopyBuffer(Unix2DOSWriter(&output), jr{bytes.NewReader(input)}, buf[:i])
			if err != nil {
				t.Errorf("test %d.%d: unexpected error: %s", n+1, i, err)
			} else if !bytes.Equal(output.Bytes(), test.Output) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}
