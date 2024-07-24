// 还记得之前编写的图片生成器 吗？我们再来编写另外一个，不过这次它将会返回一个 image.Image 的实现而非一个数据切片。

// 定义你自己的 Image 类型，实现必要的方法并调用 pic.ShowImage。

// Bounds 应当返回一个 image.Rectangle ，例如 image.Rect(0, 0, w, h)。

// ColorModel 应当返回 color.RGBAModel。

// At 应当返回一个颜色。上一个图片生成器的值 v 对应于此次的 color.RGBA{v, v, 255, 255}。

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	v := x * 255 / (i.width - 1)
	return color.RGBA{uint8(v), uint8(v), 255, 255}
}

func main_image() {
	m := Image{width: 100, height: 100}
	pic.ShowImage(m)
}
