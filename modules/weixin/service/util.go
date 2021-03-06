package service

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"matrix/core"
	"sort"
)

// 微信公众号 url 签名.
func Sign(token, timestamp, nonce string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce))
	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

// 微信公众号/企业号 消息体签名.
func MsgSign(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strs.Sort()

	h := sha1.New()

	bufw := bufio.NewWriterSize(h, 128) // sha1.BlockSize 的整数倍
	bufw.WriteString(strs[0])
	bufw.WriteString(strs[1])
	bufw.WriteString(strs[2])
	bufw.WriteString(strs[3])
	bufw.Flush()

	hashsum := h.Sum(nil)
	return hex.EncodeToString(hashsum)
}

func Base64Decode(base64Str []byte) []byte {
	decodedStr := make([]byte, base64.StdEncoding.DecodedLen(len(base64Str)))
	encryptedMsgLen, err := base64.StdEncoding.Decode(decodedStr, base64Str)
	core.HandleError(err)
	decodedStr = decodedStr[:encryptedMsgLen]
	return decodedStr
}

// ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
func AesEncryptMsg(random, rawXMLMsg []byte, appId string, aesKey []byte) (ciphertext []byte) {
	const (
		BLOCK_SIZE = 32             // PKCS#7
		BLOCK_MASK = BLOCK_SIZE - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	appIdOffset := 20 + len(rawXMLMsg)
	contentLen := appIdOffset + len(appId)
	amountToPad := BLOCK_SIZE - contentLen&BLOCK_MASK
	plaintextLen := contentLen + amountToPad

	plaintext := make([]byte, plaintextLen)

	// 拼接
	copy(plaintext[:16], random)
	encodeNetworkByteOrder(plaintext[16:20], uint32(len(rawXMLMsg)))
	copy(plaintext[20:], rawXMLMsg)
	copy(plaintext[appIdOffset:], appId)

	// PKCS#7 补位
	for i := contentLen; i < plaintextLen; i++ {
		plaintext[i] = byte(amountToPad)
	}

	// 加密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, aesKey[:16])
	mode.CryptBlocks(plaintext, plaintext)

	ciphertext = plaintext
	return
}

// ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
func AesDecryptMsg(ciphertext []byte, aesKey []byte) (random, rawXMLMsg, appId []byte, err error) {
	const (
		BLOCK_SIZE = 32             // PKCS#7
		BLOCK_MASK = BLOCK_SIZE - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(ciphertext) < BLOCK_SIZE {
		err = fmt.Errorf("the length of ciphertext too short: %d", len(ciphertext))
		return
	}
	if len(ciphertext)&BLOCK_MASK != 0 {
		err = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(ciphertext))
		return
	}

	plaintext := make([]byte, len(ciphertext)) // len(plaintext) >= BLOCK_SIZE

	// 解密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, aesKey[:16])
	mode.CryptBlocks(plaintext, ciphertext)

	// PKCS#7 去除补位
	amountToPad := int(plaintext[len(plaintext)-1])
	if amountToPad < 1 || amountToPad > BLOCK_SIZE {
		err = fmt.Errorf("the amount to pad is incorrect: %d", amountToPad)
		return
	}
	plaintext = plaintext[:len(plaintext)-amountToPad]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(plaintext) <= 20 {
		err = fmt.Errorf("plaintext too short, the length is %d", len(plaintext))
		return
	}
	rawXMLMsgLen := int(decodeNetworkByteOrder(plaintext[16:20]))
	if rawXMLMsgLen < 0 {
		err = fmt.Errorf("incorrect msg length: %d", rawXMLMsgLen)
		return
	}
	appIdOffset := 20 + rawXMLMsgLen
	if len(plaintext) <= appIdOffset {
		err = fmt.Errorf("msg length too large: %d", rawXMLMsgLen)
		return
	}

	random = plaintext[:16:20]
	rawXMLMsg = plaintext[20:appIdOffset:appIdOffset]
	appId = plaintext[appIdOffset:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func encodeNetworkByteOrder(b []byte, n uint32) {
	b[0] = byte(n >> 24)
	b[1] = byte(n >> 16)
	b[2] = byte(n >> 8)
	b[3] = byte(n)
}

// 从 4 字节的网络字节序里解析出整数
func decodeNetworkByteOrder(b []byte) (n uint32) {
	return uint32(b[0])<<24 |
		uint32(b[1])<<16 |
		uint32(b[2])<<8 |
		uint32(b[3])
}
