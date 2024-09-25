package strings

import (
	"regexp"
	"strings"
	"unicode"
)

type StringCleaner interface {
	KeepFirstNChars(n int) StringCleaner
	RemoveNonLetters() StringCleaner
	RemoveWhiteSpace() StringCleaner
	ToLower() StringCleaner
	String() string
}

type stringCleaner struct {
	str string
}

func NewStringCleaner(str string) StringCleaner {
	return &stringCleaner{str}
}

func (sc *stringCleaner) KeepFirstNChars(n int) StringCleaner {
	if len(sc.str) > (n-1) {
		sc.str = sc.str[0:(n-1)]
	}
	return sc
}

func (sc *stringCleaner) RemoveNonLetters() StringCleaner {
	var builder strings.Builder
	for _, r := range sc.str {
		if unicode.IsLetter(r) || r == ' ' {
			builder.WriteRune(r)
		}
	}
	sc.str = builder.String()
	return sc
}

func (sc *stringCleaner) RemoveWhiteSpace() StringCleaner {
	sc.str = strings.Trim(sc.str, " ")
	re, _ := regexp.Compile(`\s+`)
	sc.str = re.ReplaceAllString(sc.str, " ")
	return sc
}

func (sc *stringCleaner) ToLower() StringCleaner {
	sc.str = strings.ToLower(sc.str)
	return sc
}

func (sc *stringCleaner) String() string {
	return sc.str
}