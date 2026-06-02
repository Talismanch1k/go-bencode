package encoder

func WriteBytes(dst []byte, data []byte) []byte {
	WriteInt(dst, int64(len(data)))
	dst = append(dst, ':')
	dst = append(dst, data...)

	return dst
}
