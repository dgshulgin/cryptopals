package crypto

import "errors"

func FixedXOR(inp1 []byte, inp2 []byte) ([]byte, error) {

	if len(inp1) == 0 || len(inp2) == 0 {
		return []byte{}, errors.New("input buffer is empty")
	}

	var ret []byte
	for i := 0; i < len(inp1); i++ {
		ret = append(ret, inp1[i]^inp2[i])
	}
	return ret, nil
}
