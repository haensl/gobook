package charcount

import (
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

type CharAnalysis struct {
	counts  *map[rune]int
	utflen  [utf8.UTFMax + 1]int
	invalid int
}

func (c *CharAnalysis) String() string {
	return fmt.Sprintf("counts: %d\tutflen: %d\tinvalid: %d", len(*c.counts), c.utflen, c.invalid)
}

func Charcount(in io.RuneReader) (*CharAnalysis, error) {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	return &CharAnalysis{&counts, utflen, invalid}, nil
}
