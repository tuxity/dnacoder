package dnacoder

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x44, 0x4e, 0x41, 0x21} // "Hello, DNA!"
	expectedDNA := "GAGAGCACACGTATGATACATCTCTCTACGTCAGCTATATAGATCTCTGATCTGA"

	encodedDNA := encode(data)
    if encodedDNA != expectedDNA {
		t.Errorf("Encode(%s) = %s; want %s", data, encodedDNA, expectedDNA)
	}
}

func TestDecode(t *testing.T) {
	dnaSeq := "GAGAGCACACGTATGATACATCTCTCTACGTCAGCTATATAGATCTCTGATCTGA"
	expectedData := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x44, 0x4e, 0x41, 0x21} // "Hello, DNA!"

	decodedData, err := decode(dnaSeq)
	if err != nil {
		t.Fatalf("Error decoding DNA sequence: %v", err)
	}

    if !bytes.Equal(decodedData, expectedData) {
		t.Errorf("Decode(%s) = %s; want %s", dnaSeq, decodedData, expectedData)
	}
}

func TestEncodeAndDecode(t *testing.T) {
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x44, 0x4e, 0x41, 0x21} // "Hello, DNA!"

	// Encode the string to DNA
	encodedDNA := encode(data)

	// Decode back to the original string
	decodedData, err := decode(encodedDNA)
	if err != nil {
		t.Fatalf("Error decoding DNA sequence: %v", err)
	}

    if !bytes.Equal(decodedData, data) {
		t.Errorf("Encode and Decode mismatch: got %s, want %s", decodedData, data)
	}
}
