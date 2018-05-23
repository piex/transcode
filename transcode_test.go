package transcode

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

const path = "./testdata"

var utf8Byte, _ = ioutil.ReadFile(filepath.Join(path, "UTF-8.txt"))
var utf8String = string(utf8Byte)
var gbkByte, _ = ioutil.ReadFile(filepath.Join(path, "GBK.txt"))
var gbkString = string(gbkByte)
var gb18030Byte, _ = ioutil.ReadFile(filepath.Join(path, "GB18030.txt"))
var gb18030String = string(gb18030Byte)
var gb2312Byte, _ = ioutil.ReadFile(filepath.Join(path, "GB2312.txt"))
var gb2312String = string(gb2312Byte)

func Test_Trans_GBK_ByteArray_TO_UTF8(t *testing.T) {
	d := new(Transcode).FromByte(gbkByte).Decode("GBK")
	s := d.ToString()
	b := d.ToByteArray()
	if s != utf8String {
		t.Error("Failed to transcode GBK ByteArray to UTF-8 String.")
	}
	if !bytes.Equal(b, utf8Byte) {
		t.Error("Failed to transcode GBK ByteArray to UTF-8 ByteArray.")
	}
}

func Test_Trans_GB18030_ByteArray_TO_UTF8(t *testing.T) {
	d := new(Transcode).FromByte(gb18030Byte).Decode("GB18030")
	s := d.ToString()
	b := d.ToByteArray()
	if s != utf8String {
		t.Error("Failed to transcode GB18030 ByteArray to UTF-8 String.")
	}
	if !bytes.Equal(b, utf8Byte) {
		t.Error("Failed to transcode GB18030 ByteArray to UTF-8 ByteArray.")
	}
}

func Test_Trans_GB2312_ByteArray_TO_UTF8(t *testing.T) {
	d := new(Transcode).FromByte(gb2312Byte).Decode("GB2312")
	s := d.ToString()
	b := d.ToByteArray()
	if s != utf8String {
		t.Error("Failed to transcode GB2312 ByteArray to UTF-8 String.")
	}
	if !bytes.Equal(b, utf8Byte) {
		t.Error("Failed to transcode GB2312 ByteArray to UTF-8 ByteArray.")
	}
}

func Test_Trans_GBK_String_TO_UTF8(t *testing.T) {
	d := new(Transcode).FromString(gbkString).Decode("GBK")
	s := d.ToString()
	b := d.ToByteArray()
	if s != utf8String {
		t.Error("Failed to transcode GBK String to UTF-8 String.")
	}
	if !bytes.Equal(b, utf8Byte) {
		t.Error("Failed to transcode GBK String to UTF-8 ByteArray.")
	}
}

func Test_Trans_GB18030_String_TO_UTF8(t *testing.T) {
	d := new(Transcode).FromString(gb18030String).Decode("GB18030")
	s := d.ToString()
	b := d.ToByteArray()
	if s != utf8String {
		t.Error("Failed to transcode GB18030 String to UTF-8 String.")
	}
	if !bytes.Equal(b, utf8Byte) {
		t.Error("Failed to transcode GB18030 String to UTF-8 ByteArray.")
	}
}

func Test_Trans_GB2312_String_TO_UTF8(t *testing.T) {
	d := new(Transcode).FromString(gb2312String).Decode("GB2312")
	s := d.ToString()
	b := d.ToByteArray()
	if s != utf8String {
		t.Error("Failed to transcode GB2312 String to UTF-8 String.")
	}
	if !bytes.Equal(b, utf8Byte) {
		t.Error("Failed to transcode GB2312 String to UTF-8 ByteArray.")
	}
}

func Test_Trans_UTF8_ByteArray_TO_GBK(t *testing.T) {
	gbk := new(Transcode).FromByte(utf8Byte).Encode("GBK")
	s := gbk.ToString()
	b := gbk.ToByteArray()
	if s != gbkString {
		t.Error("Failed to transcode UTF-8 ByteArray to GBK String.")
	}
	if !bytes.Equal(b, gbkByte) {
		t.Error("Failed to transcode UTF-8 ByteArray to GBK ByteArray.")
	}
}

func Test_Trans_UTF8_ByteArray_TO_GB18030(t *testing.T) {
	gb18030 := new(Transcode).FromByte(utf8Byte).Encode("GB18030")
	s := gb18030.ToString()
	b := gb18030.ToByteArray()
	if s != gbkString {
		t.Error("Failed to transcode UTF-8 ByteArray to GB18030 String.")
	}
	if !bytes.Equal(b, gb18030Byte) {
		t.Error("Failed to transcode UTF-8 ByteArray to GB18030 ByteArray.")
	}
}

func Test_Trans_UTF8_ByteArray_TO_GB2312(t *testing.T) {
	gb2312 := new(Transcode).FromByte(utf8Byte).Encode("GB2312")
	s := gb2312.ToString()
	b := gb2312.ToByteArray()
	if s != gbkString {
		t.Error("Failed to transcode UTF-8 ByteArray to GB2312 String.")
	}
	if !bytes.Equal(b, gb2312Byte) {
		t.Error("Failed to transcode UTF-8 ByteArray to GB2312 ByteArray.")
	}
}

func Test_Trans_UTF8_String_TO_GBK(t *testing.T) {
	gbk := new(Transcode).FromString(utf8String).Encode("GBK")
	s := gbk.ToString()
	b := gbk.ToByteArray()
	if s != gbkString {
		t.Error("Failed to transcode UTF-8 String to GBK String.")
	}
	if !bytes.Equal(b, gbkByte) {
		t.Error("Failed to transcode UTF-8 String to GBK ByteArray.")
	}
}

func Test_Trans_UTF8_String_TO_GB18030(t *testing.T) {
	gb18030 := new(Transcode).FromString(utf8String).Encode("GB18030")
	s := gb18030.ToString()
	b := gb18030.ToByteArray()
	if s != gbkString {
		t.Error("Failed to transcode UTF-8 String to GB18030 String.")
	}
	if !bytes.Equal(b, gb18030Byte) {
		t.Error("Failed to transcode UTF-8 String to GB18030 ByteArray.")
	}
}

func Test_Trans_UTF8_String_TO_GB2312(t *testing.T) {
	gb2312 := new(Transcode).FromString(utf8String).Encode("GB2312")
	s := gb2312.ToString()
	b := gb2312.ToByteArray()
	if s != gbkString {
		t.Error("Failed to transcode UTF-8 String to GB2312 String.")
	}
	if !bytes.Equal(b, gb2312Byte) {
		t.Error("Failed to transcode UTF-8 String to GB2312 ByteArray.")
	}
}
