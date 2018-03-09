package PDFWritter

import (
	"net/http"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
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

func (hpdf *HTMLToPDF) Generate(w http.ResponseWriter) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	for _, page := range hpdf.pages {
		pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(page)))
	}

	if err = pdfg.Create(); err != nil {
		return err
	}

	_, err = w.Write(pdfg.Bytes())
	return err
}
