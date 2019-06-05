// package main

// import (
// 	"fmt"
// 	"path/filepath"

// 	"github.com/jung-kurt/gofpdf"
// 	// unipdf "github.com/zouhuigang/pdf"
// )

// var gofpdfDir string

// func PdfDir() string {
// 	return filepath.Join(gofpdfDir, "pdf")
// }
// func ImageFile(fileStr string) string {
// 	return filepath.Join(gofpdfDir, "image", fileStr)
// }

// func Filename(baseStr string) string {
// 	return PdfFile(baseStr + ".pdf")
// }

// func PdfFile(fileStr string) string {
// 	return filepath.Join(PdfDir(), fileStr)
// }

// func Summary(err error, fileStr string) {
// 	if err == nil {
// 		fileStr = filepath.ToSlash(fileStr)
// 		fmt.Printf("Successfully generated %s\n", fileStr)
// 	} else {
// 		fmt.Println(err)
// 	}
// }

// func main() {
// 	const (
// 		fontPtSize = 16.0
// 		wd         = 100.0
// 		sigFileStr = "4.svg"
// 	)
// 	var (
// 		sig gofpdf.SVGBasicType
// 		err error
// 	)
// 	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
// 	pdf.SetFont("Times", "", fontPtSize)
// 	lineHt := pdf.PointConvert(fontPtSize)
// 	pdf.AddPage()
// 	pdf.SetMargins(10, 10, 10)
// 	htmlStr := `This example renders a simple ` +
// 		`<a href="http://www.w3.org/TR/SVG/">SVG</a> (scalable vector graphics) ` +
// 		`image that contains only basic path commands without any styling, ` +
// 		`color fill, reflection or endpoint closures. In particular, the ` +
// 		`type of vector graphic returned from a ` +
// 		`<a href="http://willowsystems.github.io/jSignature/#/demo/">jSignature</a> ` +
// 		`web control is supported and is used in this example.`
// 	html := pdf.HTMLBasicNew()
// 	html.Write(lineHt, htmlStr)
// 	sig, err = gofpdf.SVGBasicFileParse(ImageFile(sigFileStr))
// 	if err == nil {
// 		scale := 100 / sig.Wd
// 		scaleY := 30 / sig.Ht
// 		if scale > scaleY {
// 			scale = scaleY
// 		}
// 		pdf.SetLineCapStyle("round")
// 		pdf.SetLineWidth(0.25)
// 		pdf.SetDrawColor(0, 0, 128)
// 		pdf.SetXY((210.0-scale*sig.Wd)/2.0, pdf.GetY()+10)
// 		pdf.SVGBasicWrite(&sig, scale)
// 	} else {
// 		pdf.SetError(err)
// 	}
// 	fileStr := Filename("Fpdf_SVGBasicWrite")
// 	err = pdf.OutputFileAndClose(fileStr)
// 	Summary(err, fileStr)

// 	// err := unipdf.GeneratePDF("1.svg", "1_svg.pdf")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
