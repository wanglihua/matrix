package core

import (
	"testing"
)

func Test_convert_num_to_cny(t *testing.T) {
	if Convert_num_to_cny(12345.78) != "壹万贰仟叁佰肆拾伍元柒角捌分" {
		t.Error("转换12345.78至中文大写出错！")
	}
}

func Test_convert_num_to_cap(t *testing.T) {
	if Convert_num_to_cap(123450000.78) != "1亿2345万.78" {
		t.Error("转换123450000.78为中文分割失败")
	}
}
