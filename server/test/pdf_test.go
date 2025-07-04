package test

import (
	"code.sajari.com/docconv"
	"fmt"
	"log"
	"testing"
)

func TestParsePDF(t *testing.T) {
	res, err := docconv.ConvertPath("testdocx.docx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
