package bytesize_test

import (
	"fmt"
	"testing"

	"github.com/Andrew-M-C/go-bytesize"
	"github.com/smartystreets/goconvey/convey"
)

var (
	cv = convey.Convey
	so = convey.So

	eq = convey.ShouldEqual
)

func TestByteSize(t *testing.T) {
	cv("Comma", t, func() { testComma(t) })
	cv("Base2Full", t, func() { testBase2Full(t) })
	cv("Base2", t, func() { testBase2(t) })
	cv("Base10Full", t, func() { testBase10Full(t) })
	cv("Base10", t, func() { testBase10(t) })
}

func testComma(t *testing.T) {
	n := -1234567890
	s := fmt.Sprint(bytesize.Comma(n))
	so(s, eq, "-1,234,567,890")

	u := 1234567
	s = fmt.Sprint(bytesize.Comma(u))
	so(s, eq, "1,234,567")
}

func testBase2Full(t *testing.T) {
	s := fmt.Sprint(bytesize.Base2Full(0))
	so(s, eq, "0 byte")

	s = fmt.Sprint(bytesize.Base2Full(1024))
	so(s, eq, "1 KiB")

	s = fmt.Sprint(bytesize.Base2Full(0xFFFFFFFFFFFFFFFF))
	so(s, eq, "15 EiB 1023 PiB 1023 TiB 1023 GiB 1023 MiB 1023 KiB 1023 bytes")

	// TODO:
}

func testBase2(t *testing.T) {
	s := fmt.Sprint(bytesize.Base2(0))
	so(s, eq, "0 byte")

	s = fmt.Sprint(bytesize.Base2(1024))
	so(s, eq, "1 KiB")

	s = fmt.Sprint(bytesize.Base2(0xFFFFFFFFFFFFFFFF))
	so(s, eq, "15.999 EiB")

	s = fmt.Sprint(bytesize.Base2(15 << 60))
	so(s, eq, "15 EiB")

	s = fmt.Sprint(bytesize.Base2(1024 + 1))
	so(s, eq, "1.001 KiB")

	s = fmt.Sprint(bytesize.Base2(1024*1024 + 1))
	so(s, eq, "1.000 MiB")

	s = fmt.Sprint(bytesize.Base2(1024*1024 + 1024))
	so(s, eq, "1.001 MiB")
}

func testBase10Full(t *testing.T) {
	s := fmt.Sprint(bytesize.SIFull(0))
	so(s, eq, "0 byte")

	s = fmt.Sprint(bytesize.SIFull(1000))
	so(s, eq, "1 KB")

	s = fmt.Sprint(bytesize.Base10Full(0xFFFFFFFFFFFFFFFF))
	so(s, eq, "18 EB 446 PB 744 TB 73 GB 709 MB 551 KB 615 bytes")

	// TODO:
}

func testBase10(t *testing.T) {
	s := fmt.Sprint(bytesize.Base10(0))
	so(s, eq, "0 byte")

	s = fmt.Sprint(bytesize.Base10(1000))
	so(s, eq, "1 KB")
	s = fmt.Sprint(bytesize.Base10(1024))
	so(s, eq, "1.024 KB")

	s = fmt.Sprint(bytesize.Base10(0xFFFFFFFFFFFFFFFF))
	so(s, eq, "18.446 EB")

	s = fmt.Sprint(bytesize.Base10(15000000000000000000))
	so(s, eq, "15 EB")

	s = fmt.Sprint(bytesize.Base10(1000 + 1))
	so(s, eq, "1.001 KB")

	s = fmt.Sprint(bytesize.Base10(1000*1000 + 1))
	so(s, eq, "1.000 MB")

	s = fmt.Sprint(bytesize.Base10(1000*1000 + 1000))
	so(s, eq, "1.001 MB")
}
