package simplesubstitution

import "testing"

func TestDecrypt(t *testing.T) {
	encryptionMap := EncryptionMap{
		'a': 'w',
		'b': 'y',
		'c': 'h',
		'd': 'f',
		'e': 'x',
		'f': 'u',
		'g': 'm',
		'h': 't',
		'i': 'j',
		'j': 'v',
		'k': 's',
		'l': 'g',
		'm': 'e',
		'n': 'n',
		'o': 'b',
		'p': 'r',
		'q': 'd',
		'r': 'z',
		's': 'l',
		't': 'q',
		'u': 'a',
		'v': 'p',
		'w': 'c',
		'x': 'o',
		'y': 'k',
		'z': 'i',
	}

	type args struct {
		encrypted string
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
				encrypted: "kbltjsb",
			},
			want:    "yoshiko",
			wantErr: false,
		},
		{
			name: "include uppercase",
			args: args{
				encrypted: "HogeHuga",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "include symbol",
			args: args{
				encrypted: "!@#$%)*",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "include emoji",
			args: args{
				encrypted: "✊👏☔️",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.encrypted, encryptionMap)
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

func TestEncrypt(t *testing.T) {
	encryptionMap := EncryptionMap{
		'a': 'w',
		'b': 'y',
		'c': 'h',
		'd': 'f',
		'e': 'x',
		'f': 'u',
		'g': 'm',
		'h': 't',
		'i': 'j',
		'j': 'v',
		'k': 's',
		'l': 'g',
		'm': 'e',
		'n': 'n',
		'o': 'b',
		'p': 'r',
		'q': 'd',
		'r': 'z',
		's': 'l',
		't': 'q',
		'u': 'a',
		'v': 'p',
		'w': 'c',
		'x': 'o',
		'y': 'k',
		'z': 'i',
	}

	type args struct {
		plain string
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
				"yoshiko",
			},
			want:    "kbltjsb",
			wantErr: false,
		},
		{
			name: "include uppercase",
			args: args{
				"HogeHuga",
			},
			want:    "",
			wantErr: true,
		},
		{
			"include symbol",
			args{
				"!@#$%",
			},
			"",
			true,
		},
		{
			"include emoji",
			args{
				"👏✊☔️",
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.plain, encryptionMap)
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
