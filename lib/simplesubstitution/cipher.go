package simplesubstitution

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

func Encrypt(plain string, encryptionMap EncryptionMap) (string, error) {
	encrypted := make([]rune, utf8.RuneCountInString(plain))

	for i, r := range plain {
		if !unicode.IsLower(r) {
			return "", errors.New("parameters must be lowercase")
		}

		if v, ok := encryptionMap[r]; ok {
			encrypted[i] = v
		} else {
			return "", errors.New("encryption map doesn't have corresponding rune")
		}
	}
	return string(encrypted), nil
}

func Decrypt(encrypted string, encryptionMap EncryptionMap) (string, error) {
	decrypted := make([]rune, utf8.RuneCountInString(encrypted))

	for i, r := range encrypted {
		if !unicode.IsLower(r) {
			return "", errors.New("parameters must be lowercase")
		}
		k, err := encryptionMap.getKey(r)
		if err != nil {
			return "", err
		}
		decrypted[i] = k
	}
	return string(decrypted), nil
}

type EncryptionMap map[rune]rune

func (m EncryptionMap) getKey(value rune) (rune, error) {
	for k, v := range m {
		if v == value {
			return k, nil
		}
	}
	return 0, errors.New("key not found")
}
