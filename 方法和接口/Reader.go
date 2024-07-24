// 练习：Reader
// 实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 为 MyReader 添加一个 Read([]byte) (int, error) 方法。

func (reader MyReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	b[0] = 'A'
	return 1, nil
}

func main_reader() {
	reader.Validate(MyReader{})
}
