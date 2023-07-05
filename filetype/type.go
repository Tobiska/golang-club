package main

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
)

func main() {

	xlsMime := mimetype.Lookup("application/x-ole-storage")

	xlsDetector := func(buf []byte, limit uint32) bool {
		if len(buf) < 8 {
			return false
		}
		return buf[0] == 0xD0 && buf[1] == 0xCF &&
			buf[2] == 0x11 && buf[3] == 0xE0 &&
			buf[4] == 0xA1 && buf[5] == 0xB1 &&
			buf[6] == 0x1A && buf[7] == 0xE1
	}

	xlsMime.Extend(xlsDetector, "application/xls", ".xls", "application/msexcel")

	kind, err := mimetype.DetectFile("filetype/file/prices.xls")
	fmt.Println(kind, err)
}
