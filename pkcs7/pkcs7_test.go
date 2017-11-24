package pkcs7

import (
	"testing"
	"encoding/base64"
)

func TestEncode(t *testing.T) {
	data := []byte("1234567890")
	data = Encode(data, 32)
	ret := base64.StdEncoding.EncodeToString(data)
	if ret != "MTIzNDU2Nzg5MBYWFhYWFhYWFhYWFhYWFhYWFhYWFhY=" {
		t.Error("pkcs7.Encode test failed")
		return
	}
}

func TestDecode(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString("MTIzNDU2Nzg5MBYWFhYWFhYWFhYWFhYWFhYWFhYWFhY=")
	if err != nil {
		t.Error(err)
		return
	}
	data = Decode(data)
	if string(data) != "1234567890" {
		t.Error("pkcs7.Decode test failed")
	}
}
