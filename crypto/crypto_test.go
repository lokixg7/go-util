package crypto

import "testing"

// go test -v ./crypto -run '^TestMD5$'
func TestMD5(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{name: "hello", text: "hello", want: "5d41402abc4b2a76b9719d911017c592"},
		{name: "empty", text: "", want: "d41d8cd98f00b204e9800998ecf8427e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.text); got != tt.want {
				t.Errorf("MD5() = %q, want %q", got, tt.want)
			}
		})
	}
}

// go test -v ./crypto -run '^TestSHA1$'
func TestSHA1(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{name: "hello", text: "hello", want: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{name: "empty", text: "", want: "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1(tt.text); got != tt.want {
				t.Errorf("SHA1() = %q, want %q", got, tt.want)
			}
		})
	}
}

// go test -v ./crypto -run '^TestSHA256$'
func TestSHA256(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{name: "hello", text: "hello", want: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{name: "empty", text: "", want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA256(tt.text); got != tt.want {
				t.Errorf("SHA256() = %q, want %q", got, tt.want)
			}
		})
	}
}

// go test -v ./crypto -run '^TestHMACSHA256$'
func TestHMACSHA256(t *testing.T) {
	tests := []struct {
		name string
		key  string
		text string
		want string
	}{
		{
			name: "rfc-4231-test-case-1",
			key:  "\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b",
			text: "Hi There",
			want: "b0344c61d8db38535ca8afceaf0bf12b881dc200c9833da726e9376c2e32cff7",
		},
		{
			name: "empty",
			key:  "",
			text: "",
			want: "b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HMACSHA256(tt.key, tt.text); got != tt.want {
				t.Errorf("HMACSHA256() = %q, want %q", got, tt.want)
			}
		})
	}
}

// go test -v ./crypto -run '^TestBase64Encode$'
func TestBase64Encode(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{name: "hello", text: "hello", want: "aGVsbG8="},
		{name: "empty", text: "", want: ""},
		{name: "unicode", text: "你好", want: "5L2g5aW9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Encode(tt.text); got != tt.want {
				t.Errorf("Base64Encode() = %q, want %q", got, tt.want)
			}
		})
	}
}

// go test -v ./crypto -run '^TestBase64Decode$'
func TestBase64Decode(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    string
		wantErr bool
	}{
		{name: "hello", text: "aGVsbG8=", want: "hello"},
		{name: "empty", text: "", want: ""},
		{name: "unicode", text: "5L2g5aW9", want: "你好"},
		{name: "invalid", text: "!!!not-base64!!!", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64Decode(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Base64Decode() = %q, want %q", got, tt.want)
			}
		})
	}
}
