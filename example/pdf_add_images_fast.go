/*
 * Add images to a PDF file, one image per page.
 * Faster version, optimized using bimg/LibVPS.
 *
 * Run as: go run pdf_add_images_fast.go output.pdf img1.jpg img2.jpg img3.png ...
因为虚拟机磁盘空间有限，所以没有测试
 安裝文档
 https://github.com/h2non/bimg/tree/v1.0.19

 go get -u gopkg.in/h2non/bimg.v1
curl -s https://raw.githubusercontent.com/h2non/bimg/master/preinstall.sh | sudo bash -


Perhaps you should add the directory containing `vips.pc'


### 使用libvips来操作图像

查找:
pkg-config --cflags vips vips vips vips


vips -v

yum  install libvips-dev
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/h2non/bimg.v1"

	//  unicommon "github.com/unidoc/unidoc/common"
	//  unilicense "github.com/unidoc/unidoc/license"
	unipdf "github.com/zouhuigang/pdf"
)

// Fast image handling with bimg/LibVPS.
type FastImageHandler struct{}

// Read an input image file and prepare a native Image object.
func (this FastImageHandler) Read(reader io.Reader) (*unipdf.Image, error) {
	buffer, err := ioutil.ReadAll(reader)

	img, err := bimg.NewImage(buffer).Convert(bimg.JPEG)
	if err != nil {
		return nil, err
	}

	size, err := bimg.NewImage(img).Size()
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(img)

	image := unipdf.Image{}
	image.Width = int64(size.Width)
	image.Height = int64(size.Height)
	image.Data = buf

	return &image, nil
}

// Not implemented yet.
func (this FastImageHandler) Compress(input *unipdf.Image, quality int64) (*unipdf.Image, error) {
	return input, nil
}

//  func initUniDoc(licenseKey string) error {
// 	 if len(licenseKey) > 0 {
// 		 err := unilicense.SetLicenseKey(licenseKey)
// 		 if err != nil {
// 			 return err
// 		 }
// 	 }

// 	 // To make the library log we just have to initialise the logger which satisfies
// 	 // the unipdf.Logger interface, unicommon.DummyLogger is the default and
// 	 // does not do anything. Very easy to implement your own.
// 	 unicommon.SetLogger(unicommon.DummyLogger{})

// 	 // Set the fast image handler as the image handler.
// 	 unipdf.SetImageHandler(FastImageHandler{})

// 	 return nil
//  }

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: go run pdf_add_images_fast.go output.pdf img1.jpg img2.jpg ...\n")
		os.Exit(1)
	}

	outputPath := os.Args[1]
	inputPaths := os.Args[2:len(os.Args)]

	//  err := initUniDoc("")
	//  if err != nil {
	// 	 fmt.Printf("Error: %v\n", err)
	// 	 os.Exit(1)
	//  }

	err := imagesToPdf(inputPaths, outputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Complete, see output file: %s\n", outputPath)
}

// Images to PDF.
func imagesToPdf(inputPaths []string, outputPath string) error {
	pdfWriter := unipdf.NewPdfWriter()

	unipdf.Log.Debug("Inputs: %v", inputPaths)

	// Make the document structure.
	for idx, imgPath := range inputPaths {
		unipdf.Log.Debug("Image: %s", imgPath)
		// Open the image file.
		reader, err := os.Open(imgPath)
		if err != nil {
			unipdf.Log.Error("Error opening file: %s", err)
			return err
		}
		defer reader.Close()

		img, err := unipdf.ImageHandling.Read(reader)
		if err != nil {
			unipdf.Log.Error("Error loading image: %s", err)
			return err
		}

		height := 612 * float64(img.Height) / float64(img.Width)

		// Make a page.
		page := unipdf.NewPdfPage()
		bbox := unipdf.PdfRectangle{0, 0, 612, height}
		page.MediaBox = &bbox

		imgName := unipdf.PdfObjectName(fmt.Sprintf("Im%d", idx+1))

		ximg, err := unipdf.NewXObjectImage(imgName, img)

		if err != nil {
			unipdf.Log.Error("Failed to create xobject image: %s", err)
			return err
		}
		page.AddImageResource(imgName, ximg)

		gs0 := unipdf.PdfObjectDictionary{}
		name := unipdf.PdfObjectName("Normal")
		gs0["BM"] = &name
		page.AddExtGState("GS0", &gs0)

		contentStr := fmt.Sprintf("q\n"+
			"/GS0 gs\n"+
			"612 0 0 %.0f 0 0 cm\n"+
			"/%s Do\n"+
			"Q", height, imgName)
		page.AddContentStreamByString(contentStr)

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
