package wxcrypto

import (
	"encoding/base64"
	"errors"
	"crypto/aes"
	"wxcrypto-demo/pkcs7"
	"crypto/cipher"
	"bytes"
	"encoding/binary"
	"io"
	"crypto/rand"
)

type WxCrypto struct {
	aesKey []byte
	iv     []byte
	appid  string
}

func NewWxCrypto(aesKey, appid string) (*WxCrypto, error) {
	if len(aesKey) != 43 {
		return nil, errors.New("EncodingAesKey length invalid")
	}
	buf, err := base64.StdEncoding.DecodeString(aesKey + "=")
	if err != nil {
		return nil, err
	}
	return &WxCrypto{aesKey: buf, appid: appid, iv: buf[:16]}, nil
}

func (w *WxCrypto) Encrypt(str string) (string, error) {
	// 获取16字节随机字符串
	randBuf := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, randBuf); err != nil {
		return "", err
	}
	// 获取消息长度
	msg := []byte(str)
	msgLenBuf := new(bytes.Buffer)
	if err := binary.Write(msgLenBuf, binary.BigEndian, int32(len(msg))); err != nil {
		return "", err
	}
	// 根据微信规则组装消息体
	plainData := bytes.Join([][]byte{randBuf, msgLenBuf.Bytes(), msg, []byte(w.appid)}, nil)
	block, err := aes.NewCipher(w.aesKey)
	if err != nil {
		return "", err
	}
	plainData = pkcs7.Encode(plainData, block.BlockSize())
	cbc := cipher.NewCBCEncrypter(block, w.iv)
	cipherData := make([]byte, len(plainData))
	cbc.CryptBlocks(cipherData, plainData)
	return base64.StdEncoding.EncodeToString(cipherData), nil
}

func (w *WxCrypto) Decrypt(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("invalid encrypt data")
	}
	cipherData, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(w.aesKey)
	if err != nil {
		return "", err
	}
	cbc := cipher.NewCBCDecrypter(block, w.iv)
	if err != nil {
		return "", err
	}
	plainData := make([]byte, len(cipherData))
	cbc.CryptBlocks(plainData, cipherData)
	// 去除补位
	plainData = pkcs7.Decode(plainData)
	// 获取消息体长度
	var msgLen int32
	if err := binary.Read(bytes.NewBuffer(plainData[16:20]), binary.BigEndian, &msgLen); err != nil {
		return "", nil
	}
	// 读取消息
	msgBuf := plainData[20:20+msgLen]
	appid := plainData[20+msgLen:]
	if string(appid) != w.appid {
		return "", errors.New("invalid appid")
	}
	return string(msgBuf), nil
}
