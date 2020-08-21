package webhook

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// decrypter 是用于解密订阅事件，详见 https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM
type decrypter struct {
	block cipher.Block
}

// newDecrypter 新建一个 decrypter, key 是应用的 Encrypt Key
func newDecrypter(key string) *decrypter {
	sum := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(sum[:]) // AES-256
	if err != nil {
		// sha256 返回 32 位，这里不应该报错
		panic(err)
	}
	return &decrypter{block}
}

func (d *decrypter) Decrypt(b64CipherText string) ([]byte, error) {
	cipherText, err := base64.StdEncoding.DecodeString(b64CipherText)
	if err != nil {
		return nil, err
	}

	blockSize := d.block.BlockSize()

	// 第一个 block 是 iv
	if len(cipherText) < blockSize {
		return nil, fmt.Errorf("Invalid iv")
	}
	iv := cipherText[:blockSize]

	// 剩下的是实际数据 blocks
	cipherText = cipherText[blockSize:]
	l := len(cipherText)
	if l <= 0 {
		return nil, fmt.Errorf("No cipher text")
	}
	if l%blockSize != 0 {
		return nil, fmt.Errorf("Length of cipher text must be multiply of %d, but got %d", blockSize, l)
	}

	plainText := make([]byte, len(cipherText))
	cipher.NewCBCDecrypter(d.block, iv).CryptBlocks(plainText, cipherText)

	// unpad
	pad := int(plainText[len(plainText)-1])
	plainText = plainText[:l-pad]

	return plainText, nil
}
