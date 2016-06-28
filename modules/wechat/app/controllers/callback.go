package controllers

import (
    "encoding/xml"
    "github.com/revel/revel"
    matrixService "matrix/service"
    wechatService "matrix/modules/wechat/service"
    "fmt"
    "io/ioutil"
    "encoding/base64"
)

type WechatCallback struct {
    *revel.Controller
    matrixService.BaseController
}

func (c WechatCallback) Sign() revel.Result {
    var signature, timestamp, nonce, echostr string

    c.Params.Bind(&signature, "signature")
    c.Params.Bind(&timestamp, "timestamp")
    c.Params.Bind(&nonce, "nonce")
    c.Params.Bind(&echostr, "echostr")

    if echostr != "" {
        signatureComputed := wechatService.Sign(wechatService.GetToken(), timestamp, nonce)
        if signature == signatureComputed {
            return c.RenderText(echostr)
        }
    }
    return c.RenderText("")
}

func (c WechatCallback) Reply() revel.Result {
    var signature, msgSignature, encrypt_type, timestamp, nonce string

    c.Params.Bind(&encrypt_type, "encrypt_type")
    c.Params.Bind(&signature, "signature")
    c.Params.Bind(&msgSignature, "msg_signature")
    c.Params.Bind(&timestamp, "timestamp")
    c.Params.Bind(&nonce, "nonce")

    fmt.Println(encrypt_type)
    fmt.Println(signature)
    fmt.Println(msgSignature)
    fmt.Println(timestamp)
    fmt.Println(nonce)

    signatureComputed := wechatService.Sign(wechatService.GetToken(), timestamp, nonce)
    if signature != signatureComputed {
        fmt.Println("signature != signatureComputed")
        return c.RenderText("")
    }

    /*
    bodyBuffer := new(bytes.Buffer)
    bodyBuffer.ReadFrom(c.Request.Body)
    msgXml := bodyBuffer.String()
    fmt.Println(msgXml)
    */

    msgRequestBody, _ := ioutil.ReadAll(c.Request.Body)
    msg := new(Message)
    xml.Unmarshal(msgRequestBody, msg)

    msgSignatureComputed := wechatService.MsgSign(wechatService.GetToken(), timestamp, nonce, msg.Encrypt)
    if msgSignature != msgSignatureComputed {
        return c.RenderText("")
    }

    fmt.Println(msg.ToUserName)
    fmt.Println(msg.Encrypt)

    //msg.Encrypt base64 decode
    msgEncryptByte := []byte(msg.Encrypt)
    encryptedMsg := make([]byte, base64.StdEncoding.DecodedLen(len(msgEncryptByte)))
    encryptedMsgLen, _ := base64.StdEncoding.Decode(encryptedMsg, msgEncryptByte)
    encryptedMsg = encryptedMsg[:encryptedMsgLen]

    fmt.Println(encryptedMsg)

    //aeskey base64 decode
    encodedAesKey := []byte(wechatService.GetAesKey() + "=")
    aesKey := make([]byte, base64.StdEncoding.DecodedLen(len(encodedAesKey)))
    aesKeyLen, _ := base64.StdEncoding.Decode(aesKey, encodedAesKey)
    aesKey = aesKey[: aesKeyLen]

    random, msgPlaintext, haveAppIdBytes, _ := wechatService.AesDecryptMsg(encryptedMsg, aesKey)
    fmt.Println(random)
    fmt.Println(string(msgPlaintext))
    fmt.Println(haveAppIdBytes)

    return c.RenderText("")
}

/*
<xml>
    <ToUserName>...</ToUserName>
    <Encrypt>...</Encrypt>
</xml>
*/
type Message struct {
    ToUserName string  `xml:"ToUserName"`
    Encrypt    string  `xml:"Encrypt"`
}


/*
package main

import "fmt"
import "crypto/aes"

func main() {
    bc, err := aes.NewCipher([]byte("key3456789012345"))
    if (err != nil) {
        fmt.Println(err);
    }
    fmt.Printf("The block size is %d\n", bc.BlockSize())

    var dst = make([]byte, 16)
    var src = []byte("sensitive1234567")

    bc.Encrypt(dst, src)
    fmt.Println(dst)
}
 */