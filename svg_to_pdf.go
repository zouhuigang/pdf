package pdf

// import (
// 	"io/ioutil"

// 	"github.com/jung-kurt/gofpdf"
// )

// func GeneratePDF(svgFile string, outFile string) error {
// 	var (
// 		sig gofpdf.SVGBasicType
// 		err error
// 	)

// 	read, err := ioutil.ReadFile(svgFile)
// 	if err != nil {
// 		return err
// 	}

// 	sig, err = gofpdf.SVGBasicParse(read)
// 	if err != nil {
// 		return err
// 	}
// 	scale := 100 / sig.Wd
// 	scaleY := 30 / sig.Ht
// 	if scale > scaleY {
// 		scale = scaleY
// 	}
// 	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
// 	pdf.SVGBasicWrite(&sig, scale)

// 	err = pdf.OutputFileAndClose(outFile)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
