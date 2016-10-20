package controllers

import (
    "encoding/xml"
    "github.com/revel/revel"
    "matrix/core"
    weixin_service  "matrix/modules/weixin/service"
    "gopkg.in/xmlpath"
    "io/ioutil"
    "time"
    "encoding/base64"
    "strconv"
    "bytes"
    "strings"
)

type WechatCallback struct {
    *revel.Controller
}

func (c WechatCallback) Sign() revel.Result {
    session := c.DbSession

    var signature, timestamp, nonce, echostr string

    c.Params.Bind(&signature, "signature")
    c.Params.Bind(&timestamp, "timestamp")
    c.Params.Bind(&nonce, "nonce")
    c.Params.Bind(&echostr, "echostr")

    if echostr != "" {
        signatureComputed := weixin_service.Sign(weixin_service.GetToken(session), timestamp, nonce)
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

    signatureComputed := weixin_service.Sign(weixin_service.GetToken(session), timestamp, nonce)
    if signature != signatureComputed {
        return c.RenderText("")
    }

    requestMessageBody, _ := ioutil.ReadAll(c.Request.Body)
    requestCipherMsg := new(weixin_service.RequestCipherMessage)
    xml.Unmarshal(requestMessageBody, requestCipherMsg)

    msgSignatureComputed := weixin_service.MsgSign(weixin_service.GetToken(session), timestamp, nonce, requestCipherMsg.Encrypt)
    if msgSignature != msgSignatureComputed {
        return c.RenderText("")
    }

    encryptedMsg := weixin_service.Base64Decode([]byte(requestCipherMsg.Encrypt))
    aesKey := weixin_service.Base64Decode([]byte(weixin_service.GetEncodingAesKey(session) + "="))

    _, msgPlainXmlBytes, _, err := weixin_service.AesDecryptMsg(encryptedMsg, aesKey)
    core.HandleError(err)

    revel.TRACE.Println("request plain xml: ")
    revel.TRACE.Println(string(msgPlainXmlBytes))

    replyText := ""

    requestMsgNode, err := xmlpath.Parse(bytes.NewBuffer(msgPlainXmlBytes))
    requestMsgPath := xmlpath.MustCompile("/xml/MsgType")
    msgType, _ := requestMsgPath.String(requestMsgNode)

    requestMsgFrom := ""
    requestMsgTo := ""
    if strings.ToLower(msgType) == "text" {
        //解析为文本消息
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
        requestTextMessage := new(weixin_service.RequestTextMessage)
        xml.Unmarshal(msgPlainXmlBytes, requestTextMessage)

        requestMsgFrom = requestTextMessage.FromUserName
        requestMsgTo = requestTextMessage.ToUserName

        replyText = "哈哈，你发了消息 " + requestTextMessage.Content + " !"

    } else if strings.ToLower(msgType) == "event" {
        // Click 事件类型 Event
        /*
        <xml>
            <ToUserName><![CDATA[gh_c03e85cdf6af]]></ToUserName>
            <FromUserName><![CDATA[oZhq_wr-7ZA71qZshMo_Zbj_1IJQ]]></FromUserName>
            <CreateTime>1468128180</CreateTime>
            <MsgType><![CDATA[event]]></MsgType>
            <Event><![CDATA[CLICK]]></Event>
            <EventKey><![CDATA[V1001_TODAY_MUSIC]]></EventKey>
        </xml>
        */

        requestMsgPath = xmlpath.MustCompile("/xml/Event")
        event, _ := requestMsgPath.String(requestMsgNode)
        if strings.ToUpper(event) == "CLICK" {
            requestMsgPath = xmlpath.MustCompile("/xml/FromUserName")
            requestMsgFrom, _ = requestMsgPath.String(requestMsgNode)

            requestMsgPath = xmlpath.MustCompile("/xml/ToUserName")
            requestMsgTo, _ = requestMsgPath.String(requestMsgNode)

            requestMsgPath = xmlpath.MustCompile("/xml/EventKey")
            eventKey, _ := requestMsgPath.String(requestMsgNode)
            replyText = "哈哈，你点击了 " + eventKey + " !"
        } else {
            return c.RenderText("")
        }

    } else {
        return c.RenderText("")
    }

    //extra response logic here

    //发送文本消息
    /*
    <xml>
        <ToUserName><![CDATA[oZhq_wr-7ZA71qZshMo_Zbj_1IJQ]]></ToUserName>
        <FromUserName><![CDATA[gh_c03e85cdf6af]]></FromUserName>
        <CreateTime>1467115784</CreateTime>
        <MsgType><![CDATA[text]]></MsgType>
        <Content><![CDATA[test]]></Content>
    </xml>
    */
    msgTimeStampStr := strconv.FormatInt(time.Now().Unix(), 10)
    responseTextMessage := new(weixin_service.ResponseTextMessage)
    responseTextMessage.ToUserName = requestMsgFrom
    responseTextMessage.FromUserName = requestMsgTo
    responseTextMessage.CreateTime = msgTimeStampStr
    responseTextMessage.MsgType = "text"
    responseTextMessage.Content = replyText
    responseMessageBytes, _ := xml.Marshal(responseTextMessage)

    responseTextMessageEncrypted := weixin_service.AesEncryptMsg([]byte("12345678"), responseMessageBytes, weixin_service.GetAppId(session), aesKey)
    responseTextMessageBase64 := base64.StdEncoding.EncodeToString(responseTextMessageEncrypted)
    responseMsgSignature := weixin_service.MsgSign(weixin_service.GetToken(session), msgTimeStampStr, nonce, responseTextMessageBase64)

    responseCliperMessage := new(weixin_service.ResponseCliperMessage)
    responseCliperMessage.Encrypt = responseTextMessageBase64
    responseCliperMessage.MsgSignature = responseMsgSignature
    responseCliperMessage.TimeStamp = msgTimeStampStr
    responseCliperMessage.Nonce = nonce
    responseMsgXMLBytes, _ := xml.Marshal(responseCliperMessage)

    return c.RenderText(string(responseMsgXMLBytes))
}
