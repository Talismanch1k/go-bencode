package encoder

import "strconv"

func WriteInt(dst []byte, n int64) []byte {
	dst = append(dst, 'i')
	dst = strconv.AppendInt(dst, n, 10)
	dst = append(dst, 'e')

	return dst
}

func WriteUint(dst []byte, n uint64) []byte {
	dst = append(dst, 'i')
	dst = strconv.AppendUint(dst, n, 10)
	dst = append(dst, 'e')

	return dst
}
