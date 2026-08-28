package maps

import (
	"reflect"
	"sort"
	"testing"
)

// go test -v ./map -run '^TestMap2List$'
func TestMap2List(t *testing.T) {
	type args struct {
		m interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "string-map-values",
			args: args{
				m: map[string]string{"a": "1", "b": "2", "c": "3"},
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "empty-map",
			args: args{
				m: map[string]string{},
			},
			want: []string{},
		},
		{
			name: "not-a-map",
			args: args{
				m: []string{"1", "2"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ret []string
			if got := Map2List(tt.args.m, &ret); (got != nil) != tt.wantErr {
				t.Errorf("Map2List() error = %v, wantErr %v", got, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			sort.Strings(ret)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("Map2List() = %v, want %v", ret, tt.want)
			}
			t.Log(ret)
		})
	}
}
