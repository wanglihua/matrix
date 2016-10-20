package core

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func Test_create_qrcode(t *testing.T) {
	code := Create_qrcode("123456")
	write_png("z:/qrcode_123456.png", code)
}

func write_png(file_name string, img image.Image) {
	file, err := os.Create(file_name)
	HandleError(err)

	err = png.Encode(file, img)
	// err = jpeg.Encode(file, img, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	HandleError(err)

	file.Close()
}
