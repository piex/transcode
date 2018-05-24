package transcode

import (
	"bytes"
	"io/ioutil"
	"log"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Transcode code type
type Transcode struct {
	b []byte
}

// Trans 转化字符串
type Trans interface {
	Decode(string) Trans
	Encode(string) Trans
	ToByteArray() []byte
	ToString() string
}

// Decode decode string to UTF-8
func (t *Transcode) Decode(fromcode string) Trans {
	var decoder *encoding.Decoder

	switch fromcode {
	case "GBK":
		decoder = simplifiedchinese.GBK.NewDecoder()
	case "GB18030":
		decoder = simplifiedchinese.GB18030.NewDecoder()
	case "HZGB2312":
		decoder = simplifiedchinese.HZGB2312.NewDecoder()
	default:
		decoder = simplifiedchinese.GB18030.NewDecoder()
	}

	reader := transform.NewReader(bytes.NewReader(t.b), decoder)
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Panic(err)
	}
	t.b = d
	return t
}

// Encode encode UTF-8 to GBK
func (t *Transcode) Encode(tocode string) Trans {
	var encoder *encoding.Encoder

	switch tocode {
	case "GBK":
		encoder = simplifiedchinese.GBK.NewEncoder()
	case "GB18030":
		encoder = simplifiedchinese.GB18030.NewEncoder()
	case "HZGB2312":
		encoder = simplifiedchinese.HZGB2312.NewEncoder()
	default:
		encoder = simplifiedchinese.GB18030.NewEncoder()
	}

	reader := transform.NewReader(bytes.NewReader(t.b), encoder)
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Panic(err)
	}
	t.b = d
	return t
}

// ToByteArray 输出转化后的 byte 数组
func (t *Transcode) ToByteArray() []byte {
	return t.b
}

// ToString 输出转化后的字符串
func (t *Transcode) ToString() string {
	return string(t.b)
}
