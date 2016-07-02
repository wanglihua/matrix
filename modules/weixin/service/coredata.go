package service


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

type WeixinError struct {
    ErrCode int  `json:"errcode"`
    ErrMsg  string `json:"errmsg"`
}

type AccessToken struct {
    Token     string `json:"access_token"`
    ExpiresIn int    `json:"expires_in"`
}

