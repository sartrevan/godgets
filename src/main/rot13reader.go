package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (z rot13Reader) Read(b []byte) (int, error) {

	a := make([]byte, 10, 10)
	n, e := z.r.Read(a)

	for i := 0; i < n; i++ {
		if a[i] >= 'A' && a[i] <= 'Z' {
			dc := a[i] + 13
			if dc > 'Z' {
				dc = dc - 26
			}
			b[i] = dc
		} else if a[i] >= 'a' && a[i] <= 'z' {
			dc := a[i] + 13
			if dc > 'z' {
				dc = dc - 26
			}
			b[i] = dc
		} else {
			b[i] = a[i]
		}

	}

	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
