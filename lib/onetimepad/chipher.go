package onetimepad

import "errors"

func Encrypt(plane []byte, key []byte) ([]byte, error) {
	if len(plane) != len(key) {
		return nil, errors.New("key length is mismatch")
	}

	encrypted := make([]byte, len(plane))
	for i := 0; i < len(plane); i++ {
		encrypted[i] = plane[i] ^ key[i]
	}
	return encrypted, nil
}

func Decrypt(encrypted []byte, key []byte) ([]byte, error) {
	if len(encrypted) != len(key) {
		return nil, errors.New("key length is mismatch")
	}

	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i++ {
		decrypted[i] = encrypted[i] ^ key[i]
	}
	return decrypted, nil
}
