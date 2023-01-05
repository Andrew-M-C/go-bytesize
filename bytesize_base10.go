package bytesize

import (
	"fmt"
	"strconv"
	"strings"
)

func base10Parts(n uint64) (parts [7]string) {
	s := strconv.FormatUint(n, 10)
	// max: 18,446,744,073,709,551,615, 20 bytes max
	switch len(s) {
	// case 0 is impossible
	case 1, 2, 3:
		parts[6] = s
	case 4, 5, 6:
		i := len(s) - 3
		parts[5] = s[:i]
		parts[6] = s[i:]
	case 7, 8, 9:
		i := len(s) - 6
		parts[4] = s[:i]
		parts[5] = s[i : i+3]
		parts[6] = s[i+3:]
	case 10, 11, 12:
		i := len(s) - 9
		parts[3] = s[:i]
		parts[4] = s[i : i+3]
		parts[5] = s[i+3 : i+6]
		parts[6] = s[i+6:]
	case 13, 14, 15:
		i := len(s) - 12
		parts[2] = s[:i]
		parts[3] = s[i : i+3]
		parts[4] = s[i+3 : i+6]
		parts[5] = s[i+6 : i+9]
		parts[6] = s[i+9:]
	case 16, 17, 18:
		i := len(s) - 15
		parts[1] = s[:i]
		parts[2] = s[i : i+3]
		parts[3] = s[i+3 : i+6]
		parts[4] = s[i+6 : i+9]
		parts[5] = s[i+9 : i+12]
		parts[6] = s[i+12:]
	default:
		i := len(s) - 18
		parts[0] = s[:i]
		parts[1] = s[i : i+3]
		parts[2] = s[i+3 : i+6]
		parts[3] = s[i+6 : i+9]
		parts[4] = s[i+9 : i+12]
		parts[5] = s[i+12 : i+15]
		parts[6] = s[i+15:]
	}
	return
}

// Base10Full uses KiloByte(KB), MegaByte(MB), GigaByte(GB), TeraByte(TB), PetaByte(PB), ExaByte(EB).
// Showing full parts like 123 MB 456 KB 789 Bytes
type SIFull uint64

type Base10Full = SIFull

// String implements fmt.Stringer
func (n SIFull) String() string {
	if n == 0 {
		return "0 byte"
	}

	parts := base10Parts(uint64(n))
	res := make([]string, 0, 7)

	exact := func(idx int) string {
		s := parts[idx]
		if s == "" {
			return ""
		}
		for i, c := range parts[idx] {
			if c != '0' {
				return parts[idx][i:]
			}
		}
		return ""
	}
	iterate := func(idx int, suffix string) {
		s := exact(idx)
		if s == "" {
			return
		}
		res = append(res, fmt.Sprintf("%s %s", s, suffix))
	}

	iterate(0, "EB")
	iterate(1, "PB")
	iterate(2, "TB")
	iterate(3, "GB")
	iterate(4, "MB")
	iterate(5, "KB")

	if s := exact(6); s == "1" {
		res = append(res, "1 byte")
	} else if s == "" {
		// do nothing
	} else {
		res = append(res, s)
		res = append(res, "bytes")
	}

	return strings.Join(res, " ")
}

// SI uses KiloByte(KB), MegaByte(MB), GigaByte(GB), TeraByte(TB), PetaByte(PB), ExaByte(EB).
// Showing full parts like 123.456 MB
type SI uint64

type Base10 = SI

// String implements fmt.Stringer
func (n SI) String() string {
	if n == 0 {
		return "0 byte"
	}

	parts := base10Parts(uint64(n))

	exact := func(idx int) string {
		s := parts[idx]
		if s == "" {
			return ""
		}
		for i, c := range parts[idx] {
			if c != '0' {
				return parts[idx][i:]
			}
		}
		return ""
	}
	compose := func(itg string, decIndex int, suffix string) string {
		dec := parts[decIndex]
		if dec == "000" {
			for i := decIndex + 1; i < len(parts); i++ {
				if parts[i] != "000" {
					return itg + ".000 " + suffix
				}
			}
			return itg + " " + suffix
		}
		return fmt.Sprintf("%s.%s %s", itg, strings.TrimSuffix(parts[decIndex], "0"), suffix)
	}

	if itg := exact(0); itg != "" {
		return compose(itg, 1, "EB")
	}
	if itg := exact(1); itg != "" {
		return compose(itg, 2, "PB")
	}
	if itg := exact(2); itg != "" {
		return compose(itg, 3, "TB")
	}
	if itg := exact(3); itg != "" {
		return compose(itg, 4, "GB")
	}
	if itg := exact(4); itg != "" {
		return compose(itg, 5, "MB")
	}
	if itg := exact(5); itg != "" {
		return compose(itg, 6, "KB")
	}

	s := exact(6)
	if s == "1" {
		return "1 byte"
	}
	return fmt.Sprintf("%s bytes", s)
}
