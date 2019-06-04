/*
 * List bookmarks from a pdf file (get the table of contents).
 *
 * Run as: go run pdf_list_bookmarks.go input.pdf
 */

package main

import (
	"fmt"
	"os"

	unipdf "github.com/zouhuigang/pdf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run pdf_list_bookmarks.go input.pdf\n")
		os.Exit(1)
	}

	inputPath := os.Args[1]

	err := listBookmarks(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func listBookmarks(inputPath string) error {
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

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return err
		}
	}

	_, flattenedTitles, err := pdfReader.GetOutlinesFlattened()
	fmt.Printf("--------------------\n")
	fmt.Printf("Table of contents:\n")
	fmt.Printf("--------------------\n")
	for _, title := range flattenedTitles {
		fmt.Printf("%s\n", title)
	}

	return nil
}
