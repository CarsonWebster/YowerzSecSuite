package base64

import (
	"fmt"
	"strings"
)

// Base64 character set
var base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Encode(input string) string {
	data := []byte(input)
	output := make([]byte, 4*((len(data)+2)/3))
	j, padding := 0, 0

	for i := 0; i < len(data); i += 3 {
		chunk := uint32(data[i])<<16 | uint32(data[i+1])<<8 | uint32(data[i+2])
		output[j] = base64Chars[(chunk>>18)&0x3F]
		output[j+1] = base64Chars[(chunk>>12)&0x3F]
		output[j+2] = base64Chars[(chunk>>6)&0x3F]
		output[j+3] = base64Chars[chunk&0x3F]
		j += 4
		padding = len(data) - i
	}

	if padding == 1 {
		output[len(output)-2] = '='
		output[len(output)-1] = '='
	} else if padding == 2 {
		output[len(output)-1] = '='
	}

	return string(output)
}

func Decode(encoded string) (string, error) {
	data := []byte(encoded)
	output := []byte{}
	padding := 0

	for i := 0; i < len(data); i += 4 {
		chunk := uint32(0)

		for j := 0; j < 4; j++ {
			if i+j < len(data) {
				chunk <<= 6
				index := strings.IndexByte(base64Chars, data[i+j])
				if index == -1 {
					return "", fmt.Errorf("invalid base64 character")
				}
				chunk |= uint32(index)
			} else {
				padding++
			}
		}

		for k := 0; k < 3-padding; k++ {
			value := byte((chunk >> uint(8*(2-k))) & 0xFF)
			output = append(output, value)
		}
	}

	return string(output), nil
}
