package caesarcipher

import (
	"errors"
	"unicode"
)

func Encrypt(plane string, key int32) (string, error) {
	encrypted := make([]rune, len([]rune(plane)))
	for i, s := range plane {
		if !unicode.IsLower(s) {
			return "", errors.New("parameters must be lowercase")
		}

		if s +key > 'z' {
			encrypted[i] = 'a' + s +key- 'z' - 1
			continue
		}
		encrypted[i] = s +key
	}
	return string(encrypted), nil
}

func Decrypt(encrypted string, key int32) (string, error) {
	decrypted := make([]rune, len([]rune(encrypted)))
	for i, s := range encrypted {
		if !unicode.IsLower(s) {
			return "", errors.New("parameters must be lowercase")
		}

		if s - key < 'a' {
			decrypted[i] = 'z' + s - key - 'a' + 1
			continue
		}
		decrypted[i] = s - key
	}
	return string(decrypted), nil
}
