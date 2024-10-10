package dnacoder

import (
	"testing"
)

func TestEncode(t *testing.T) {
	data := "Hello, DNA!"
	expectedDNA := "GAGAGCACACGTATGATACATCTCTCTACGTCAGCTATATAGATCTCTGATCTGA"

	encodedDNA := encode(data)
	if encodedDNA != expectedDNA {
		t.Errorf("Encode(%s) = %s; want %s", data, encodedDNA, expectedDNA)
	}
}

func TestDecode(t *testing.T) {
	dnaSeq := "GAGAGCACACGTATGATACATCTCTCTACGTCAGCTATATAGATCTCTGATCTGA"
	expectedString := "Hello, DNA!"

	decodedString, err := decode(dnaSeq)
	if err != nil {
		t.Fatalf("Error decoding DNA sequence: %v", err)
	}

	if decodedString != expectedString {
		t.Errorf("Decode(%s) = %s; want %s", dnaSeq, decodedString, expectedString)
	}
}

func TestEncodeAndDecode(t *testing.T) {
	data := "Test String"

	// Encode the string to DNA
	encodedDNA := encode(data)

	// Decode back to the original string
	decodedString, err := decode(encodedDNA)
	if err != nil {
		t.Fatalf("Error decoding DNA sequence: %v", err)
	}

	if decodedString != data {
		t.Errorf("Encode and Decode mismatch: got %s, want %s", decodedString, data)
	}
}
