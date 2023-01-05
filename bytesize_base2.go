package bytesize

import (
	"fmt"
	"strings"
)

func base2Parts(n uint64) (parts [7]uint16) {
	parts[0] = uint16(n & (1024 - 1))
	parts[1] = uint16((n >> 10) & (1024 - 1))
	parts[2] = uint16((n >> 20) & (1024 - 1))
	parts[3] = uint16((n >> 30) & (1024 - 1))
	parts[4] = uint16((n >> 40) & (1024 - 1))
	parts[5] = uint16((n >> 50) & (1024 - 1))
	parts[6] = uint16((n >> 60) & (1024 - 1))
	return
}

// Base2Full uses KibiByte(KiB), MebiByte(MiB), GibiByte(GiB), TebiByte(TiB), PebiByte(PiB), Exbiyte(EiB).
// Showing full parts like 123 MiB 456 KiB 789 Bytes
type Base2Full uint64

// String implements fmt.Stringer
func (n Base2Full) String() string {
	if n == 0 {
		return "0 byte"
	}

	parts := base2Parts(uint64(n))
	res := make([]string, 0, 6)

	iterate := func(part uint16, suffix string) {
		if part == 0 {
			return
		}
		res = append(res, fmt.Sprintf("%d %s", part, suffix))
	}

	iterate(parts[6], "EiB")
	iterate(parts[5], "PiB")
	iterate(parts[4], "TiB")
	iterate(parts[3], "GiB")
	iterate(parts[2], "MiB")
	iterate(parts[1], "KiB")

	if parts[0] == 1 {
		res = append(res, "1 byte")
	} else if parts[0] > 1 {
		res = append(res, fmt.Sprintf("%d bytes", parts[0]))
	}

	return strings.Join(res, " ")
}

// Base2 uses KibiByte(KiB), MebiByte(MiB), GibiByte(GiB), TebiByte(TiB), PebiByte(PiB), Exbiyte(EiB).
// Showing decimal format like 123.456 MiB
type Base2 uint64

// String implements fmt.Stringer
func (n Base2) String() string {
	if n == 0 {
		return "0 byte"
	} else if n == 1 {
		return "1 byte"
	} else if n < 1024 {
		return fmt.Sprintf("%d bytes", n)
	}

	itg := uint64(n >> 60)
	div := uint64(1 << 60)
	dec := uint64(n) & (div - 1)

	shift := func(offset int) {
		itg = uint64(n >> offset)
		div = uint64(1 << offset)
		dec = uint64(n) & (div - 1)
	}

	format := func(suffix string) string {
		if dec == 0 {
			return fmt.Sprintf("%d %s", itg, suffix)
		}
		d := round(float32(dec) / float32(div) * 1000)
		if d >= 1000 {
			return fmt.Sprintf("%d.999 %s", itg, suffix)
		} else if d == 0 {
			return fmt.Sprintf("%d.000 %s", itg, suffix)
		}
		s := fmt.Sprintf("%03d", d)
		return fmt.Sprintf("%d.%s %s", itg, strings.TrimRight(s, "0"), suffix)
	}

	if itg > 0 {
		return format("EiB")
	}

	shift(50)
	if itg > 0 {
		return format("PiB")
	}

	shift(40)
	if itg > 0 {
		return format("TiB")
	}

	shift(30)
	if itg > 0 {
		return format("GiB")
	}

	shift(20)
	if itg > 0 {
		return format("MiB")
	}

	shift(10)
	return format("KiB")
}
