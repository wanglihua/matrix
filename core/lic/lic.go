package lic

import (
	"bytes"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"crypto/des"
	"crypto/cipher"
	"log"
	"io/ioutil"
	"time"
)

var encrypt_salt = "Jiaxing Huadu Information Technolony Co.,Ltd"

func EncryptLic(lic_app_name string, date string) string {
	app_name_md5 := make_md5(lic_app_name + encrypt_salt)
	date_encrypt := des_encrypt([]byte(date), []byte(encrypt_salt))
	date_encrypt_hex_str := bytes_to_hex_string(date_encrypt)
	md5_check := make_md5(lic_app_name + date + app_name_md5 + date_encrypt_hex_str + encrypt_salt)
	return lic_app_name + date + app_name_md5 + date_encrypt_hex_str + md5_check
}

func DecryptLic(lic_app_name string, lic string) string {
	if len(lic) <= 88 {
		log.Fatal(fmt.Sprintf("[L000] lic not valid: %s", lic))
	}
	lic_runes := []rune(lic)
	lic_len := len(lic_runes)
	if make_md5(string(lic_runes[0:lic_len - 32]) + encrypt_salt) != string(lic_runes[lic_len - 32:lic_len]) {
		log.Fatal(fmt.Sprintf("[L001] lic not valid: %s", lic))
	}
	encrypt_date := lic_runes[lic_len - 48:lic_len - 32]
	decrypt_date := string(des_decrypt(hex_string_to_bytes(string(encrypt_date)), []byte(encrypt_salt)))
	if make_md5(lic_app_name + encrypt_salt) != string(lic_runes[lic_len - 80:lic_len - 48]) {
		log.Fatal(fmt.Sprintf("[L002] lic not valid: %s", lic))
	}
	if string(lic_runes[lic_len - 88:lic_len - 80]) != decrypt_date {
		log.Fatal(fmt.Sprintf("[L003] lic not valid: %s", lic))
	}
	if string(lic_runes[0:lic_len - 88]) != lic_app_name {
		log.Fatal(fmt.Sprintf("[L004] lic not valid: %s", lic))
	}
	return decrypt_date
}

func ValidAppLic(lic_app_name string, lic_file_path string) {
	lic_bytes, err := ioutil.ReadFile(lic_file_path)
	if err != nil {
		log.Fatal(fmt.Sprintf("[L005] lic file path not valid: %s", lic_file_path))
	}
	lic_date := DecryptLic(lic_app_name, string(lic_bytes))
	current_date := time.Now().Format("20060102")
	if lic_date < current_date {
		log.Fatal(fmt.Sprintf("[L006] lic not valid: %s", string(lic_bytes)))
	}
}

func make_md5(clear_text string) string {
	h := md5.New()
	h.Write([]byte(clear_text))
	cipher_str := h.Sum(nil)
	return hex.EncodeToString(cipher_str)
}

func des_encrypt(orig_data []byte, key []byte) []byte {
	key = make_right_len_des_key(key)
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(orig_data))
	blockMode.CryptBlocks(crypted, orig_data)
	return crypted
}

func des_decrypt(crypted []byte, key []byte) []byte {
	key = make_right_len_des_key(key)
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	return origData
}

func make_right_len_des_key(key_bytes []byte) []byte {
	var des_key_len = 8
	key_bytes_len := len(key_bytes)
	if key_bytes_len > des_key_len {
		key_bytes = key_bytes[0:des_key_len]
	} else if key_bytes_len != des_key_len {
		key_bytes = append(key_bytes, bytes.Repeat([]byte("0"), des_key_len - key_bytes_len)...)
	}
	return key_bytes
}

func bytes_to_hex_string(bytes_data []byte) string {
	var buffer bytes.Buffer
	bytes_data_len := len(bytes_data)
	for i := 0; i < bytes_data_len; i++ {
		hex_str := fmt.Sprintf("%x", bytes_data[i])
		if len(hex_str) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(hex_str)
	}
	return buffer.String()
}

func hex_string_to_bytes(hex_str string) []byte {
	hex_runes := []rune(hex_str)
	hex_runes_len := len(hex_runes)
	var hex_bytes = make([]byte, 0)
	for i := 0; i < hex_runes_len / 2; i++ {
		var b byte = byte(hex_digit_rune_to_num(hex_runes[2 * i]) * 16 + hex_digit_rune_to_num(hex_runes[2 * i + 1]))
		hex_bytes = append(hex_bytes, b)
	}
	return hex_bytes
}

func hex_digit_rune_to_num(hex_digit_rune rune) int {
	if hex_digit_rune >= rune('a') {
		return int(hex_digit_rune - rune('a')) + 10
	} else {
		return int(hex_digit_rune - rune('0'))
	}
}
