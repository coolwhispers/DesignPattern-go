package factory

type iFormat interface {
	Run() string
}

type jsonFormat struct {
}

func (f *jsonFormat) Run() string {
	return "json"
}

type xmlFormat struct {
}

func (f *xmlFormat) Run() string {
	return "xml"
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
