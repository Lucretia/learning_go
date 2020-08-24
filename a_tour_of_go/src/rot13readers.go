package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(ch byte) byte {
	const offset = 13

	getCharOffset := func(n byte, upper byte) byte {
		return (offset - (upper - ch)) - 1
	}

	rch := ch + offset

	if ch >= 'A' && ch <= 'Z' {
		if rch > 'Z' {
			return 'A' + getCharOffset(ch, 'Z')
		}

		return rch
	} else if ch >= 'a' && ch <= 'z' {
		if rch > 'z' {
			return 'a' + getCharOffset(ch, 'z')
		}

		return rch
	}

	return ch
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)

	// fmt.Println("Start: ", n, err, b[:n], string(b[:n]), len(b))

	if err == nil {
		for i := 0; i < n; i++ {
			b[i] = rot13(b[i])
		}

		// fmt.Println("End: ", n, err, b[:n], string(b[:n]), len(b))

		return len(b), nil
	}

	return 0, err
}

func main() {
	// You cracked the code!
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
