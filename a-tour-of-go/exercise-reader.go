package main

import (
	"golang.org/x/tour/reader"
)

// / https://cs.opensource.google/go/x/tour/+/refs/tags/v0.1.0:reader/validate.go;l=13
// func Validate(r io.Reader) {
// 	b := make([]byte, 1024, 2048)
// 	i, o := 0, 0
// 	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
// 		n, err := r.Read(b)
// 		for i, v := range b[:n] {
// 			if v != 'A' {
// 				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
// 				return
// 			}
// 		}
// 		o += n
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
// 			return
// 		}
// 	}
// 	if o == 0 {
// 		fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
// 		return
// 	}
// 	fmt.Println("OK!")
// }

// Reader is the interface that wraps the basic Read method.
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
