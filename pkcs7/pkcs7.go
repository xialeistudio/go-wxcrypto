package pkcs7

import "bytes"

func Encode(buf []byte, blockSize int) []byte {
	padSize := blockSize - len(buf)%blockSize
	padBuf := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(buf, padBuf...)
}

func Decode(buf []byte) []byte {
	size := len(buf)
	if size == 0 {
		return nil
	}
	padding := int(buf[size-1])
	if padding < 1 || padding > 32 {
		padding = 0
	}
	return buf[:size-padding]
}
