package bytesize

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Comma formats decimal digits into comma-separated format. Like 123,456,789
//
// Reference:
//   - [StackOverflow](https://stackoverflow.com/questions/13020308/how-to-fmt-printf-an-integer-with-thousands-comma)
type Comma int64

// String implements fmt.Springer
func (c Comma) String() string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", c)
}
