/*
 * Add watermark to each page of a PDF file.
 *
 * Run as: go run pdf_watermark_image.go input.pdf output.pdf watermark.png
 */

package main

import (
	"fmt"
	"os"

	unipdf "github.com/zouhuigang/pdf"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("go run pdf_watermark_image.go input.pdf output.pdf watermark.jpg\n")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]
	watermarkPath := os.Args[3]

	err := addWatermarkImage(inputPath, outputPath, watermarkPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Complete, see output file: %s\n", outputPath)
}

// Watermark pdf file based on an image.
func addWatermarkImage(inputPath string, outputPath string, watermarkPath string) error {

	unipdf.Log.Debug("Input PDF: %v", inputPath)
	unipdf.Log.Debug("Watermark image: %s", watermarkPath)

	pdfWriter := unipdf.NewPdfWriter()

	// Open the watermark image file.
	reader, err := os.Open(watermarkPath)
	if err != nil {
		unipdf.Log.Error("Error opening file: %s", err)
		return err
	}
	defer reader.Close()

	watermarkImg, err := unipdf.ImageHandling.Read(reader)
	if err != nil {
		unipdf.Log.Error("Error loading image: %s", err)
		return err
	}

	// Read the input pdf file.
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

	// If encrypted, try to decrypt with an empty password.
	if isEncrypted {
		// Fails, try fallback with empty password.
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return err
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	imgName := unipdf.PdfObjectName("Imw0")
	ximg, err := unipdf.NewXObjectImage(imgName, watermarkImg)
	if err != nil {
		unipdf.Log.Error("Failed to create xobject image: %s", err)
		return err
	}

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		// Read the page.
		page, err := pdfReader.GetPageAsPdfPage(pageNum)
		if err != nil {
			return err
		}

		wmOpt := unipdf.WatermarkImageOptions{}
		wmOpt.Alpha = 0.5
		wmOpt.FitToWidth = true
		wmOpt.PreserveAspectRatio = true

		err = page.AddWatermarkImage(ximg, wmOpt)
		if err != nil {
			return err
		}

		err = pdfWriter.AddPage(page.GetPageAsIndirectObject())
		if err != nil {
			unipdf.Log.Error("Failed to add page: %s", err)
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
