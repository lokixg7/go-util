package strings

import "testing"

// go test -v ./strings -run '^TestSubstr$'
func TestSubstr(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		start  int
		length int
		want   string
	}{
		{
			name:   "from-start",
			s:      "hello world",
			start:  0,
			length: 5,
			want:   "hello",
		},
		{
			name:   "from-middle",
			s:      "hello world",
			start:  6,
			length: 5,
			want:   "world",
		},
		{
			name:   "start-beyond-end",
			s:      "hello",
			start:  10,
			length: 3,
			want:   "",
		},
		{
			name:   "negative-start",
			s:      "hello world",
			start:  -5,
			length: 5,
			want:   "world",
		},
		{
			name:   "negative-start-beyond",
			s:      "hello",
			start:  -10,
			length: 2,
			want:   "he",
		},
		{
			name:   "length-clamped",
			s:      "hello",
			start:  2,
			length: 10,
			want:   "llo",
		},
		{
			name:   "negative-length",
			s:      "hello",
			start:  1,
			length: -1,
			want:   "",
		},
		{
			name:   "unicode",
			s:      "你好世界",
			start:  1,
			length: 2,
			want:   "好世",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substr(tt.s, tt.start, tt.length); got != tt.want {
				t.Errorf("Substr() = %q, want %q", got, tt.want)
			}
		})
	}
}

// go test -v ./strings -run '^TestTruncate$'
func TestTruncate(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		maxLen int
		suffix string
		want   string
	}{
		{
			name:   "no-truncation-needed",
			s:      "hello",
			maxLen: 10,
			suffix: "...",
			want:   "hello",
		},
		{
			name:   "truncate-with-suffix",
			s:      "hello world",
			maxLen: 8,
			suffix: "...",
			want:   "hello...",
		},
		{
			name:   "suffix-longer-than-max",
			s:      "hello world",
			maxLen: 3,
			suffix: "....",
			want:   "hel",
		},
		{
			name:   "zero-max-len",
			s:      "hello world",
			maxLen: 0,
			suffix: "...",
			want:   "",
		},
		{
			name:   "empty-suffix",
			s:      "hello world",
			maxLen: 5,
			suffix: "",
			want:   "hello",
		},
		{
			name:   "unicode",
			s:      "你好世界",
			maxLen: 3,
			suffix: "…",
			want:   "你好…",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.s, tt.maxLen, tt.suffix); got != tt.want {
				t.Errorf("Truncate() = %q, want %q", got, tt.want)
			}
		})
	}
}

// go test -v ./strings -run '^TestSnakeCase$'
func TestSnakeCase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "pascal-case", s: "CamelCase", want: "camel_case"},
		{name: "camel-case", s: "camelCase", want: "camel_case"},
		{name: "acronym", s: "HTTPServer", want: "http_server"},
		{name: "already-snake", s: "hello_world", want: "hello_world"},
		{name: "digit-boundary", s: "Foo2Bar", want: "foo2_bar"},
		{name: "empty", s: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.s); got != tt.want {
				t.Errorf("SnakeCase() = %q, want %q", got, tt.want)
			}
		})
	}
}

// go test -v ./strings -run '^TestCamelCase$'
func TestCamelCase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "snake-case", s: "hello_world", want: "helloWorld"},
		{name: "pascal-case", s: "HelloWorld", want: "helloWorld"},
		{name: "acronym", s: "HTTPServer", want: "httpServer"},
		{name: "all-caps-snake", s: "HELLO_WORLD", want: "helloWorld"},
		{name: "hyphen", s: "foo-bar", want: "fooBar"},
		{name: "space", s: "hello world", want: "helloWorld"},
		{name: "empty", s: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.s); got != tt.want {
				t.Errorf("CamelCase() = %q, want %q", got, tt.want)
			}
		})
	}
}

// go test -v ./strings -run '^TestPascalCase$'
func TestPascalCase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "snake-case", s: "hello_world", want: "HelloWorld"},
		{name: "camel-case", s: "helloWorld", want: "HelloWorld"},
		{name: "acronym", s: "HTTPServer", want: "HttpServer"},
		{name: "empty", s: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PascalCase(tt.s); got != tt.want {
				t.Errorf("PascalCase() = %q, want %q", got, tt.want)
			}
		})
	}
}

// go test -v ./strings -run '^TestMask$'
func TestMask(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		start    int
		end      int
		maskChar string
		want     string
	}{
		{
			name:  "phone-number",
			s:     "13812345678",
			start: 3,
			end:   7,
			want:  "138****5678",
		},
		{
			name:     "default-mask-char",
			s:        "13812345678",
			start:    3,
			end:      7,
			maskChar: "",
			want:     "138****5678",
		},
		{
			name:  "mask-from-start",
			s:     "hello@example.com",
			start: 0,
			end:   5,
			want:  "*****@example.com",
		},
		{
			name:  "start-out-of-range",
			s:     "abc",
			start: 5,
			end:   6,
			want:  "abc",
		},
		{
			name:  "end-clamped",
			s:     "abcdef",
			start: 2,
			end:   99,
			want:  "ab****",
		},
		{
			name:  "negative-start-clamped",
			s:     "abcdef",
			start: -2,
			end:   3,
			want:  "***def",
		},
		{
			name:  "unicode",
			s:     "你好世界",
			start: 1,
			end:   3,
			want:  "你**界",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mask(tt.s, tt.start, tt.end, tt.maskChar); got != tt.want {
				t.Errorf("Mask() = %q, want %q", got, tt.want)
			}
		})
	}
}
