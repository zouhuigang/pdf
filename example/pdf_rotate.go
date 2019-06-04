package main

import (
	"errors"
	"fmt"
	"os"

	pdf "github.com/zouhuigang/pdf"
)

func main() {

	inputPath := "8_16.pdf"
	outputPath := "rotate_8_16.pdf"

	var degrees int64 = 180

	if degrees%90 != 0 {
		fmt.Printf("Degrees needs to be a multiple of 90\n")
		os.Exit(1)
	}

	err := rotatePdf(inputPath, degrees, outputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Complete, see output file: %s\n", outputPath)
}

// Rotate all pages by 90 degrees.
func rotatePdf(inputPath string, degrees int64, outputPath string) error {
	pdfWriter := pdf.NewPdfWriter()

	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return err
	}

	// Try decrypting both with given password and an empty one if that fails.
	if isEncrypted {
		auth, err := pdfReader.Decrypt([]byte(""))
		if err != nil {
			return err
		}
		if !auth {
			return errors.New("Unable to decrypt pdf with empty pass")
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		// page, err := pdfReader.GetPage(pageNum)
		page, err := pdfReader.GetPageAsPdfPage(pageNum)
		if err != nil {
			return err
		}

		// Do the rotation.
		var rotation int64 = 0
		if page.Rotate != nil {
			rotation = *(page.Rotate)
		}
		rotation += degrees // Rotate by 90 deg.
		page.Rotate = &rotation

		// err = pdfWriter.AddPage(page)
		pageObj := page.GetPageAsIndirectObject()
		err = pdfWriter.AddPage(pageObj)
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
