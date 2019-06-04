/*
pdf分割
*/
package main

import (
	"fmt"
	"os"

	unipdf "github.com/zouhuigang/pdf"
)

func main() {
	pdfWriter := unipdf.NewPdfWriter()
	pageFrom:=3
	pageTo:=8

	f, err := os.Open("ori_16.pdf")
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

	if numPages < pageTo {
		//return fmt.Errorf("numPages (%d) < pageTo (%d)", numPages, pageTo)
		fmt.Println(err.Error())
	}

	for i := pageFrom; i <= pageTo; i++ {
		pageNum := i

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = pdfWriter.AddPage(page)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	outFile := fmt.Sprintf("sp_%d_%d.pdf", pageFrom, pageTo)
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
