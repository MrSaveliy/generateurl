package service

import (
	"golang.org/x/text/unicode/bidi"
	"math"
	"strings"
)

const (
	DefaultAlphabetLength uint64 = 63
	DefaultShortLength    int    = 10
)

func getDefaultAlphabet() [DefaultAlphabetLength]string {
	return [DefaultAlphabetLength]string{
		"a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x",
		"y", "z", "A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V",
		"W", "X", "Y", "Z", "1", "2", "3", "4",
		"5", "6", "7", "8", "9", "0", "_",
	}
}

func NewUrlGenerator() *UrlGenerator {
	return &UrlGenerator{
		shortLength:    DefaultShortLength,
		alphabetLength: DefaultAlphabetLength,
		alphabet:       getDefaultAlphabet(),
	}
}

type UrlGenerator struct {
	alphabet       [DefaultAlphabetLength]string
	alphabetLength uint64
	shortLength    int
}

func (s *UrlGenerator) GetShortByID(id uint64) string {
	char := uint64(0)

	index := id + uint64(math.Pow(float64(s.alphabetLength), float64(s.shortLength)))

	beta := new(strings.Builder)

	defer beta.Reset()

	limit := 0
	for {
		part := char + (index % s.alphabetLength)
		beta.WriteString(s.alphabet[int(part)])
		index = (index / s.alphabetLength) >> 0
		limit++

		if index < 1 || limit >= s.shortLength {
			break
		}
	}

	return bidi.ReverseString(beta.String())
}
