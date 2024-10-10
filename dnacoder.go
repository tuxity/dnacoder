package dnacoder

import (
	"fmt"
	"math/big"
	"strings"
)

// Lookup table based on the previous nucleotide to prevent homopolymers
var lookupTable = map[byte][3]byte{
	'A': {'C', 'G', 'T'},
	'C': {'G', 'T', 'A'},
	'G': {'T', 'A', 'C'},
	'T': {'A', 'C', 'G'},
}
// since it's a fixed table, decode more efficiently and avoid looping
var reverseLookupTable = map[byte]map[byte]byte{
	'A': {'C': 0, 'G': 1, 'T': 2},
	'C': {'G': 0, 'T': 1, 'A': 2},
	'G': {'T': 0, 'A': 1, 'C': 2},
	'T': {'A': 0, 'C': 1, 'G': 2},
}

func byteArrayToTernary(data []byte) string {
	bigInt := new(big.Int)
	bigInt.SetBytes(data)
	ternaryStr := bigInt.Text(3)
	return ternaryStr
}

func ternaryToByteArray(ternaryStr string) []byte {
	bigInt := new(big.Int)
	bigInt.SetString(ternaryStr, 3)
	return bigInt.Bytes()
}

// Encodes a ternary string (sequence of trits) to a DNA sequence (sequence of nucleotides)
func encodeTernaryToDNA(ternaryStr string) string {
	var dnaSeq strings.Builder
	previousNucleotide := byte('A')

	for _, tritChar := range ternaryStr {
        trit := tritChar - '0' // Convert to integer
		nextNucleotide := lookupTable[previousNucleotide][trit]
		dnaSeq.WriteByte(nextNucleotide)
		previousNucleotide = nextNucleotide
	}

	return dnaSeq.String()
}

// Decodes a DNA sequence (nucleotides) back to a ternary string (sequence of trits)
func decodeDNAToTernary(dnaSeq string) (string, error) {
	var ternaryStr strings.Builder
	if len(dnaSeq) == 0 {
		return "", fmt.Errorf("empty DNA sequence")
	}

	previousNucleotide := byte('A')
	for i := 0; i < len(dnaSeq); i++ {
		nextNucleotide := dnaSeq[i]

		trit := reverseLookupTable[previousNucleotide][nextNucleotide]
		ternaryStr.WriteByte(byte(trit) + '0') // Convert to ASCII character
		previousNucleotide = nextNucleotide
	}

	return ternaryStr.String(), nil
}

func encode(data string) string {
	ternaryStr := byteArrayToTernary([]byte(data))
	dnaSequence := encodeTernaryToDNA(ternaryStr)

	return dnaSequence
}

func decode(dnaSeq string) (string, error) {
	ternaryStr, err := decodeDNAToTernary(dnaSeq)
	if err != nil {
		return "", err
	}

	byteArray := ternaryToByteArray(ternaryStr)

	return string(byteArray), nil
}
