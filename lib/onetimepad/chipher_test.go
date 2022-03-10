package onetimepad

import (
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		plane []byte
		key   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"string",
			args{
				[]byte("hogeHOGE"),
				[]byte("hugaHUGA"),
			},
			[]byte{0, 26, 0, 4, 0, 26, 0, 4},
			false,
		},
		{
			"include symbols",
			args{
				[]byte("hoge!@#$%$"),
				[]byte("HUGAhuga!@"),
			},
			[]byte{32, 58, 32, 36, 73, 53, 68, 69, 4, 100},
			false,
		},
		{
			"include emoji",
			args{
				[]byte("✊"),
				[]byte{12, 55, 66},
			},
			[]byte{238, 171, 200},
			false,
		},
		{
			"key length is different",
			args{
				[]byte("hogehoge"),
				[]byte("huga"),
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.plane, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		encrypted []byte
		key       []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"string",
			args{
				[]byte{0, 26, 0, 4, 0, 26, 0, 4},
				[]byte("hugaHUGA"),
			},
			[]byte("hogeHOGE"),
			false,
		},
		{
			"include symbols",
			args{
				[]byte{32, 58, 32, 36, 73, 53, 68, 69, 4, 100},
				[]byte("HUGAhuga!@"),
			},
			[]byte("hoge!@#$%$"),
			false,
		},
		{
			"include emoji",
			args{
				[]byte{238, 171, 200},
				[]byte{12, 55, 66},
			},
			[]byte("✊"),
			false,
		},
		{
			"key length is different",
			args{
				[]byte("hogehoge"),
				[]byte("huga"),
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.encrypted, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
