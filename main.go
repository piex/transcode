package transcode

// FromByte 输入要转码的 byte 数组
func FromByte(b []byte) Trans {
	t := new(Transcode)
	t.b = b
	return t
}

// FromString 输入要转码的字符串
func FromString(s string) Trans {
	t := new(Transcode)
	t.b = []byte(s)
	return t
}
