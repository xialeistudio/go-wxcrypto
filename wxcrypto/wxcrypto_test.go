package wxcrypto

import (
	"testing"
)

var wx, err = NewWxCrypto("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG", "wxb11529c136998cb6")
var plainData = "<xml><ToUserName><![CDATA[oia2Tj我是中文jewbmiOUlr6X-1crbLOvLw]]></ToUserName><FromUserName><![CDATA[gh_7f083739789a]]></FromUserName><CreateTime>1407743423</CreateTime><MsgType><![CDATA[video]]></MsgType><Video><MediaId><![CDATA[eYJ1MbwPRJtOvIEabaxHs7TX2D-HV71s79GUxqdUkjm6Gs2Ed1KF3ulAOA9H1xG0]]></MediaId><Title><![CDATA[testCallBackReplyVideo]]></Title><Description><![CDATA[testCallBackReplyVideo]]></Description></Video></xml>"
var encryptData = ""

func TestNewWxCrypto(t *testing.T) {
	if err != nil {
		t.Error("NewWxCrypto", err)
		return
	}
}

func TestWxCrypto_Encrypt(t *testing.T) {
	encryptData, err = wx.Encrypt(plainData)
	if err != nil {
		t.Error("WxCrypto.Encrypt", err)
		return
	}
}

func TestWxCrypto_Decrypt(t *testing.T) {
	decryptData, err := wx.Decrypt(encryptData)
	if err != nil {
		t.Error("WxCrypto.Decrypt", err)
		return
	}
	if decryptData != plainData {
		t.Error("WxCrypto.Decrypt invalid result")
	}
}
