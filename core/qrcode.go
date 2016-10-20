package core

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func Create_qrcode(data string) barcode.Barcode {

	code, err := qr.Encode(data, qr.L, qr.Unicode)
	HandleError(err)

	if data != code.Content() {
		panic("Create_qrcode data differs")
	}

	code, err = barcode.Scale(code, 300, 300)
	HandleError(err)

	return code
}
