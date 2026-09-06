package codec

import (
	"crypto/rsa"
	"errors"
)

var (
	ErrPrivateKey = errors.New("private key error")

	ErrPublicKey = errors.New("failed to parse PEM block containing the public key")

	ErrNotRsaKey = errors.New("key type is not RSA")
)

type (
	RsaDecrypter interface {
		Decrypt(input []byte) ([]byte, error)
		DecryptBase64(input string) ([]byte, error)
	}

	RsaEncrypter interface {
		Encrypt(input []byte) ([]byte, error)
	}

	rsaBase struct {
		bytesLimit int
	}

	rsaDecrypter struct {
		rsaBase
		privateKey *rsa.PrivateKey
	}

	rsaEncrypter struct {
		rsaBase
		publicKey *rsa.PublicKey
	}
)

func NewRsaDecrypter(file string) (RsaDecrypter, error) {
	_ = "STUB: not implemented"
	return *new(RsaDecrypter), nil
}

func (r *rsaDecrypter) Decrypt(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *rsaDecrypter) DecryptBase64(input string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRsaEncrypter(key []byte) (RsaEncrypter, error) {
	_ = "STUB: not implemented"
	return *new(RsaEncrypter), nil
}

func (r *rsaEncrypter) Encrypt(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *rsaBase) crypt(input []byte, cryptFn func([]byte) ([]byte, error)) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rsaDecryptBlock(privateKey *rsa.PrivateKey, block []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rsaEncryptBlock(publicKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRsaOAEPDecrypter(file string) (RsaDecrypter, error) {
	_ = "STUB: not implemented"
	return *new(RsaDecrypter), nil
}

func NewRsaOAEPEncrypter(key []byte) (RsaEncrypter, error) {
	_ = "STUB: not implemented"
	return *new(RsaEncrypter), nil
}

type rsaOAEPDecrypter struct {
	rsaBase
	privateKey *rsa.PrivateKey
}

func (r *rsaOAEPDecrypter) Decrypt(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *rsaOAEPDecrypter) DecryptBase64(input string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type rsaOAEPEncrypter struct {
	rsaBase
	publicKey *rsa.PublicKey
}

func (r *rsaOAEPEncrypter) Encrypt(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
