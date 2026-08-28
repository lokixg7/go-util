package maps

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v ./maps -run '^TestMap2List$'
func TestMap2List(t *testing.T) {
	type args struct {
		m      interface{}
		refRet interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "string-map-values",
			args: args{
				m:      map[string]string{"a": "1", "b": "2", "c": "3"},
				refRet: new([]string),
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "int-map-values",
			args: args{
				m:      map[string]int{"a": 1, "b": 2, "c": 3},
				refRet: new([]int),
			},
			want: []int{1, 2, 3},
		},
		{
			name: "non-string-keys",
			args: args{
				m:      map[int]string{1: "a", 2: "b"},
				refRet: new([]string),
			},
			want: []string{"a", "b"},
		},
		{
			name: "empty-map",
			args: args{
				m:      map[string]string{},
				refRet: new([]string),
			},
			want: []string{},
		},
		{
			name: "nil-map",
			args: args{
				m:      map[string]string(nil),
				refRet: new([]string),
			},
			want: []string{},
		},
		{
			name: "not-a-map",
			args: args{
				m:      []string{"1", "2"},
				refRet: new([]string),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Map2List(tt.args.m, tt.args.refRet)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.ElementsMatch(t, tt.want, reflect.ValueOf(tt.args.refRet).Elem().Interface())
		})
	}
}
