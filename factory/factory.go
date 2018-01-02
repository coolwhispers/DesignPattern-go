package factory

import (
	"fmt"
)

type iFormat interface {
	Run()
}

type jsonFormat struct {
}

func (f *jsonFormat) Run() {
	fmt.Println("json")
}

type xmlFormat struct {
}

func (f *xmlFormat) Run() {
	fmt.Println("xml")
}

func getFormat(formatType string) iFormat {
	var foramt iFormat

	switch formatType {
	case "JSON":
		foramt = &jsonFormat{}
	case "XML":
		foramt = &xmlFormat{}
	default:
		panic("no foramt")
	}
	return foramt
}

func factory() {

	f := getFormat("XML")

	f.Run()

}
