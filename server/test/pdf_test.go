package test

import (
	"code.sajari.com/docconv"
	"fmt"
	"log"
	"testing"
)

func TestParsePDF(t *testing.T) {
	res, err := docconv.ConvertPath("2025信息安全专业答辩-第6组.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
