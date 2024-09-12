package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThomasGobbs(t *testing.T) {

	var (
		input  string = "Man"
		output string = "TWFu"
	)

	o, err := Hex2Base64([]byte(input))
	assert.NoError(t, err)
	assert.EqualValues(t, output, string(o))
}
func TestMoreThomasGobbs(t *testing.T) {

	var (
		input string = "Man is distinguished, not only by his reason, but by this singular passion from other animals, which is a lust of the mind, that by a perseverance of delight in the continued and indefatigable generation of knowledge, exceeds the short vehemence of any carnal pleasure."

		output string = "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0aGlzIHNpbmd1bGFyIHBhc3Npb24gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1c3Qgb2YgdGhlIG1pbmQsIHRoYXQgYnkgYSBwZXJzZXZlcmFuY2Ugb2YgZGVsaWdodCBpbiB0aGUgY29udGludWVkIGFuZCBpbmRlZmF0aWdhYmxlIGdlbmVyYXRpb24gb2Yga25vd2xlZGdlLCBleGNlZWRzIHRoZSBzaG9ydCB2ZWhlbWVuY2Ugb2YgYW55IGNhcm5hbCBwbGVhc3VyZS4="
	)

	o, err := Hex2Base64([]byte(input))
	assert.NoError(t, err)
	assert.EqualValues(t, output, string(o))
}

func TestChallenge1(t *testing.T) {
	var input []byte = []byte{49, 27, 0x6d, 20, 0x6b, 69, 0x6c, 0x6c, 69, 0x6e, 67, 20, 79, 0x6f, 75, 72, 20, 62, 72, 61, 69, 0x6e, 20, 0x6c, 69, 0x6b, 65, 20, 61, 20, 70, 0x6f, 69, 73, 0x6f, 0x6e, 0x6f, 75, 73, 20, 0x6d, 75, 73, 68, 72, 0x6f, 0x6f, 0x6d}
	var output string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	o, err := Hex2Base64(input)
	assert.NoError(t, err)
	assert.EqualValues(t, output, string(o))
}
