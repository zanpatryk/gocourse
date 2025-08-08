package iolimit

import "io"

type limited struct {
	r io.Reader
	n int64 // bytes left to read
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limited{r: r, n: n}
}

func (l *limited) Read(p []byte) (int, error) {
	// if no bytes left, return
	if l.n <= 0 {
		return 0, io.EOF
	}

	// read at most l.n bytes
	if int64(len(p)) > l.n {
		p = p[:l.n]
	}

	n, err := l.r.Read(p)

	// decrease remainig bytes by how many were read
	if n > 0 {
		l.n -= int64(n)
	}

	if err != nil {
		return n, err
	}

	// signal EOF to caller, if exhausted the limit
	if l.n <= 0 {
		return n, io.EOF
	}

	return n, nil
}
