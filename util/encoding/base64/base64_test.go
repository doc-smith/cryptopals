package base64

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{[]byte{}}, ""},
		{"no-padding", args{[]byte{0x12, 0xaa, 0x7f}}, "Eqp/"},
		{"single-byte", args{[]byte{0x12}}, "Eg=="},
		{"padding", args{[]byte{0x12, 0xaa}}, "Eqo="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.b); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"empty", args{""}, nil, false},
		{"no-padding", args{"Eqp/"}, []byte{0x12, 0xaa, 0x7f}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
