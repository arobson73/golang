package main

import (
	"fmt"
	"io"
	"strings"
)

type LimReader struct {
	R io.Reader // underlying reader
	N int64     // max bytes remaining
}

func LReader(r io.Reader, n int64) *LimReader {
	var lm LimReader
	lm.R = r
	lm.N = n
	return &lm
}

func (l *LimReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func main() {
	data := ` 
<!DOCTYPE html>
<html>
<body>

<p>The image is a link. You can click on it.</p>

<a href="default.asp">
  <img src="smiley.gif" alt="HTML tutorial" style="width:42px;height:42px;border:0">
</a>

<p>We have added "border:0" to prevent IE9 (and earlier) from displaying a border around the image.</p>

</body>
</html>
`
	lr := LReader(strings.NewReader(data), 50)
	buf := make([]byte, 2048)
	n, err := lr.Read(buf)
	buf = buf[:n]
	fmt.Println(n, err, string(buf))

}
