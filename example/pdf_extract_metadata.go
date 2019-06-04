/*

列出pdf信息:
输出:
===================
Input file: 1_2.pdf
PDF Encrypted: false
PDF Num Pages: 2
===================

 * Prints PDF metadata, tries to decrypt encrypted documents with the given password,
 * if that fails it tries an empty password as best effort.
 *
 * Run as: go run pdf_extract_metadata.go 1_2.pdf
*/

package main

import (
	"fmt"
	"os"

	unipdf "github.com/zouhuigang/pdf"
)

type PdfMetadata struct {
	Encrypted bool
	NumPages  int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Requires at least 1 argument: password input.pdf\n")
		fmt.Printf("Usage: To print information about input.pdf run: go run pdf_extract_metadata.go input.pdf\n")
		os.Exit(1)
	}

	inputPath := os.Args[1]

	ret, err := extractPdfMetadata(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input file: %s\n", inputPath)
	fmt.Printf("PDF Encrypted: %t\n", ret.Encrypted)
	fmt.Printf("PDF Num Pages: %d\n", ret.NumPages)
}

func extractPdfMetadata(inputPath string) (*PdfMetadata, error) {
	ret := PdfMetadata{}

	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	pdfReader, err := unipdf.NewPdfReader(f)
	if err != nil {
		return nil, err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return nil, err
	}

	// Try decrypting with an empty one.
	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			// Encrypted and we cannot do anything about it.
			ret.Encrypted = true
			return &ret, nil
		}
	}

	ret.Encrypted = isEncrypted

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, err
	}

	ret.NumPages = numPages

	return &ret, nil
}
