// 练习：rot13Reader
// 有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。

// 例如，gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）并返回一个同样实现了 io.Reader 的 *gzip.Reader（解压后的数据流）。

// 编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，通过应用 rot13 代换密码对数据流进行修改。

// rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader。

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		b := p[i]

		if 'A' <= b && b <= 'Z' {
			p[i] = 'A' + (b-'A'+13)%26
		} else if 'a' <= b && b <= 'z' {
			p[i] = 'a' + (b-'a'+13)%26
		}
	}
	return n, nil
}

func main_rot13() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
