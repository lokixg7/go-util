package strings

import (
	"bytes"
	"unicode"
)

// Substr returns the substring of s starting at the rune index start with the
// given rune length. A negative start counts backward from the end of s.
// Out-of-range values are clamped, and a negative length yields an empty
// string.
func Substr(s string, start, length int) string {
	rs := []rune(s)
	if length < 0 {
		return ""
	}
	if start < 0 {
		start = len(rs) + start
	}
	if start < 0 {
		start = 0
	}
	if start >= len(rs) {
		return ""
	}
	if start+length > len(rs) {
		length = len(rs) - start
	}
	return string(rs[start : start+length])
}

// Truncate shortens s to at most maxLen runes, appending suffix when s is
// truncated. If maxLen is not greater than the suffix length, only the first
// maxLen runes are kept.
func Truncate(s string, maxLen int, suffix string) string {
	rs := []rune(s)
	if maxLen <= 0 {
		return ""
	}
	if len(rs) <= maxLen {
		return s
	}
	suffixRunes := []rune(suffix)
	if len(suffixRunes) >= maxLen {
		return string(rs[:maxLen])
	}
	return string(rs[:maxLen-len(suffixRunes)]) + suffix
}

// SnakeCase converts CamelCase, PascalCase, or separated input into
// lower_snake_case.
func SnakeCase(s string) string {
	words := splitWords(s)
	var b bytes.Buffer
	for i, w := range words {
		if i > 0 {
			b.WriteByte('_')
		}
		b.WriteString(toLower(w))
	}
	return b.String()
}

// CamelCase converts input into lower camelCase, splitting words on
// underscores, hyphens, dots, and capitalization boundaries.
func CamelCase(s string) string {
	return toCamelCase(s, false)
}

// PascalCase converts input into UpperCamelCase, splitting words on
// underscores, hyphens, dots, and capitalization boundaries.
func PascalCase(s string) string {
	return toCamelCase(s, true)
}

// Mask replaces the runes in the half-open interval [start, end) with maskChar
// (default "*"). Out-of-range indices are clamped; invalid ranges return s
// unchanged.
func Mask(s string, start, end int, maskChar string) string {
	rs := []rune(s)
	if maskChar == "" {
		maskChar = "*"
	}
	if start < 0 {
		start = 0
	}
	if end > len(rs) {
		end = len(rs)
	}
	if start >= end {
		return s
	}

	var b bytes.Buffer
	b.WriteString(string(rs[:start]))
	for i := start; i < end; i++ {
		b.WriteRune([]rune(maskChar)[0])
	}
	b.WriteString(string(rs[end:]))
	return b.String()
}

func toCamelCase(s string, upperFirst bool) string {
	words := splitWords(s)
	var b bytes.Buffer
	for i, w := range words {
		rs := []rune(toLower(w))
		if len(rs) == 0 {
			continue
		}
		if i > 0 || upperFirst {
			b.WriteRune(unicode.ToUpper(rs[0]))
			b.WriteString(string(rs[1:]))
			continue
		}
		b.WriteString(string(rs))
	}
	return b.String()
}

func toLower(s string) string {
	var b bytes.Buffer
	for _, r := range s {
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

func splitWords(s string) []string {
	rs := []rune(s)
	var (
		words []string
		cur   []rune
	)
	for i, r := range rs {
		if isSeparator(r) {
			if len(cur) > 0 {
				words = append(words, string(cur))
				cur = nil
			}
			continue
		}
		if len(cur) > 0 && isWordBoundary(rs, i) {
			words = append(words, string(cur))
			cur = nil
		}
		cur = append(cur, r)
	}
	if len(cur) > 0 {
		words = append(words, string(cur))
	}
	return words
}

func isWordBoundary(rs []rune, i int) bool {
	prev := rs[i-1]
	r := rs[i]
	if unicode.IsUpper(r) && (unicode.IsLower(prev) || unicode.IsDigit(prev)) {
		return true
	}
	// Split acronyms like "HTTPServer" into "HTTP" and "Server".
	if unicode.IsUpper(r) && unicode.IsUpper(prev) && i+1 < len(rs) && unicode.IsLower(rs[i+1]) {
		return true
	}
	return false
}

func isSeparator(r rune) bool {
	return r == '_' || r == '-' || r == '.' || unicode.IsSpace(r)
}
