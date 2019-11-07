// Package copy provides ...
package copier

import (
	"io"
)

type Copier struct {
	src io.Reader
	dst io.Writer
}

func NewCopier(src io.Reader, dst io.Writer) *Copier {
	return &Copier{
		src: src,
		dst: dst,
	}
}

func (c Copier) Copy() (written int64, err error) {
	buf := make([]byte, 1024)

	for {
		nr, ok := c.src.Read(buf)
		if ok != nil {
			if ok != io.EOF {
				err = ok
			}
			break
		}
		if nr > 0 {
			nw, ok := c.dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ok != nil {
				err = ok
				break
			}
		}
	}

	return written, err
}
