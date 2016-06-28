package controllers

import (
    "encoding/xml"
    "github.com/revel/revel"
    matrixService "matrix/service"
    wechatService "matrix/modules/wechat/service"
    "io/ioutil"
    "time"
    "encoding/base64"
    "strconv"
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

    signatureComputed := wechatService.Sign(wechatService.GetToken(), timestamp, nonce)
    if signature != signatureComputed {
        return c.RenderText("")
    }

    requestMessageBody, _ := ioutil.ReadAll(c.Request.Body)
    requestCipherMsg := new(RequestCipherMessage)
    xml.Unmarshal(requestMessageBody, requestCipherMsg)

    msgSignatureComputed := wechatService.MsgSign(wechatService.GetToken(), timestamp, nonce, requestCipherMsg.Encrypt)
    if msgSignature != msgSignatureComputed {
        return c.RenderText("")
    }

    encryptedMsg := wechatService.Base64Decode([]byte(requestCipherMsg.Encrypt))
    aesKey := wechatService.Base64Decode([]byte(wechatService.GetAesKey() + "="))

    _, msgPlainXmlBytes, _, err := wechatService.AesDecryptMsg(encryptedMsg, aesKey)
    matrixService.HandleError(err)
    requestTextMessage := new(RequestTextMessage)
    xml.Unmarshal(msgPlainXmlBytes, requestTextMessage)

    //extra response logic here

    msgTimeStampStr := strconv.FormatInt(time.Now().Unix(), 10)
    responseTextMessage := new(ResponseTextMessage)
    responseTextMessage.ToUserName = requestTextMessage.FromUserName
    responseTextMessage.FromUserName = requestTextMessage.ToUserName
    responseTextMessage.CreateTime = msgTimeStampStr
    responseTextMessage.MsgType = "text"
    responseTextMessage.Content = "哈哈，你发了 " + requestTextMessage.Content + " !"
    responseMessageBytes, _ := xml.Marshal(responseTextMessage)

    responseTextMessageEncrypted := wechatService.AesEncryptMsg([]byte("12345678"), responseMessageBytes, wechatService.GetAppId(), aesKey)
    responseTextMessageBase64 := base64.StdEncoding.EncodeToString(responseTextMessageEncrypted)
    responseMsgSignature := wechatService.MsgSign(wechatService.GetToken(), msgTimeStampStr, nonce, responseTextMessageBase64)

    responseCliperMessage := new(ResponseCliperMessage)
    responseCliperMessage.Encrypt = responseTextMessageBase64
    responseCliperMessage.MsgSignature = responseMsgSignature
    responseCliperMessage.TimeStamp = msgTimeStampStr
    responseCliperMessage.Nonce = nonce
    responseMsgXMLBytes, _ := xml.Marshal(responseCliperMessage)

    return  c.RenderText(string(responseMsgXMLBytes))
}

/*
<xml>
    <ToUserName><![CDATA[gh_c03e85cdf6af]]></ToUserName>
    <Encrypt><![CDATA[Co42g9......3Q2aOjM9Z0IBXKBw=]]></Encrypt>
</xml>
*/
type RequestCipherMessage struct {
    ToUserName string    `xml:"ToUserName"`
    Encrypt    string    `xml:"Encrypt"`
}

/*
<xml>
    <ToUserName><![CDATA[gh_c03e85cdf6af]]></ToUserName>
    <FromUserName><![CDATA[oZhq_wr-7ZA71qZshMo_Zbj_1IJQ]]></FromUserName>
    <CreateTime>1467115784</CreateTime>
    <MsgType><![CDATA[text]]></MsgType>
    <Content><![CDATA[test]]></Content>
    <MsgId>6301214312140396469</MsgId>
</xml>
*/
type RequestTextMessage struct {
    ToUserName   string    `xml:"ToUserName"`
    FromUserName string    `xml:"FromUserName"`
    CreateTime   string    `xml:"CreateTime"`
    MsgType      string    `xml:"MsgType"`
    Content      string    `xml:"Content"`
    MsgId        string    `xml:"MsgId"`
}


/*
<xml>
    <Encrypt>...</Encrypt>
    <MsgSignature>...</MsgSignature>
    <TimeStamp>...</TimeStamp>
    <Nonce>...</Nonce>
</xml>
*/

type ResponseCliperMessage struct {
    XMLName      struct{}  `xml:"xml"`

    Encrypt      string    `xml:"Encrypt"`
    MsgSignature string    `xml:"MsgSignature"`
    TimeStamp    string    `xml:"TimeStamp"`
    Nonce        string    `xml:"Nonce"`
}

/*
<xml>
    <ToUserName><![CDATA[oZhq_wr-7ZA71qZshMo_Zbj_1IJQ]]></ToUserName>
    <FromUserName><![CDATA[gh_c03e85cdf6af]]></FromUserName>
    <CreateTime>1467115784</CreateTime>
    <MsgType><![CDATA[text]]></MsgType>
    <Content><![CDATA[test]]></Content>
</xml>
*/
type ResponseTextMessage struct {
    XMLName      struct{}  `xml:"xml"`

    ToUserName   string    `xml:"ToUserName"`
    FromUserName string    `xml:"FromUserName"`
    CreateTime   string    `xml:"CreateTime"`
    MsgType      string    `xml:"MsgType"`
    Content      string    `xml:"Content"`
}