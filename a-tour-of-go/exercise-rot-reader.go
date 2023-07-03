package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		// 別解:ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ という文字列を作ると場合分けがなくなる
		if b[i] >= 'A' && b[i] <= 'Z' {
			if b[i]+13 >= 'Z' {
				b[i] += 'A' + (b[i] + 13 - 'Z' - 1)
			} else {
				b[i] += 13
			}
		} else if b[i] >= 'a' && b[i] <= 'z' {
			if b[i]+13 >= 'z' {
				b[i] = 'a' + (b[i] + 13 - 'z' - 1)
			} else {
				b[i] += 13
			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
