package msgcrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	//"fmt"
)

var key = []byte("abcdefghijklmnop")

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length < 1 {
		return nil, errors.New("the cipther length less than 1")
	}
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)], nil
}

func AesEncrypt(str string) (string, error) {
	origData := []byte(str)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	encodestring := base64.StdEncoding.EncodeToString(crypted)
	return encodestring, nil
}

func AesDecrypt(crypted string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "", err
	}
	//fmt.Println(decodeBytes)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(decodeBytes))
	//fmt.Println("test origdata ", origData)
	blockMode.CryptBlocks(origData, decodeBytes)
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		//fmt.Println("PKCS7UnPadding error")
		return "", err
	}
	decryptstr := string(origData)
	return decryptstr, nil
}
