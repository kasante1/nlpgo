package getpdfcontent

import (
    "log"
	"bytes"
	"github.com/ledongthuc/pdf"
)

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)

	if err != nil {
		return "", err
	}

    // remember close file
    defer f.Close()

	var buf bytes.Buffer
    b, err := r.GetPlainText()
    if err != nil {
        return "", err
    }
    buf.ReadFrom(b)
	return buf.String(), nil
}

func Getpdfcontent(filename string) (string, error){
	pdf.DebugOn = true
	content, err := readPdf(filename)
	if err != nil {
		log.Fatal(err)
	}
	return content, err

	
}