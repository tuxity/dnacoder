# dnacoder

A Go module for encoding and decoding data using a ternary system and DNA nucleotide mapping. The module converts strings into DNA sequences and decodes DNA sequences back into strings, ensuring no homopolymers in the DNA sequence.

## Installation

To install the package, run the following command:

```bash
go get github.com/tuxity/dnacoder
```

## Usage

### Encoding a string to a DNA sequence

You can encode any string into a DNA sequence using `encode`.

```go
package main

import (
    "io/ioutil"
    "log"
    "fmt"
    "github.com/tuxity/dnacoder"
)

func main() {
    fileData, err := ioutil.ReadFile("example.txt") // Read the file as binary data
    if err != nil {
        log.Fatal(err)
    }
    dnaSequence := dnacoder.encode(fileData)
    fmt.Println("Encoded DNA Sequence:", dnaSequence)
}
```

### Decoding a DNA sequence back to a string

You can decode a DNA sequence back into the original string using `decode`.

```go
package main

import (
    "io/ioutil"
    "log"
    "github.com/tuxity/dnacoder"
)

func main() {
    dnaSequence := "GAGAGCACACGTATGATACATCTCTCTACGTCAGCTATATAGATCTCTGATCTGA"
    decodedData, err := dnacoder.decode(dnaSequence)
    if err != nil {
        log.Fatal(err)
    }
    err = ioutil.WriteFile("decoded_example.txt", decodedData, 0644)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Example Output

For the input string `"Hello, DNA!"`, the output might look like:

```
Encoded DNA Sequence: GAGAGCACACGTATGATACATCTCTCTACGTCAGCTATATAGATCTCTGATCTGA
Decoded String: Hello, DNA!
```

## Running Tests

This module includes unit tests. You can run the tests using the Go testing tool:

```bash
go test
```

If all tests pass, you should see output like:

```
ok  	github.com/tuxity/dnaencoder	0.123s
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
