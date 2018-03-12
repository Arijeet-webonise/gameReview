package PDFWritter

import (
	"net/http"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/jung-kurt/gofpdf"
)

type HTMLToPDF struct {
	pages []string
}

type HTMLToPDFInterface interface {
	AppPage(...string)
	Generate(...string) error
}

func (hpdf *HTMLToPDF) AddPage(newpages ...string) {
	pages := hpdf.pages
	pages = append(pages, newpages...)
	hpdf.pages = pages
}

func (hpdf *HTMLToPDF) Generate2(w http.ResponseWriter) (err error) {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 20)
	_, lineHt := pdf.GetFontSize()
	pdf.Write(lineHt, "Test PDF")
	pdf.SetFont("", "U", 0)
	link := pdf.AddLink()
	pdf.WriteLinkID(lineHt, "here", link)
	pdf.SetFont("", "", 0)

	for _, htmlStr := range hpdf.pages {
		pdf.AddPage()
		pdf.SetLink(link, 0, -1)
		pdf.SetLeftMargin(15)
		pdf.SetFontSize(14)
		_, lineHt = pdf.GetFontSize()
		html := pdf.HTMLBasicNew()
		html.Write(lineHt, htmlStr)
	}

	err = pdf.Output(w)
	return

}

func (hpdf *HTMLToPDF) Generate(w http.ResponseWriter) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	for _, page := range hpdf.pages {
		pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(page)))
	}
	pdfg.Title.Set("Test PDF")
	if err = pdfg.Create(); err != nil {
		return err
	}

	_, err = w.Write(pdfg.Bytes())
	return err
}
