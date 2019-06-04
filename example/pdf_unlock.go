/*
 * Unlocks PDF files, tries to decrypt encrypted documents with the given password,
 * if that fails it tries an empty password as best effort.
 *
 * Run as: go run pdf_unlock.go password output.pdf input.pdf
 * To unlock input.pdf with password 'test' and save as output.pdf run: go run pdf_unlock.go test output.pdf input.pdf
 */

package main

import (
	"fmt"
	"os"

	unipdf "github.com/zouhuigang/pdf"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Requires at least 3 arguments: password output.pdf input.pdf\n")
		fmt.Printf("Usage: To unlock input.pdf with password 'test' and save as output.pdf run: go run pdf_unlock.go test output.pdf input.pdf\n")
		os.Exit(1)
	}

	password := os.Args[1]

	outputPath := os.Args[2]
	inputPath := os.Args[3]

	err := unlockPdf(inputPath, outputPath, password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Complete, see output file: %s\n", outputPath)
}

func unlockPdf(inputPath string, outputPath string, password string) error {
	pdfWriter := unipdf.NewPdfWriter()

	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := unipdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return err
	}

	// Try decrypting both with given password and an empty one if that fails.
	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(password))
		if err != nil {
			// Fails, try fallback with empty password.
			_, err = pdfReader.Decrypt([]byte(""))
			if err != nil {
				return err
			}
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		err = pdfWriter.AddPage(page)
		if err != nil {
			return err
		}
	}

	fWrite, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}
