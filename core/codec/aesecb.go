package codec

import (
	"crypto/cipher"
	"errors"
)

var ErrPaddingSize = errors.New("padding size error")

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb { _ = "STUB: not implemented"; return nil }

type ecbEncrypter ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	_ = "STUB: not implemented"
	return *new(cipher.BlockMode)
}

func (x *ecbEncrypter) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) { _ = "STUB: not implemented"; return }

type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	_ = "STUB: not implemented"
	return *new(cipher.BlockMode)
}

func (x *ecbDecrypter) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) { _ = "STUB: not implemented"; return }

func EcbDecrypt(key, src []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func EcbDecryptBase64(key, src string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func EcbEncrypt(key, src []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func EcbEncryptBase64(key, src string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getKeyBytes(key string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func pkcs5Padding(ciphertext []byte, blockSize int) []byte { _ = "STUB: not implemented"; return nil }

func pkcs5Unpadding(src []byte, blockSize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
