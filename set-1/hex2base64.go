package crypto

import "unicode/utf8"

/*
Base64 — стандарт кодирования двоичных данных при помощи только 64 символов ASCII.
В формате электронной почты MIME Base64 — это схема, по которой произвольная последовательность байтов преобразуется в последовательность печатных ASCII-символов.
Стандартные 62 символа дополняют +, / и = — в качестве специального кода суффикса.
*/

const (
	baseMime string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func decodeMime(input []byte) (out []byte, err error) {
	// bytes0 6 верхних байт
	b1 := input[0] >> 2
	rr, _ := utf8.DecodeRuneInString(baseMime[b1:])
	out = append(out, byte(rr))

	// bytes0 два нижних + bytes1 четыре верхних
	b2 := ((input[0] &^ 0b11111100) << 4) | (input[1]&^0b00001111)>>4
	rr, _ = utf8.DecodeRuneInString(baseMime[b2:])
	out = append(out, byte(rr))

	if input[1] == 0 {
		var rr rune = '='
		out = append(out, byte(rr))
	} else {
		// bytes1 четрые нижних + bytes2 два верхних
		b3 := ((input[1] &^ 0b11110000) << 2) | ((input[2] &^ 0b00111111) >> 6)
		rr, _ = utf8.DecodeRuneInString(baseMime[b3:])
		out = append(out, byte(rr))

	}

	if input[2] == 0 {
		var rr rune = '='
		out = append(out, byte(rr))

	} else {
		// bytes2 шесть нижних
		b4 := input[2] &^ 0b11000000
		rr, _ = utf8.DecodeRuneInString(baseMime[b4:])
		out = append(out, byte(rr))

	}

	return out, nil
}
func Hex2Base64Mime(input []byte) (out []byte, err error) {
	pos := 0
	for {
		if len(input[pos:]) <= 0 {
			break
		}
		triple := make([]byte, 3)
		w := copy(triple, input[pos:])

		// ...преобразуются в соотв с таблицей baseMime
		dec, err := decodeMime(triple)
		if err != nil {
			return []byte{}, err
		}

		out = append(out, dec...)

		pos += w
	}
	return out, err
}
