package caesarcipher

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		plane string
		key   int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "only lowercase",
			args: args{
				plane: "abcxyz",
				key:   3,
			},
			want:    "defabc",
			wantErr: false,
		},
		{
			name: "include uppercase",
			args: args{
				plane: "ABCabcXYZxyz",
				key:   3,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "include symbol",
			args: args{
				plane: "!@#&*(",
				key:   3,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "include emoji",
			args: args{
				plane: "🙏️✋",
				key:   3,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.plane, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		encrypted string
		key       int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "only lowercase",
			args: args{
				encrypted: "defabc",
				key:       3,
			},
			want:    "abcxyz",
			wantErr: false,
		},
		{
			name: "include uppercase",
			args: args{
				encrypted: "ABChoge",
				key:       3,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "include symbol",
			args: args{
				encrypted: "!@#$%^&*()_+|`",
				key:       3,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "include emoji",
			args: args{
				encrypted: "🙏✊🏖",
				key:       3,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.encrypted, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
