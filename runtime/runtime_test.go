package runtime

import (
	"testing"
)

// go test -v ./runtime -run '^TestGetCurCalleeFunc$'
func TestGetCurCalleeFunc(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "func1",
			want: "func1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurCalleeFunc(); got != tt.want {
				t.Errorf("GetCurCalleeFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v ./runtime -run '^TestGetParentCallFunc$'
func TestGetParentCallFunc(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "tRunner",
			want: "tRunner",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetParentCallFunc(); got != tt.want {
				t.Errorf("GetParentCallFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
