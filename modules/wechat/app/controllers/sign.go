package controllers

import (
    "github.com/revel/revel"
    "matrix/service"
    "sort"
    "crypto/sha1"
    "encoding/hex"
)

type WechatSign struct {
    *revel.Controller
    service.BaseController
}

var token = "matrix_dev_token"

func (c WechatSign) Index(signature, timestamp, nonce, echostr string) revel.Result {
    signatureComputed := sign(token, timestamp, nonce)
    if signature == signatureComputed {
        return c.RenderText(echostr)
    } else {
        return c.RenderText("")
    }
    return nil
}

// Sign 微信公众号 url 签名.
func sign(token, timestamp, nonce string) (signature string) {
    strs := sort.StringSlice{token, timestamp, nonce}
    strs.Sort()

    buf := make([]byte, 0, len(token) + len(timestamp) + len(nonce))
    buf = append(buf, strs[0]...)
    buf = append(buf, strs[1]...)
    buf = append(buf, strs[2]...)

    hashsum := sha1.Sum(buf)
    return hex.EncodeToString(hashsum[:])
}
