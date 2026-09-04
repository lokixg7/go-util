package structs

import (
	"reflect"
	"testing"
)

// go test -v ./structs -run '^TestSetStructFields$'
func TestSetStructFields(t *testing.T) {
	type foo struct {
		Name string
		Age  int64
	}

	type args struct {
		s           interface{}
		fieldValues map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "set-structs-fields",
			args: args{
				s: &foo{
					Name: "Andrew",
					Age:  26,
				},
				fieldValues: map[string]interface{}{
					"Name": "Blue",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetStructFields(tt.args.s, tt.args.fieldValues); (err != nil) != tt.wantErr {
				t.Errorf("SetStructFields() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("new structs = %+v", tt.args.s)
		})
	}
}

// go test -v ./structs -run '^TestConvertToMap$'
func TestConvertToMap(t *testing.T) {
	type User struct {
		Id   uint64
		Name string
	}
	type args struct {
		s   interface{}
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    map[interface{}]interface{}
		wantErr bool
	}{
		{
			name: "elem struct value",
			args: args{
				s: []User{
					{
						Id:   1,
						Name: "blue",
					},
					{
						Id:   2,
						Name: "crank",
					},
				},
				key: "Id",
			},
			want: map[interface{}]interface{}{
				uint64(1): User{Id: 1, Name: "blue"},
				uint64(2): User{Id: 2, Name: "crank"},
			},
		},
		{
			name: "elem struct pointer",
			args: args{
				s: []*User{
					{
						Id:   1,
						Name: "blue",
					},
					{
						Id:   2,
						Name: "crank",
					},
				},
				key: "Id",
			},
			want: map[interface{}]interface{}{
				uint64(1): &User{Id: 1, Name: "blue"},
				uint64(2): &User{Id: 2, Name: "crank"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToMap(tt.args.s, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToMap() got = %v, want %v", got, tt.want)
			}
			if tt.name == "elem struct pointer" {
				t.Log(got[uint64(2)].(*User))
			}
		})
	}
}
