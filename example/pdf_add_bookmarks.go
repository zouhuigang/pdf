/*
 * Add bookmarks to pdf file.
 * Adds each page as a bookmark in the output file.
 *
 * Run as: go run pdf_add_bookmarks.go 1_2.pdf output.pdf
 */

package main

import (
	"fmt"
	"os"

	unipdf "github.com/zouhuigang/pdf"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Example adds bookmarks with page number referring to each page.\n")
		fmt.Printf("Usage: go run pdf_add_bookmarks.go input.pdf output.pdf\n")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	err := addBookmarks(inputPath, outputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Complete, see output file: %s\n", outputPath)
}

func addBookmarks(inputPath string, outputPath string) error {
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

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return err
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	// The outer containing tree.
	outlineTree := unipdf.NewPdfOutlineTree()
	isFirst := true
	var node *unipdf.PdfOutlineItem

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPageAsPdfPage(pageNum)
		if err != nil {
			return err
		}

		outputPage := page.GetPageAsIndirectObject()
		err = pdfWriter.AddPage(outputPage)
		if err != nil {
			return err
		}

		item := unipdf.NewOutlineBookmark(fmt.Sprintf("Page %d", i+1), outputPage)
		item.Parent = &outlineTree.PdfOutlineTreeNode

		if isFirst {
			outlineTree.First = &item.PdfOutlineTreeNode
			node = item
			isFirst = false
		} else {
			node.Next = &item.PdfOutlineTreeNode
			item.Prev = &node.PdfOutlineTreeNode
			node = item
		}
	}
	outlineTree.Last = &node.PdfOutlineTreeNode
	pdfWriter.AddOutlineTree(&outlineTree.PdfOutlineTreeNode)

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
