package iomanager

type IOManager interface {
	ReadLinesFromFile() ([]string, error)
	WriteJSONToFile(data interface{}) error
}
