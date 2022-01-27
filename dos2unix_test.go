package dos2unix

import (
	"bytes"
	"io"
	"testing"
)

func TestDOS2Unix(t *testing.T) {
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
			io.CopyBuffer(&output, DOS2Unix(bytes.NewReader(input)), buf[:i])
			if !bytes.Equal(output.Bytes(), test.Output) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}

func TestUnix2DOS(t *testing.T) {
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
			io.CopyBuffer(&output, Unix2DOS(bytes.NewReader(input)), buf[:i])
			if !bytes.Equal(output.Bytes(), test.Output) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}
