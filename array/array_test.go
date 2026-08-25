package array

import (
	"reflect"
	"testing"
)

// go test -v ./array -run '^TestInArray$'
func TestInArray(t *testing.T) {
	type args struct {
		target interface{}
		arr    interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-true",
			args: args{
				target: int64(7),
				arr:    []int64{1, 3, 5, 7},
			},
			want: true,
		},
		{
			name: "test-false",
			args: args{
				target: int64(10),
				arr:    []int64{1, 3, 5, 7},
			},
			want: false,
		},
		{
			name: "test-string",
			args: args{
				target: "blue",
				arr:    []string{"red", "orange", "yellow", "green", "blue"},
			},
			want: true,
		},
		{
			name: "test-empty",
			args: args{
				target: "blue",
				arr:    []string{},
			},
			want: false,
		},
		{
			name: "test-err-type",
			args: args{
				target: "blue",
				arr:    []int64{1, 2, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v ./array -run '^TestIntersect$'
func TestIntersect(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int64-slice-交集",
			args: args{
				a: []int64{1, 2, 3, 4, 4, 5, 6},
				b: []int64{1, 3, 5, 7, 9, 11},
			},
			want: []int64{1, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ret []int64
			if got := Intersect(tt.args.a, tt.args.b, &ret); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v ./array -run '^TestDiff$'
func TestDiff(t *testing.T) {
	type args struct {
		X interface{}
		Y interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int64-slice-diff",
			args: args{
				X: []int64{1, 2, 3, 4},
				Y: []int64{2, 3, 5, 6},
			},
			want: []int64{1, 4},
		},
		{
			name: "int64-slice-diff",
			args: args{
				X: []int64{2, 3, 5, 6},
				Y: []int64{1, 2, 3, 4},
			},
			want: []int64{5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ret []int64
			if got := Diff(tt.args.X, tt.args.Y, &ret); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v ./array -run '^TestUnique$'
func TestUnique(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "去重-int64-slice",
			args: args{
				slice: []int64{1, 2, 2, 3, 4, 4, 5},
			},
			want: []int64{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ret []int64
			if got := Unique(tt.args.slice, &ret); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
			t.Log(ret)
		})
	}
}

// go test -v ./array -run '^TestExplode$'
func TestExplode(t *testing.T) {
	type args struct {
		delimiter string
		array     interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "join uint64 array",
			args: args{
				delimiter: ",",
				array:     []uint64{1, 2, 3, 4},
			},
			want: "1,2,3,4",
		},
		{
			name: "join int64 array",
			args: args{
				delimiter: ",",
				array:     []int64{1, 2, 3, 4},
			},
			want: "1,2,3,4",
		},
		{
			name: "join string array",
			args: args{
				delimiter: ",",
				array:     []string{"a", "b", "c", "d"},
			},
			want: "a,b,c,d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Explode(tt.args.delimiter, tt.args.array); got != tt.want {
				t.Errorf("Explode() = %v, want %v", got, tt.want)
			}
		})
	}
}
