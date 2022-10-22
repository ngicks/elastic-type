package estype

import (
	"encoding/base64"
	"encoding/json"
)

// Binary is elasticsearch type that represents binary data as base64 encoded string.
//
// Binary keeps it lazily encode/decoded.
type Binary struct {
	encoded string
	decoded []byte
}

// BytesToBinary makes Binary from bytes.
func BytesToBinary(data []byte) Binary {
	return Binary{
		decoded: data,
	}
}

// StringToBinary makes Binary from base64-encoded string.
// If storeDecoded is true, decoded bytes also stored in returned Binary.
func StringToBinary(data string, storeDecoded bool) (Binary, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return Binary{}, err
	}

	if storeDecoded {
		return Binary{encoded: data, decoded: decoded}, nil
	}
	return Binary{encoded: data}, nil
}

// StringToBinaryUnchecked makes Binary from base64-encoded string.
// It does no check for input string validity, whether it is base64 encoded or not.
func StringToBinaryUnchecked(data string) Binary {
	return Binary{encoded: data}
}

// MarshalJSON marshals this type into byte slice representing JSON boolean literal, true or false.
func (b Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Binary) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &b.encoded)
	if err != nil {
		return err
	}

	_, err = base64.StdEncoding.DecodeString(b.encoded)
	if err != nil {
		return err
	}

	return nil
}

// String returns base64 encoded string
func (b Binary) String() string {
	if b.encoded == "" {
		b.encoded = base64.StdEncoding.EncodeToString(b.decoded)
	}
	return b.encoded
}

// Bytes returns decoded []byte.
func (b Binary) Bytes() []byte {
	if b.decoded == nil {
		// Validity is checked in constructor function or Unmarshal method.
		b.decoded, _ = base64.StdEncoding.DecodeString(b.encoded)
	}
	return b.decoded
}
