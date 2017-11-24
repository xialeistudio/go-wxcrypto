# go-wxcrypto
微信开放平台/公众平台加解密golang开发库，目前只提供基本的消息加解密。

## 功能
+ [x] 消息加密
+ [x] 消息解密

## 下一步
+ [x] 微信GET参数验证
+ [x] 微信XML解析

## 微信加密
```go
wx, err:=wxcrypto.NewWxCrypto("EncodingAesKey","APPID")
if err != nil {
	panic(err)
}
plainData:="test"
encryptData,err:=wx.Encrypt(plainData)
if err != nil {
	panic(err)
}
fmt.Print(encryptData)
```

## 微信解密
```go
wx, err:=wxcrypto.NewWxCrypto("EncodingAesKey","APPID")
if err != nil {
	panic(err)
}
encryptData:="encryptData"
plainData,err:=wx.Decrypt(encryptData)
if err != nil {
	panic(err)
}
fmt.Print(plainData)
```