package controllers

import (
    "encoding/xml"
    "github.com/revel/revel"
    "matrix/core"
    "matrix/modules/weixin/service"
    "io/ioutil"
    "time"
    "encoding/base64"
    "strconv"
)

type WechatCallback struct {
    *revel.Controller
    core.BaseController
}

func (c WechatCallback) Sign() revel.Result {
    session := c.DbSession

    var signature, timestamp, nonce, echostr string

    c.Params.Bind(&signature, "signature")
    c.Params.Bind(&timestamp, "timestamp")
    c.Params.Bind(&nonce, "nonce")
    c.Params.Bind(&echostr, "echostr")

    if echostr != "" {
        signatureComputed := service.Sign(service.GetToken(session), timestamp, nonce)
        if signature == signatureComputed {
            return c.RenderText(echostr)
        }
    }
    return c.RenderText("")
}

func (c WechatCallback) Reply() revel.Result {
    session := c.DbSession

    var signature, msgSignature, encrypt_type, timestamp, nonce string

    c.Params.Bind(&encrypt_type, "encrypt_type")
    c.Params.Bind(&signature, "signature")
    c.Params.Bind(&msgSignature, "msg_signature")
    c.Params.Bind(&timestamp, "timestamp")
    c.Params.Bind(&nonce, "nonce")

    signatureComputed := service.Sign(service.GetToken(session), timestamp, nonce)
    if signature != signatureComputed {
        return c.RenderText("")
    }

    requestMessageBody, _ := ioutil.ReadAll(c.Request.Body)
    requestCipherMsg := new(service.RequestCipherMessage)
    xml.Unmarshal(requestMessageBody, requestCipherMsg)

    msgSignatureComputed := service.MsgSign(service.GetToken(session), timestamp, nonce, requestCipherMsg.Encrypt)
    if msgSignature != msgSignatureComputed {
        return c.RenderText("")
    }

    encryptedMsg := service.Base64Decode([]byte(requestCipherMsg.Encrypt))
    aesKey := service.Base64Decode([]byte(service.GetEncodingAesKey(session) + "="))

    _, msgPlainXmlBytes, _, err := service.AesDecryptMsg(encryptedMsg, aesKey)
    core.HandleError(err)
    requestTextMessage := new(service.RequestTextMessage)
    xml.Unmarshal(msgPlainXmlBytes, requestTextMessage)

    //extra response logic here

    msgTimeStampStr := strconv.FormatInt(time.Now().Unix(), 10)
    responseTextMessage := new(service.ResponseTextMessage)
    responseTextMessage.ToUserName = requestTextMessage.FromUserName
    responseTextMessage.FromUserName = requestTextMessage.ToUserName
    responseTextMessage.CreateTime = msgTimeStampStr
    responseTextMessage.MsgType = "text"
    responseTextMessage.Content = "哈哈，你发了 " + requestTextMessage.Content + " !"
    responseMessageBytes, _ := xml.Marshal(responseTextMessage)

    responseTextMessageEncrypted := service.AesEncryptMsg([]byte("12345678"), responseMessageBytes, service.GetAppId(session), aesKey)
    responseTextMessageBase64 := base64.StdEncoding.EncodeToString(responseTextMessageEncrypted)
    responseMsgSignature := service.MsgSign(service.GetToken(session), msgTimeStampStr, nonce, responseTextMessageBase64)

    responseCliperMessage := new(service.ResponseCliperMessage)
    responseCliperMessage.Encrypt = responseTextMessageBase64
    responseCliperMessage.MsgSignature = responseMsgSignature
    responseCliperMessage.TimeStamp = msgTimeStampStr
    responseCliperMessage.Nonce = nonce
    responseMsgXMLBytes, _ := xml.Marshal(responseCliperMessage)

    return  c.RenderText(string(responseMsgXMLBytes))
}
