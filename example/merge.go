/*
pdf合成
*/
package main

import (
	"fmt"
	"os"

	unipdf "github.com/zouhuigang/pdf"
)

func main() {
	pdfWriter := unipdf.NewPdfWriter()
	var inputPaths []string = []string{"1_2.pdf", "3_7.pdf", "8_16.pdf"}
	for _, inputPath := range inputPaths {
		f, err := os.Open(inputPath)
		if err != nil {
			fmt.Println(err.Error())
		}

		defer f.Close()

		pdfReader, err := unipdf.NewPdfReader(f)
		if err != nil {
			fmt.Println(err.Error())
		}

		isEncrypted, err := pdfReader.IsEncrypted()
		if err != nil {
			fmt.Println(err.Error())
		}

		if isEncrypted {
			_, err = pdfReader.Decrypt([]byte(""))
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		numPages, err := pdfReader.GetNumPages()
		if err != nil {
			fmt.Println(err.Error())
		}

		for i := 0; i < numPages; i++ {
			pageNum := i + 1

			page, err := pdfReader.GetPage(pageNum)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = pdfWriter.AddPage(page)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

	}

	outFile := fmt.Sprintf("%s.pdf", "ori_16")
	fWrite, err := os.Create(outFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		fmt.Println(err.Error())
	}
}
